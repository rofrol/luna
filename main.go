package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

func main() {
	client := &http.Client{}

	u, err := url.Parse("http://lunamademyday.pl/")
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux i686; rv:17.0) Gecko/20100101 Firefox/17.0")

	values := make(url.Values)
	values.Set("action", "ulikeprocess")
	values.Set("id", "170")
	fmt.Printf("%v\n", values)

	u2 := "http://lunamademyday.pl/wp-admin/admin-ajax.php"

	for i := 0; i < 2500; i++ {
		jar, err := cookiejar.New(nil)
		if err != nil {
			log.Fatal(err)
		}
		client.Jar = jar

		_, err = client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", client.Jar.Cookies(u)[0])

		res, err := client.PostForm(u2, values)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", res.Status)

		time.Sleep(500 * time.Millisecond)
	}
}
