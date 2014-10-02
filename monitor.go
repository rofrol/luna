package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/yosssi/gohtml"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func monitor() {
	doc, err := goquery.NewDocument("http://lunamademyday.pl/lista-zgloszen/")
	ifErrLogFatal(err)

	doc.Find(".fLeft").Each(func(i int, s *goquery.Selection) {
		sel := s.Find(".wpulike")
		// ommit empty/white-space-only nodes
		if strings.TrimSpace(sel.Text()) != "" {
			var count int
			countStr := s.Find(".count-box").Text()
			if countStr != "" {
				count, err = strconv.Atoi(countStr)
				ifErrLogFatal(err)
			}

			idStr, ok := sel.Attr("id")
			if !ok {
				log.Fatal("Expected a value for the id attribute.")
			}
			re := regexp.MustCompile(`wp-ulike-([0-9]+)$`)
			var id int
			if foundId := re.FindStringSubmatch(idStr); foundId != nil {
				id, err = strconv.Atoi(foundId[1])
				ifErrLogFatal(err)
			}
			if count > 0 {
				fmt.Printf("i: %d, count: %d, id: %d\n", i, count, id)
			}
		}
	})

}

func printSel(s *goquery.Selection) {
	html, err := s.Html()
	ifErrLogFatal(err)
	fmt.Printf("%s\n", gohtml.Format(html))
}

func ifErrLogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
