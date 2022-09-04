package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

var rootCmd = &cobra.Command{
	Use:   "gitApp",
	Short: "it's a simple app fot fetching data from github api",
	Long:  "it's a very simple app written in golang to fetch data from github rest api",
	Run: func(cmd *cobra.Command, args []string) {

		uri := "https://api.github.com/users/"
		newUrl := uri + args[0]

		response, err := http.Get(newUrl)
		if err != nil {
			log.Fatal("something went wrong")
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		var response1 Response
		if err := json.Unmarshal(responseData, &response1); err != nil {
			log.Fatal(err)
		}

		println(response1.Name)
		println(response1.Bio)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
