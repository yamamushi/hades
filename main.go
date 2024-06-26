package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Variables used for command line parameters
var (
	ConfPath string
	Conf     Config
)

func init() {
	// Read our command line options
	flag.StringVar(&ConfPath, "c", "hades.conf", "Path to Config File")
	flag.Parse()

	_, err := os.Stat(ConfPath)
	if err != nil {
		log.Fatal("Config file is missing: ", ConfPath)
	}

	// Verify we can actually read our config file
	err = ReadConfig(ConfPath)
	if err != nil {
		log.Fatal("error reading config file at: ", ConfPath)
		return
	}

}

func main() {
	fmt.Println("\n\n|| Starting Hades Bot ||")
	log.SetOutput(io.Discard)

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Conf.DiscordConfig.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	defer func(dg *discordgo.Session) {
		err := dg.Close()
		if err != nil {
			panic(err)
		}
	}(dg)

	dg.Identify.Intents |= discordgo.IntentsAll

	dg.AddHandler(Ready)

	dg.AddHandler(FilterUsers)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	go RainbowRole(dg)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	err = dg.Close()
	if err != nil {
		panic(err)
	}
}

func Ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("Bot is up!")
}
