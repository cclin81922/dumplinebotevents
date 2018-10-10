//    Copyright 2018 cclin
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func init() {
	b, err := linebot.New(
		os.Getenv("LineChannelSecret"),
		os.Getenv("LineChannelToken"),
	)
	if err != nil {
		log.Fatal(err)
	}

	bot = b
}

func main() {
	http.HandleFunc("/callback", callbackHandler)

	log.Println("Callback is serving at http://localhost:8080/callback")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
