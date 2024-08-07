package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/shayan0v0n/invoice-managment/utils"
	"github.com/urfave/cli/v2"
)

func RegisterActions(ctx *cli.Context) error {
	if ctx.NArg() < 2 || ctx.Args().Get(0) == "" || ctx.Args().Get(1) == "" {
		return cli.Exit("username and password are required", 1)
	}

	username := ctx.Args().Get(0)
	password := ctx.Args().Get(1)

	var filePath = fmt.Sprintf("database/%s_%d.json", username, time.Now().Unix())

	if !utils.IsFileExist(username) {
		file, _ := os.Create(filePath)
		content := map[string]string{"password": password, "status": "unregistered"}
		contentJson, _ := json.Marshal(content)

		file.Write([]byte(contentJson))

		defer file.Close()
	} else {
		return cli.Exit(fmt.Sprintf("%s already exists", filePath), 1)
	}

	return nil
}

func LoginActions(ctx *cli.Context) error {
	if ctx.NArg() < 2 || ctx.Args().Get(0) == "" || ctx.Args().Get(1) == "" {
		return cli.Exit("username and password are required", 1)
	}

	username := ctx.Args().Get(0)
	password := ctx.Args().Get(1)

	if !utils.IsFileExist(username) {
		return cli.Exit("username does not exist", 1)
	}

	getFile := utils.GetFileExist(username)
	data, err := os.ReadFile(getFile)
	if err != nil {
		return cli.Exit("failed to read user file", 1)
	}

	var JsonData map[string]interface{}
	err = json.Unmarshal(data, &JsonData)
	if err != nil {
		return cli.Exit("failed to parse user data", 1)
	}

	if password == JsonData["password"].(string) {
		// Update the status in the JSON data
		JsonData["status"] = "logged in"

		updatedData, err := json.Marshal(JsonData)
		if err != nil {
			return cli.Exit("failed to marshal updated data", 1)
		}

		err = os.WriteFile(getFile, updatedData, 0644)
		if err != nil {
			return cli.Exit("failed to write updated data to file", 1)
		}

		return cli.Exit("Successfully logged in", 0)
	}

	return cli.Exit("Failed to login", 1)
}
