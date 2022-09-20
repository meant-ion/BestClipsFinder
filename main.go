package main

import (
	"bufio" // using this to make it easier for reading long lists of emotes
	"fmt"
	"os"
	"strings"

	"github.com/gempir/go-twitch-irc/v3"
)

// each individual emote's counter struct
type MsgText struct {
	Emote string
	Count int
}

func main() {

	// open and read contents of emotes list file into array
	file, err := os.Open("emotes.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var emotesList []MsgText

	for scanner.Scan() {
		tmp := MsgText{
			Emote: scanner.Text(),
			Count: 0,
		}
		emotesList = append(emotesList, tmp)
	}

	if err := scanner.Err(); err != nil {
		panic("err")
	}

	// No need to write anything to the IRC channel, just grab and go
	client := twitch.NewAnonymousClient()

	var msg_arr []MsgText

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
		for _, word := range strings.Split(message.Message, " ") {
			if isWantedEmote(emotesList, word) {
				tmp := MsgText{word, 0}
				msg_arr = append(msg_arr, tmp)
			}
		}

	})

	client.Join("pope_pontius")

	err = client.Connect()
	if err != nil {
		panic(err)
	}
}

func isWantedEmote(s []MsgText, str string) bool {
	for _, v := range s {
		if v.Emote == str {
			return true
		}
	}
	return false
}
