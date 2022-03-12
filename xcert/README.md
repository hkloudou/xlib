# :zap: xcert
xcert is a easy way to manager x509 certificate

## Installation
``` sh
go get -u github.com/hkloudou/xlib/xcert
```


## Quick Start self signed
``` go
os.MkdirAll("./test", 0744)
tmpl := xcert.Template(
    xcert.PkixName(pkix.Name{Organization: []string{"Acme Co"}}),
    xcert.KeyUsage(x509.KeyUsageKeyEncipherment|x509.KeyUsageDigitalSignature|x509.KeyUsageCertSign),
    xcert.ExtKeyUsage(x509.ExtKeyUsageServerAuth),
    xcert.Hosts("localhost", "test.yourdomain.com", "127.0.0.1"),
)
pem, key, err := xcert.GenerateEcdsaCert(tmpl)
if err != nil {
    panic(err)
}
ioutil.WriteFile("./test/self.pem", pem, 0644)
ioutil.WriteFile("./test/self.key", key, 0644)
```

## Quick Start two way signed
``` go
os.MkdirAll("./test", 0744)
tmpl := xcert.Template(
    xcert.PkixName(pkix.Name{CommonName: "test"}),
    xcert.KeyUsage(x509.KeyUsageDigitalSignature|
        x509.KeyUsageKeyEncipherment|x509.KeyUsageCertSign|
        x509.KeyUsageCRLSign),
    xcert.ExtKeyUsage(x509.ExtKeyUsageAny),
)
pem, key, err := xcert.GenerateEcdsaCert(tmpl)
if err != nil {
    t.Fatal(err)
}
ioutil.WriteFile("./test/ca.pem", pem, 0644)
ioutil.WriteFile("./test/ca.key", key, 0644)
ca, caKey, err := xcert.ParseCertPair(pem, key)
pem, key, err = xcert.GenerateEcdsaCertWithParent(
    xcert.Template(
        xcert.PkixName(pkix.Name{CommonName: "server"}),
        xcert.Hosts("localhost", "127.0.0.1"),
        xcert.IsCa(false),
        xcert.ExtKeyUsage(x509.ExtKeyUsageServerAuth),
    ), ca, caKey)
if err != nil {
    t.Fatal(err)
}
ioutil.WriteFile("./test/server.pem", pem, 0644)
ioutil.WriteFile("./test/server.key", key, 0644)

pem, key, err = xcert.GenerateEcdsaCertWithParent(xcert.Template(
    xcert.PkixName(pkix.Name{CommonName: "client"}),
    xcert.IsCa(false),
    xcert.ExtKeyUsage(x509.ExtKeyUsageClientAuth),
), ca, caKey)
if err != nil {
    t.Fatal(err)
}
ioutil.WriteFile("./test/client.pem", pem, 0644)
ioutil.WriteFile("./test/client.key", key, 0644)
```

## use it
``` go
cfg, err := xcert.ParseTlsConfig(nil, "./test/self.pem", "./test/self.key")
if err != nil {
    t.Fatal(err)
}
cfg, err = xcert.ParseTlsConfig("./test/ca.pem", "./test/server.pem", "./test/server.key")
if err != nil {
    t.Fatal(err)
}
tls.Listen("tcp", ":443", cfg)
```

## Todolist
- Auther interface