package main

import (
	"github.com/ChimeraCoder/anaconda"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	//api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("TOKEN_SECRET"))

	rand.Seed(time.Now().UnixNano())
	content, err := ioutil.ReadFile("desprogeades.txt")
    if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(content), "\n")

	line := lines[rand.Intn(len(lines))]

	//api.PostTweet(line, nil)
    log.Println(line)
}
