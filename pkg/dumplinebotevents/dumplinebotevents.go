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

package dumplinebotevents

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Dump ...
func Dump(e *linebot.Event) (s string) {
	template := `{` + "\n" +
		"\t" + `ReplyToken: %s` + "\n" +
		"\t" + `Type: %s` + "\n" +
		"\t" + `Timestamp: %s` + "\n" +
		"\t" + `Source: {` + "\n" +
		"\t\t" + `Type: %s` + "\n" +
		"\t\t" + `UserID: %s` + "\n" +
		"\t\t" + `GroupID: %s` + "\n" +
		"\t\t" + `RoomID: %s` + "\n" +
		"\t" + `}` + "\n" +
		"\t" + `Message: .(type) == %s` + "\n" +
		"\t" + `Postback: {` + "\n" +
		"\t\t" + `Data: -` + "\n" +
		"\t\t" + `Params: {` + "\n" +
		"\t\t\t" + `Date: -` + "\n" +
		"\t\t\t" + `Time: -` + "\n" +
		"\t\t\t" + `Datetime: -` + "\n" +
		"\t\t" + `}` + "\n" +
		"\t" + `}` + "\n" +
		"\t" + `Beacon: {` + "\n" +
		"\t\t" + `Hwid: -` + "\n" +
		"\t\t" + `Type: -` + "\n" +
		"\t\t" + `DeviceMessage: -` + "\n" +
		"\t" + `}` + "\n" +
		"\t" + `AccountLink: {` + "\n" +
		"\t\t" + `Result: -` + "\n" +
		"\t\t" + `Nonce: -` + "\n" +
		"\t" + `}` + "\n" +
		`}`

	eMessageType := "?"
	switch e.Message.(type) {
	case *linebot.TextMessage:
		eMessageType = "TextMessage"
	case *linebot.ImageMessage:
		eMessageType = "ImageMessage"
	case *linebot.VideoMessage:
		eMessageType = "VideoMessage"
	case *linebot.AudioMessage:
		eMessageType = "AudioMessage"
	case *linebot.FileMessage:
		eMessageType = "FileMessage"
	case *linebot.StickerMessage:
		eMessageType = "StickerMessage"
	case *linebot.LocationMessage:
		eMessageType = "LocationMessage"
	case *linebot.TemplateMessage:
		eMessageType = "TemplateMessage"
	case *linebot.ImagemapMessage:
		eMessageType = "ImagemapMessage"
	case *linebot.FlexMessage:
		eMessageType = "FlexMessage"
	}

	s = fmt.Sprintf(template, e.ReplyToken, e.Type, e.Timestamp,
		e.Source.Type, e.Source.UserID, e.Source.GroupID, e.Source.RoomID, eMessageType)
	return
}
