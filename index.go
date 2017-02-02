package main

import(
	"gopkg.in/olivere/elastic.v5"
	"golang.org/x/net/context"
	"log"
	"os"
	//"fmt"
	"encoding/json"
)

type people struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Institution string `json:"institution"`
	Email string `json:"email"`
}

type item struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Journal  string `json:"journal"`
	Volume int `json:"volume"`
	Number int `json:"number"`
	Pages string `json:"pages"`
	Year int `json:"year"`
	Authors []people `json:"authors"`
	Abstract string `json:"abstract"`
	Link string `json:"link"`
	Keywords []string `json:"keywords"`
	Body string `json:"body"`
}

var client *elastic.Client
var err error
func init(){
	client,err = elastic.NewClient()
	if err!=nil{
		log.Fatal(err)
	}
}

func main() {
	var data []item
	
	file,err := os.Open("data.json")
	if err!=nil{
		log.Fatal(err)
	}
	defer file.Close()
	
	jsonDeocder := 	json.NewDecoder(file)
	if err := jsonDeocder.Decode(&data); err!=nil{
		log.Fatal("Decode: ",err)
	}
	bulkIndex("library","article",data)
}

func bulkIndex(index string,typ string ,data []item){
	ctx := context.Background()
	for _,item := range data{
		_,err := client.Index().Index(index).Type(typ).Id(item.Id).BodyJson(item).Do(ctx)	
		if err !=nil{
			log.Fatal(err)
		}
	}
}