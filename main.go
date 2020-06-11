package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// Emails : email struct
type Emails struct {
	Emails []Email `json:"emails"`
}

// Email : email struct
type Email struct {
	Email string `json:"email"`
}

type inputLoad struct {
	Action string `json:"action"`
	Value  string `json:"value"`
}

// badEmailCount : bad email counts
type badEmailCount struct {
	a int
}

// CheckSpamEmail : will check if email is spam
func CheckSpamEmail(emailInput string) bool {

	jsonFile, err := os.Open("emails.json")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("Successfully Opened banned emails.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var emails Emails

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &emails)

	// parse emailInput
	var emailIn = GetSlug(emailInput)

	badEmails := []badEmailCount{}
	for i := 0; i < len(emails.Emails); i++ {
		if strings.Contains(emails.Emails[i].Email, strings.ToLower(emailIn)) == true {
			badEmails = append(badEmails, badEmailCount{1})
		}
	}
	//fmt.Println(len(badEmails) > 0)
	defer jsonFile.Close()

	if len(badEmails) > 0 {
		return true
	}
	return false

}

// GetSlug : pull domain out of email
func GetSlug(emailIn string) string {
	components := strings.Split(emailIn, "@")
	domain := components[1]
	return domain
}

// test func
func dataHandler(c net.Conn) {
	// we create a decoder that reads directly from the socket
	d := json.NewDecoder(c)

	var msg inputLoad

	err := d.Decode(&msg)
	if err != nil {
		log.Fatal("Error decoding ", err)
		return
	}
	if msg.Action == "url" {
		spamCheck := CheckSpamEmail(msg.Value)

		if spamCheck == true {
			_, err = c.Write([]byte("true\n"))
		} else {
			_, err = c.Write([]byte("false\n"))
		}
	}

	if err != nil {
		log.Fatal("write error:", err)
	}

	c.Close()
}

func main() {
	log.Println("Starting spamcontrol server")
	const SOCK = "/path/to/spamcontrol/spam.sock"
	os.Remove(SOCK)
	unixListener, err := net.Listen("unix", SOCK)
	if err != nil {
		log.Fatal("Listen (UNIX socket): ", err)
	}

	//defer unixListener.Close()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func(unixListener net.Listener, c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		unixListener.Close()
		os.Exit(0)
	}(unixListener, sigc)

	for {
		fd, err := unixListener.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		//fmt.Println(fd.LocalAddr())
		go dataHandler(fd)
		//fd.Close()
	}
}
