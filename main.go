package main
import (
    "fmt"
    "log"
    "net/http"
    //"strconv"
    "errors"
    //"reflect"

    "github.com/manifoldco/promptui"
    "github.com/PuerkitoBio/goquery"
)


func main() {

  validate := func(input string) error {
      //_, err := strconv.ParseFloat(input, 64)
      if input == ""{
        return errors.New("Input your search")
      }
      return nil
  }

  prompt := promptui.Prompt{
    Label: "Type the words that you want to know",
    Validate: validate,
  }

  var result, err = prompt.Run()


  if err != nil {
    fmt.Printf("Prompt failed %v\n", err)
    return
  }

  const base_url = "https://ja.wikipedia.org/wiki/"
  var url = base_url + result

  if result != "" {
    fmt.Println("Your choise is ", result)
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
}
