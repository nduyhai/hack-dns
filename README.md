# Hack DNS

Try investigate ```no such host``` error with ```net/http``` in Go

## Take a look

````shell
{"level":"info","ts":1718808165.168462,"caller":"hack-dns/main.go:46","msg":"done looking up dns","dns_info":{"Addrs":[],"Err":{"Err":"i/o timeout","Name":"jupiter","Server":"","IsTimeout":true,"IsTemporary":false,"IsNotFound":false},"Coalesced":false},"ms":1}
{"level":"info","ts":1718808165.1686358,"caller":"hack-dns/main.go:46","msg":"done looking up dns","dns_info":{"Addrs":[],"Err":{"Err":"i/o timeout","Name":"jupiter","Server":"","IsTimeout":true,"IsTemporary":false,"IsNotFound":false},"Coalesced":false},"ms":1}
{"level":"info","ts":1718808165.168742,"caller":"hack-dns/main.go:46","msg":"done looking up dns","dns_info":{"Addrs":[],"Err":{"Err":"i/o timeout","Name":"jupiter","Server":"","IsTimeout":true,"IsTemporary":false,"IsNotFound":false},"Coalesced":false},"ms":1}
{"level":"info","ts":1718808165.169848,"caller":"hack-dns/main.go:46","msg":"done looking up dns","dns_info":{"Addrs":[],"Err":{"Err":"no such host","Name":"jupiter","Server":"","IsTimeout":false,"IsTemporary":false,"IsNotFound":true},"Coalesced":true},"ms":78}
{"level":"info","ts":1718808165.169903,"caller":"hack-dns/main.go:46","msg":"done looking up dns","dns_info":{"Addrs":[],"Err":{"Err":"no such host","Name":"jupiter","Server":"","IsTimeout":false,"IsTemporary":false,"IsNotFound":true},"Coalesced":true},"ms":87}
{"level":"info","ts":1718808165.169908,"caller":"hack-dns/main.go:46","msg":"done looking up dns","dns_info":{"Addrs":[],"Err":{"Err":"no such host","Name":"jupiter","Server":"","IsTimeout":false,"IsTemporary":false,"IsNotFound":true},"Coalesced":true},"ms":18}
{
````

A request with a timeout that is too small will result in an ```I/O timeout```. 
Subsequently, another request with ```"Coalesced": true``` will return a ```no such host``` error. As you know, Coalesced is whether the ```Addrs``` were shared with another caller who was doing the same DNS lookup concurrently.


## Next action
TBD