package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"time"
)

func FindRoleIDByName(s *discordgo.Session, RoleName string) string {

	roles, err := s.GuildRoles(s.State.Guilds[0].ID)
	if err != nil {
		return ""
	}

	for _, role := range roles {
		if role.Name == RoleName {
			return role.ID
		}
	}
	return ""
}

func RainbowRole(s *discordgo.Session) {
	for {
		rainbowRoleID := FindRoleIDByName(s, "Meaningless Color Role")
		fmt.Println("Found role ID: " + rainbowRoleID)
		// Change the role color to a random one
		randomColor := rand.Intn(16777215)
		newColorParams := discordgo.RoleParams{Name: "Meaningless Color Role", Color: &randomColor}

		_, err := s.GuildRoleEdit(s.State.Guilds[0].ID, rainbowRoleID, &newColorParams)
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Minute * 30)
	}
}
