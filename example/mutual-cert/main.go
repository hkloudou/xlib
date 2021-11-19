package main

import (
	"crypto/x509/pkix"
	"embed"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hkloudou/xlib/xcert"
)

//go:embed views
var views embed.FS
var home, _ = os.UserHomeDir()
var basedir, _ = filepath.Abs(home + "/.mutual-cert/ca")

func getList() []string {
	arr := make([]string, 0)
	dir_list, e := ioutil.ReadDir(basedir)
	if e != nil {
		fmt.Println("read dir error")
		return nil
	}
	for _, v := range dir_list {
		if v.IsDir() {
			arr = append(arr, v.Name())
		}
	}
	return arr
}

func main() {
	log.Println("basedir", basedir)
	r := gin.Default()
	tmpl := template.Must(template.New("").ParseFS(views, "views/*"))
	r.SetHTMLTemplate(tmpl)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"list": getList(),
		})
	})
	r.POST("/download/server", func(c *gin.Context) {
		ee := `^[0-9a-zA-Z ]+$`
		reg := regexp.MustCompile(`^[0-9a-zA-Z ]+$`)
		// dns := strings.Split(strings.TrimSpace(c.PostForm("dns")), "\n")
		dns := make([]string, 0)
		ip := make([]net.IP, 0)
		for _, v := range strings.Split(strings.TrimSpace(c.PostForm("dns")), "\n") {
			tmp := strings.TrimSpace(v)
			if len(tmp) == 0 {
				continue
			}
			mt, _ := regexp.MatchString(`^[0-9a-zA-Z*.]+$`, tmp)
			if !mt {
				c.AbortWithStatusJSON(200, gin.H{"err": "err parame dns"})
				return
			}
			dns = append(dns, tmp)
		}

		for _, v := range strings.Split(strings.TrimSpace(c.PostForm("ip")), "\n") {
			tmp := strings.TrimSpace(v)
			if len(tmp) == 0 {
				continue
			}
			mt, _ := regexp.MatchString(`^[0-9.]+$`, tmp)
			if !mt {
				c.AbortWithStatusJSON(200, gin.H{"err": "err parame ip:" + tmp + "|"})
				return
			}
			// if tmpip, _, err := net.ParseIP(tmp); err != nil {
			// 	c.AbortWithStatusJSON(200, gin.H{"err": "err parame2 ip:" + tmp + "|"})
			// 	return
			// } else {
			ip = append(ip, net.ParseIP(tmp))
			// }
		}

		name := strings.TrimSpace(c.PostForm("name"))
		ca := c.PostForm("ca")

		if name == "" {
			c.AbortWithStatusJSON(200, gin.H{"err": "parame cn not found"})
			return
		} else if !reg.MatchString(name) {
			c.AbortWithStatusJSON(200, gin.H{"err": "parame name not match exp:" + ee})
			return
		}

		if !IsDir(fmt.Sprintf(basedir+"/%s/", ca)) {
			c.AbortWithStatusJSON(200, gin.H{"err": "ca not exist"})
			return
		}
		if cert, err := xcert.NewCert(); err != nil {
			c.AbortWithStatusJSON(200, gin.H{"err": "ca key not exist"})
			return
		} else if caCertByte, err := ioutil.ReadFile(fmt.Sprintf(basedir+"/%s/ca.pem", ca)); err != nil {
			c.AbortWithStatusJSON(200, gin.H{"err": "ca not exist"})
			return
		} else if caKeyByte, err := ioutil.ReadFile(fmt.Sprintf(basedir+"/%s/ca.key", ca)); err != nil {
			c.AbortWithStatusJSON(200, gin.H{"err": "ca key not exist"})
			return
		} else if caCert, caKey, err := readPEMCert(caCertByte, caKeyByte); err != nil {
			log.Println("err", err)
			c.AbortWithStatusJSON(200, gin.H{"err": "ca err"})
			return
		} else if err := cert.Sign(true, caCert, caKey, pkix.Name{
			CommonName: name,
		}, dns, ip); err != nil {
			log.Println("err", err)
			c.AbortWithStatusJSON(200, gin.H{"err": "sign err"})
			return
		} else if z, err := packZip("server", caCertByte, cert.Pub.Bytes(), cert.Pri.Bytes()); err != nil {
			log.Println("err", err)
			c.AbortWithStatusJSON(200, gin.H{"err": "zip err"})
			return
		} else {
			name := fmt.Sprintf("%s.zip", name)
			c.Header("Content-Type", "application/octet-stream")
			//强制浏览器下载
			c.Header("Content-Disposition", "attachment; filename="+name)
			//浏览器下载或预览
			c.Header("Content-Disposition", "inline;filename="+name)
			c.Header("Content-Transfer-Encoding", "binary")
			c.Data(200, "application/x-zip-compressed", z)
			// c.String(200, "生成成功")
		}
	})

	r.POST("/download/client", func(c *gin.Context) {
		/*

		 */
		ee := `^[0-9a-zA-Z ]+$`
		reg := regexp.MustCompile(`^[0-9a-zA-Z ]+$`)
		name := strings.TrimSpace(c.PostForm("name"))
		sn := strings.TrimSpace(c.PostForm("sn"))
		ca := c.PostForm("ca")

		if name == "" {
			c.AbortWithStatusJSON(200, gin.H{"err": "parame name not found"})
			return
		} else if !reg.MatchString(name) {
			c.AbortWithStatusJSON(200, gin.H{"err": "parame name not match exp:" + ee})
			return
		}

		if sn == "" {
			c.AbortWithStatusJSON(200, gin.H{"err": "parame sn not found"})
			return
		} else if !reg.MatchString(sn) {
			c.AbortWithStatusJSON(200, gin.H{"err": "parame sn not match exp:" + ee})
			return
		}

		if !IsDir(fmt.Sprintf(basedir+"/%s/", ca)) {
			c.AbortWithStatusJSON(200, gin.H{"err": "ca not exist"})
			return
		}
		if cert, err := xcert.NewCert(); err != nil {
			c.AbortWithStatusJSON(200, gin.H{"err": "ca key not exist"})
			return
		} else if caCertByte, err := ioutil.ReadFile(fmt.Sprintf(basedir+"/%s/ca.pem", ca)); err != nil {
			c.AbortWithStatusJSON(200, gin.H{"err": "ca not exist"})
			return
		} else if caKeyByte, err := ioutil.ReadFile(fmt.Sprintf(basedir+"/%s/ca.key", ca)); err != nil {
			c.AbortWithStatusJSON(200, gin.H{"err": "ca key not exist"})
			return
		} else if caCert, caKey, err := readPEMCert(caCertByte, caKeyByte); err != nil {
			log.Println("err", err)
			c.AbortWithStatusJSON(200, gin.H{"err": "ca err"})
			return
		} else if err := cert.Sign(false, caCert, caKey, pkix.Name{
			CommonName:   name,
			SerialNumber: sn,
			// ExtraNames: []pkix.AttributeTypeAndValue{
			// 	// pkix.AttributeTypeAndValue{}
			// },
			// ExtraNames: []pkix.AttributeTypeAndValue{
			// 	// pkix.
			// },
		}, nil, nil); err != nil {
			log.Println("err", err)
			c.AbortWithStatusJSON(200, gin.H{"err": "sign err"})
			return
		} else if z, err := packZip("client", caCertByte, cert.Pub.Bytes(), cert.Pri.Bytes()); err != nil {
			log.Println("err", err)
			c.AbortWithStatusJSON(200, gin.H{"err": "zip err"})
			return
		} else {
			name := fmt.Sprintf("%s.zip", name)
			c.Header("Content-Type", "application/octet-stream")
			//强制浏览器下载
			c.Header("Content-Disposition", "attachment; filename="+name)
			//浏览器下载或预览
			c.Header("Content-Disposition", "inline;filename="+name)
			c.Header("Content-Transfer-Encoding", "binary")
			c.Data(200, "application/x-zip-compressed", z)
			// c.String(200, "生成成功")
		}
	})
	r.GET("/ca/create", func(c *gin.Context) {
		ee := `^[0-9a-zA-Z ]+$`

		reg := regexp.MustCompile(ee)
		name := strings.TrimSpace(c.Query("cn"))

		// unicode.IsLetter
		if name == "" {
			c.AbortWithStatusJSON(200, gin.H{"err": "parame cn not found"})
			return
		} else if !reg.MatchString(name) {
			c.AbortWithStatusJSON(200, gin.H{"err": "parame cn not match exp:" + ee})
			return
		}
		if IsDir(fmt.Sprintf(basedir+"/%s/", name)) {
			c.AbortWithStatusJSON(200, gin.H{"err": "already exist"})
			return
		}
		os.MkdirAll(fmt.Sprintf(basedir+"/%s/", name), 0777)
		cert, err := xcert.NewCert()
		if err != nil {
			c.AbortWithError(500, err)
		} else if err := cert.SignCa(pkix.Name{CommonName: name}); err != nil {
			c.AbortWithError(500, err)
		} else if err := ioutil.WriteFile(fmt.Sprintf(basedir+"/%s/ca.pem", name), cert.Pub.Bytes(), 0644); err != nil {
			c.AbortWithError(500, err)
		} else if err := ioutil.WriteFile(fmt.Sprintf(basedir+"/%s/ca.key", name), cert.Pri.Bytes(), 0644); err != nil {
			c.AbortWithError(500, err)
		} else {
			c.Redirect(302, "/")
		}
	})
	r.Run()
}
