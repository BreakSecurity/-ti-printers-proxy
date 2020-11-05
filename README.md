# ti-printers-proxy

## A simple go program that proxy wss requests to an http IP address (raw Byte are sent) 

To use execute with the following parameters:

```
wss-proxy -cert localhost.crt -key localhost.key -from IPADDRESS:PORT -to IPADDRESS:PORT
```

You can run many instances of the program in the background to handle multiple printers

Example:

```
wss-proxy -cert localhost.crt -key localhost.key -from localhost:9002 -to 192.168.123.101:9100 
wss-proxy -cert localhost.crt -key localhost.key -from localhost:9003 -to 192.168.123.102:9100
wss-proxy -cert localhost.crt -key localhost.key -from localhost:9004 -to 192.168.123.103:9100
wss-proxy -cert localhost.crt -key localhost.key -from localhost:9005 -to 192.168.123.104:9100
```
