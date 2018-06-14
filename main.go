// Catch Open Stack Component Release Notes.

package main

import (
	"os"
	"fmt"
	"strings"
	"net/http"
	
	"github.com/PuerkitoBio/goquery"
)

const OS_BASE_URL = "https://docs.openstack.org/releasenotes/"

var components = []string{"nova", "neutron", "cinder"}
var versions = []string{"liberty", "mitaka", "newton", "ocata", "pike", "queens"}


func genUrl(component, version string) string{
	
	return fmt.Sprintf("%s/%s/%s.html", OS_BASE_URL, component, version)
}

func genFileName(component, version string) string{
	return fmt.Sprintf("docs/os_%s_%s", component, version)
}

func Do(component, version string) error {
	fmt.Println("-----------------------------------------")
	fmt.Printf("Do %s: %s document.\n", component, version)
	
	url := genUrl(component, version)
	resp, err := http.Get(url)
	if err != nil{
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Printf("status code error: %d %s\n", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}
	file, err := os.Create(genFileName(component, version))
	if err != nil{
		return err
	}
	defer file.Close()

	body := doc.Find(".docs-body").Eq(0)
	framework := body.ChildrenFiltered(".section").Eq(0)
	framework.Find(".section").Each(func(i int, s *goquery.Selection){
		id, exist := s.Attr("id")
		if !exist{
			s.Next()
		}
		if strings.Contains(id, "new-features"){
			s.Find("ul").Each(func(j int, ul *goquery.Selection){
				file.WriteString(ul.Text())
			})
		}
	})
	
	file.Sync()

	return nil
}

func main(){
	fmt.Println("Let's Do it.")
	for _, component:= range  components {
		for _, version := range  versions {
			err := Do(component, version)
			if err != nil{
				fmt.Println(err)
			}
		}
	}
	fmt.Println("Catch end")
}