package commands

import (
	"chat-api-bot/model"
	"chat-api-bot/rabbitmq"
	"chat-api-bot/service"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var StockCommand = Command{
	handleStock,
}

func handleStock(args string, cmd model.CommandTransport) error {
	resp, err := service.GetStock(args)
	fmt.Println(args)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	byteBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	parsedResp := strings.Split(string(byteBody), "\n")
	if len(parsedResp) < 2 {
		return errors.New("invalid stock code")
	}
	values := strings.Split(parsedResp[1], ",")
	if len(values) < 7 {
		return errors.New("invalid stock code")
	}

	val, err := strconv.ParseFloat(values[6], 64)

	if err != nil {
		return errors.New("invalid stock code")
	}
	data, _ := json.Marshal(model.CommandTransport{
		RoomId:   cmd.RoomId,
		Username: "BOT",
		Text:     fmt.Sprintf("%s quote is $%.2f per share", values[0], val),
	})
	rabbitmq.Push(data)

	return nil
}
