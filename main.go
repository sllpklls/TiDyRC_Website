package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

func getStringHTMLItemTradeBinded(path string, data []map[string]interface{}) string {
	htmlFilePathItemTrade := path
	htmlContentItemTrade, err := ioutil.ReadFile(htmlFilePathItemTrade)
	if err != nil {
		log.Fatalf("Failed to read HTML file: %s", err)
	}
	htmlString := string(htmlContentItemTrade)

	tmpl, err := template.New("index").Parse(string(htmlString))
	if err != nil {
		log.Fatalf("Failed to create template: %s", err)
	}

	var ListItem1 strings.Builder

	for _, item := range data {
		var renderedHTML strings.Builder
		err = tmpl.Execute(&renderedHTML, item)
		if err != nil {
			log.Fatalf("Failed to execute template: %s", err)
		}
		ListItem1.WriteString(renderedHTML.String())
	}

	return ListItem1.String()
}
func convertJSONtoMap(path string) ([]map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var data []map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func saveDataToFile(filePath string, data []map[string]interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

// Hàm để thêm dữ liệu mới vào listItem2.json
func appendDataToJSONFile(filePath string, newData map[string]interface{}) error {
	// Đọc dữ liệu hiện có từ file JSON
	var data []map[string]interface{}
	file, err := os.Open(filePath)
	if err == nil {
		defer file.Close()
		byteValue, _ := ioutil.ReadAll(file)
		json.Unmarshal(byteValue, &data)
	}

	// Thêm dữ liệu mới
	data = append(data, newData)

	// Ghi lại dữ liệu vào file
	return saveDataToFile(filePath, data)
}

func main() {
	router := gin.Default()
	router.Static("/image", "./templates/image")
	router.Static("/logo", "./templates/logo")
	router.Static("/assets", "./templates/assets")
	router.Static("/html", "./templates/html")
	router.Static("/data", "./data/json")

	router.GET("/", func(c *gin.Context) {
		dataListItem1, err := convertJSONtoMap("./data/json/listItem1.json")
		if err != nil {
			return
		}
		dataListItem2, err := convertJSONtoMap("./data/json/listItem2.json")
		if err != nil {
			return
		}

		//htmlFilePathItemTrade := "./templates/html/item_trade.html"

		htmlFilePath := "./templates/vcl.html"
		htmlContent, err := ioutil.ReadFile(htmlFilePath)
		if err != nil {
			log.Fatalf("Failed to read HTML file: %s", err)
		}
		htmlString := string(htmlContent)
		data := map[string]interface{}{
			"Title":          "TiDyRC Official",
			"LogoBrand":      "./logo/logo.jpg",
			"ImageHomePage1": "./image/anh3.jpg",
			"ImageHomePage2": "./image/anh1.jpg",
			"ImageHomePage3": "./image/anh2.jpg",
			"EmailContact":   "hoangthai.actvn@gmail.com",
			"TradingBoard":   "./image/stage1/TradingBoard.jpg",
			"BoardsDiagram":  "./image/stage1/BoardsDiagram.jpg",
			"CommunityChat":  "./image/stage1/CommunityChat.jpg",
			"location":       "167 Tay Son, Dong Da, Ha Noi",
			"ShareExp":       "./image/stage1/ShareExp.jpg",
			"LinkMessage":    "https://www.messenger.com/t/sllpklls",
			"LinkFacebook":   "https://www.facebook.com/profile.php?id=61557299456770",
			"LinkTikTok":     "https://www.tiktok.com/@sllpklls_fpv?is_from_webapp=1&sender_device=pc",
			"LinkYoutube":    "https://www.youtube.com/@thaihoang7661",
			"imageCross2":    "./image/stage2/sodomach.jpg",
			"ListItem1":      getStringHTMLItemTradeBinded("./templates/html/item_trade.html", dataListItem1),
			"ListItem2":      getStringHTMLItemTradeBinded("./templates/html/item_trade.html", dataListItem2),
		}

		tmpl, err := template.New("index").Parse(htmlString)
		if err != nil {
			log.Fatalf("Failed to create template: %s", err)
		}

		var renderedHTML strings.Builder
		err = tmpl.Execute(&renderedHTML, data)
		if err != nil {
			log.Fatalf("Failed to execute template: %s", err)
		}
		c.Data(200, "text/html; charset=utf-8", []byte(renderedHTML.String()))
	})

	router.Run(":8080")
}
