package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vicanso/go-axios"
	"gopkg.in/myesui/uuid.v1"
)

func UploadImage(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.JSON(fiber.Map{"code": 40, "message": "image is empty or invalid type", "data": nil})
	}

	// generate new uuid for image name
	uniqueId := uuid.NewV1()

	// remove "- from imageName"

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// extract image extension from original file filename

	fileExt := strings.Split(file.Filename, ".")[1]

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	// save image to ./images dir
	err = ctx.SaveFile(file, fmt.Sprintf("./images/%s", image))

	if err != nil {
		return ctx.JSON(fiber.Map{"status": 500, "message": "Can't save", "data": nil})
	}

	// generate image url to serve to client using CDN

	imageUrl := fmt.Sprintf("http://61.28.238.162:3001/%s", image)

	// create meta data and send to client

	data := map[string]interface{}{

		"imageName": image,
		"url":       imageUrl,
		"header":    file.Header,
		"size":      file.Size,
	}
	ins := axios.NewInstance(&axios.InstanceConfig{
		EnableTrace: true,
		Client: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			},
		},

		Timeout: 5 * time.Second,

		OnDone: func(config *axios.Config, resp *axios.Response, err error) {
			fmt.Println(config)
			fmt.Println(resp)
			fmt.Println(err)
		},
	})

	url := "http://61.28.238.162:3001/" + "image"
	resp, err := ins.Post(url, data)
	if err != nil {
		panic(err)
	}
	buf, _ := json.Marshal(resp.Config.HTTPTrace.Stats())
	fmt.Println(resp.Config.HTTPTrace.Stats())
	fmt.Println(string(buf))
	fmt.Println(resp.Config.HTTPTrace.Protocol)
	fmt.Println(resp.Status)
	fmt.Println(string(resp.Data))

	return ctx.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}

func main() {

}
