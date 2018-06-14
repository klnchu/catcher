// Catch Open Stack Component Release Notes.

package main

import (
	"fmt"
	"time"
	"net/http"
	
	// "github.com/PuerkitoBio/goquery"
)

const OS_BASE_URL = "https://docs.openstack.org/releasenotes/"

var components = []string{"nova", "neutron", "cinder"}
var versions = []string{"liberty", "mitaka", "newton", "ocata", "pike", "queens"}


func genUrl(component, version string) string{
	
	return fmt.Sprintf("%s/%s/%s.html", OS_BASE_URL, component, version)
}

func Do(component, version string) error {
	fmt.Println("------------------------------")
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
	// doc, err := goquery.NewDocumentFromReader(resp.Body)
	// if err != nil {
	// 	return err
	// }
	// doc.find(".section").Each(func(i int, s *goquery.Selection) {
	// 	// For each item found, get the band and title
	// 	band := s.Find("a").Text()
	// 	title := s.Find("i").Text()
	// 	fmt.Printf("Review %d: %s - %s\n", i, band, title)
	//   })
	return nil
}

func main(){
	fmt.Println("Let's Do it.")
	for _, component:= range  components {
		for _, version := range  versions {
			go func(_component, _version string){
				Do(_component, _version)
			}(component, version)
		}
	}
	time.Sleep(time.Second * 3)
	fmt.Println("Catch end")
}