
this is simple tweak to the httputil.NewSingleHostReverseProxy Director which will cause the request host to be rewritten to equal the target host.

see source code and these discussins:
* https://github.com/golang/go/issues/5692
* https://github.com/golang/go/issues/7618

example invocation:
```golang
proxy := reverseProxyHostRewrite.New("https://example.com/")
```
