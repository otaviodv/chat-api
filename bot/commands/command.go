package commands

import (
	"chat-api-bot/helper"
	"chat-api-bot/model"
	"errors"
	"strings"
)

type Command struct {
	Handle func(args string, cmd model.CommandTransport) error
}

var commands map[string]Command

func parseCommand(commandStr string) (string, string) {

	trimmed := helper.TrimFirstRune(commandStr)

	spl := strings.Split(trimmed, "=")

	if len(spl) > 1 {
		return spl[0], spl[1]
	}

	return trimmed, ""
}

func init() {
	commands = map[string]Command{
		"stock": StockCommand,
	}
}

func CheckAndRunCommand(reqCmd model.CommandTransport) error {

	command, args := parseCommand(reqCmd.Text)

	if cmd, ok := commands[command]; ok {
		return cmd.Handle(args, reqCmd)
	}

	return errors.New("command not found")
}
