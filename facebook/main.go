package main

import (
	"flag"
	"fmt"

	"github.com/huandu/facebook"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/motivation-machine/motivator/facebook/models"
	"github.com/spf13/viper"
)

func main() {
	configFilePath := flag.String("config", "config.json", "Config file path")

	// Load configuration
	viper.SetConfigFile(*configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to load configuration %v", err))
	}
	viper.AutomaticEnv()

	// DB connection
	db, err := gorm.Open("postgres",
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			viper.Get("FACEBOOK_DB_USER"),
			viper.Get("FACEBOOK_DB_PASSWORD"),
			viper.Get("FACEBOOK_DB_HOST"),
			viper.Get("FACEBOOK_DB_NAME")))
	if err != nil {
		panic(fmt.Errorf("Failed to connect to database %v", err))
	}
	defer db.Close()

	// Migration
	db.Model(&models.Source{}).Related(&models.Result{})
	db.AutoMigrate(&models.Source{}, &models.Result{})
	sources := models.InsertSources(db)

	var hasMore = true

	// Facebook session init
	app := facebook.New(viper.GetString("facebookAppId"), viper.GetString("facebookAppSecret"))
	session := app.Session(viper.GetString("facebookAccessToken"))
	err = session.Validate()
	if err != nil {
		panic(fmt.Sprintf("Failed to validate session %v", err))
	}

	// Get Posts from Source and save them
	for _, source := range sources {
		fmt.Printf("Started source ID: %d UserName: %s \n", source.ID, source.UserName)
		results, err := session.Get(
			fmt.Sprintf("/%s/posts?fields=type,full_picture,description&type=photo&limit=100", source.UserName),
			nil,
		)
		if err != nil {
			panic(fmt.Errorf("Failed to Get posts  %v", err))
		}
		paging, _ := results.Paging(session)
		pageNum := 0
		resultNum := 0

		for hasMore {
			pageNum++
			var resultItems []facebook.Result = paging.Data()
			for _, item := range resultItems {
				var r = &models.Result{
					FbID:     item["id"].(string),
					SourceID: source.ID,
				}
				description, ok := item["description"].(string)
				if ok {
					r.Description = description
				}
				pictureRawURL, ok := item["full_picture"].(string)
				if ok {
					r.PictureRawURL = pictureRawURL
				}
				// fmt.Printf("%v", r)
				db.FirstOrCreate(r, map[string]interface{}{"fb_id": item["id"].(string)})
			}
			resultNum += len(resultItems)
			noMore, err := paging.Next()
			hasMore = !noMore
			fmt.Printf("sourceID: %d SourceUserName: %s totalDownlaoded: %d PageNum: %d HasMore: %v errPaging: %v \n", source.ID, source.UserName, resultNum, pageNum, hasMore, err)
		}
	}
}
