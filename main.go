package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Storage struct {
	Lang  string            `json:"lang"`
	Snipt map[string]string `json:"snipt"`
}

var Store []Storage

func load() []Storage {
	dataByte, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}
	var dataJson []Storage
	if err = json.Unmarshal(dataByte, &dataJson); err != nil {
		return nil
	}

	return dataJson
}
func Save(s []Storage) {
	dataByte, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	if err = os.WriteFile("data.json", dataByte, os.FileMode(os.O_CREATE|os.O_APPEND)); err != nil {
		log.Fatal(err)
	}

}
func init() {
	Store = load()
}
func main() {
	reader := bufio.NewScanner(os.Stdin)

	CMD := cobra.Command{Use: "snapcode"}
	CMD.AddCommand(&cobra.Command{
		Use:   "greet",
		Short: "testing",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello !")
		},
	})
	defer func() {

		if err := CMD.Execute(); err != nil {
			log.Fatal(err)
		}
		Save(Store)
	}()
	CMD.PersistentFlags().String("lang", "py", "default language is py")
	CMD.AddCommand(&cobra.Command{
		Use:   "add",
		Short: "add the code snipt",
		Run: func(cmd *cobra.Command, args []string) {
			lang, err := cmd.Flags().GetString("lang")
			if err != nil {
				log.Fatal(err)
			}
			var title string
			fmt.Println("Enter the title of your code snipt : ")
			fmt.Scan(&title)
			fmt.Println("Enter the code : ")
			code := ""
			for {
				text := ""
				if reader.Scan() {
					text = reader.Text()
				}
				if strings.Contains(text, "end") {
					break
				}
				code += text + "\n"
			}
			snipt := make(map[string]string)
			snipt[title] = code
			Store = append(Store, Storage{Lang: lang, Snipt: snipt})
			fmt.Println("Success")
		},
	})
	CMD.AddCommand(&cobra.Command{
		Use:   "show",
		Short: "shows all the snipts",
		Run: func(cmd *cobra.Command, args []string) {
			Bubbletea()
		},
	})
}
