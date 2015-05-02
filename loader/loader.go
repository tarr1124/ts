package loader

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "encoding/json"
  "strconv"
  "time"

  "github.com/kyeah/gohunt/gohunt"
)

type Feed interface {
  Display()
  Setter(name string, title []string, url []string)
}

type JsonData struct {
  Title   string `json:"title"`
  Url     string `json:"url"`
}

type ResultData struct {
  Name string
  Title []string
  Url []string
}

func (r *ResultData) Display() {
  pp("[" + r.Name + "]\n")
  for key, _ := range r.Title {
    pp(" - " + r.Title[key] + "\n")
    pp("   - " + r.Url[key] + "\n")
  }
  pp("\n")
}

func (r *ResultData) Setter(name string, title []string, url []string) {
  r.Name = name
  r.Title = title
  r.Url = url
}

func pp(str string) {
  fmt.Printf(str)
}

func perror(err error) {
  if err != nil {
    panic(err)
  }
}

func getIdsFromUrl(url string) (ids []int) {
  resp, err := http.Get(url)
  perror(err)
  jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)
  perror(err)
  err = json.Unmarshal([]byte(jsonDataFromHttp), &ids)
  perror(err)
  return ids
}

func getJsonDataFromUrl(url string) (jsonRes JsonData) {
  resp, err := http.Get(url)
  perror(err)
  jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)
  perror(err)
  err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonRes)
  perror(err)
  return jsonRes
}

func GetHNFeed(hn chan ResultData) {
  var result ResultData
  var HNTitle, HNUrl []string
  topStoriesUrl := "https://hacker-news.firebaseio.com/v0/topstories.json"
  jsonRes := getIdsFromUrl(topStoriesUrl)
  for _, value := range jsonRes[0:5] {
    url := "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(value) + ".json"
    HNRes := getJsonDataFromUrl(url)
    HNTitle = append(HNTitle, HNRes.Title)
    HNUrl = append(HNUrl, HNRes.Url)
    time.NewTimer(time.Second * 1)
  }
  result.Setter("HackerNews", HNTitle, HNUrl)
  hn <- result
}

func GetPHFeed(ph chan ResultData) {
  var result ResultData
  var PHTitle, PHUrl []string

  client := gohunt.NewUserClient("a0ad779df4746d96d4d87a5a589f786277c4c78699445fc70abdc08d4be77b45")

  posts, err := client.GetPosts()
  perror(err)

  for _, post := range posts[0:5] {
    PHTitle = append(PHTitle, post.Name)
    PHUrl = append(PHUrl, post.RedirectUrl)
  }

  result.Setter("ProductHunt", PHTitle, PHUrl)
  ph <- result
}
