package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

func FilterUsers(s *discordgo.Session, r *discordgo.GuildMemberAdd) {
	createdAt, err := discordgo.SnowflakeTimestamp(r.Member.User.ID)
	if err != nil {
		fmt.Println("error getting user creation time,", err)
		return
	}
	// print the number of hours since the account was created
	hoursSinceCreation := int64(time.Now().Sub(createdAt).Hours())
	if hoursSinceCreation < 168 {
		// Ban the user
		err := s.GuildBanCreateWithReason(r.GuildID, r.Member.User.ID, "Hades Ban", 0)
		if err != nil {
			fmt.Println("error banning user,", err)
			return
		}
	}
}
