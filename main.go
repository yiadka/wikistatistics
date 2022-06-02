package main
import (
    "fmt"
    "log"
    "net/http"

    "github.com/manifoldco/promptui"
    "github.com/PuerkitoBio/goquery"
)

const url = "https://ja.wikipedia.org/wiki/Python"

func main() {
  prompt := promptui.Select{
    Label: "Choose the number",
    Items: []int{1},
  }

  idx, result, err := prompt.Run()


  if err != nil {
    fmt.Printf("Prompt failed %v\n", err)
    return
  }

  if result == "1" {
    fmt.Println("Your choise is %d %q\n", idx, result)
    scraping()
  }
}

func scraping() {
  res, err1 := http.Get(url)
  if err1 != nil {
    log.Println(err1)
  }
  defer res.Body.Close()

  doc, _ := goquery.NewDocumentFromReader(res.Body)
  doc.Find(".mw-headline").Each(func(i int, s *goquery.Selection) {
    fmt.Println(s.Text())
  })

  fmt.Println("\n--------------------------------------------")
  doc.Find(".tocnumber").Each(func(i int, s *goquery.Selection) {
    fmt.Println(s.Text(), " ", s.Next().Text())
  })
}
