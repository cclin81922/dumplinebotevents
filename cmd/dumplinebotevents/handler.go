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
	"fmt"
	"log"
	"net/http"

	"github.com/cclin81922/dumplinebotevents/pkg/dumplinebotevents"
	"github.com/line/line-bot-sdk-go/linebot"
)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {

		// Dump to console
		eventLiteral := dumplinebotevents.Dump(event)
		fmt.Printf(eventLiteral)

		if event.Type == linebot.EventTypeMessage {
			// Dump to line
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				echoMessage := fmt.Sprintf("%s\n%s", message.Text, eventLiteral)
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(echoMessage)).Do(); err != nil {
					log.Print(err)
				}
			default:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(eventLiteral)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
