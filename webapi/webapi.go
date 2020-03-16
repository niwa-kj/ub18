 package main

// https://godoc.org/net/http
// https://godoc.org/net/url#URL

 import  "fmt"
 import  "io/ioutil"
 import  "net/url"
 import  "net/http"

 func makeUrl(host string, keyword string) *url.URL {
    u := &url.URL{}
	u.Scheme = "https"
	u.Host = host
	q := u.Query()
	q.Set("q", keyword)
	u.RawQuery = q.Encode()
	fmt.Println("URL ", u, &u ,*u)
	return u

 }

 func main(){
	u := &url.URL{}
	u = makeUrl("google.co.jp", "golang")
	fmt.Println("URL ", u, &u)

    res, err := http.Get(u.String())
	if err != nil {
	    fmt.Println("RES ERROR ", err)
	}
	defer res.Body.Close()
	fmt.Println("RES ", res)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	    fmt.Println("BODY ERROR ", err)
	}
	fmt.Println("BODY ", string(body))
 }

