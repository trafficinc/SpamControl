<?php

$value = "demo1@hotmail.com";

$socketFile = '/path/to/spamcontrol/spam.sock';

$client = stream_socket_client("unix://$socketFile", $errno, $errorMessage);

if ($client === false) {
    throw new UnexpectedValueException("Failed to connect: $errorMessage");
}

fwrite($client, "{\"action\":\"url\",\"value\":\"$value\"}");
$isSpam = trim(stream_get_contents($client));
echo $isSpam;
var_dump($isSpam);
fclose($client);