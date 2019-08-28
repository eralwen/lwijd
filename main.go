package main

import (
    "log"
    "github.com/joho/godotenv"
    "os"
    "github.com/bwmarrin/discordgo"
)

var (
    Token string
)

func main () {
    loadEnv()
    newSession()
}

func loadEnv () {
    env_err := godotenv.Load()
    if env_err != nil {
        log.Println("Error loading env file!")
    }

    Token = os.Getenv("TOKEN")
}

func newSession() {
    dg, err := discordgo.New("Bot " + Token)
    if err != nil {
        log.Println(err)
    }

    dg.AddHandler(msgHandler)

    //Open socket and listen
    err = dg.Open()
    if err != nil {
        log.Println("error connecting," ,err)
        return
    }

    // Uhh, this makes CTRL-C exit
    <-make(chan struct{})
    return
}

func msgHandler(s *discordgo.Session, m *discordgo.MessageCreate){
    // Ignore messages made by self
    if  m.Author.ID == s.State.User.ID {
        return
    }

    if m.Content == "owo" {
        s.ChannelMessageSend(m.ChannelID, "uwu")
    }
}


