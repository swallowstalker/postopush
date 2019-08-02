package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/swallowstalker/postopush/go/forwarder"
)

func main() {
	chatID, err := strconv.Atoi(os.Getenv("CHAT_ID"))
	if err != nil {
		panic(fmt.Errorf("please provide correct chat ID"))
	}

	err = forwarder.Forward(os.Getenv("TOKEN"), os.Getenv("MESSAGE"), (int64)(chatID))
	if err != nil {
		panic(err)
	}
}
