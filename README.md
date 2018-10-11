# Installation

```
go get -u github.com/cclin81922/dumplinebotevents/cmd/dumplinebotevents
export PATH=$PATH:~/go/bin
```

# Commmand Line Usage

```
LineChannelSecret=... LineChannelToken=... dumplinebotevents

// then you can use line to test
```

// Package Usage

```
import "github.com/cclin81922/dumplinebotevents/pkg/dumplinebotevents"

func demo(event *linebot.Event) {
    eventLiteral := dumplinebotevents.Dump(event)
    fmt.Println(eventLiteral)
}
```

# For Developer

Run dumplinebotevents with ngrok

```
NgrokToken=... ngrok authtoken $NgrokToken
ngrok http 8080
LineChannelSecret=... LineChannelToken=... dumplinebotevents
```

# Related Resources

* [Tool ngrok to test linebot in development phase](https://cleanshadow.blogspot.com/2017/02/ngrokline-botwebhook.html)
