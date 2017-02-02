package main

import(
	//"gopkg.in/olivere/elastic.v5"
	"fmt"
	"log"
	//"net/http"
	"os"
	//"strings"
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
	Volume string `json:"volume"`
	Number string `json:"number"`
	Pages string `json:"pages"`
	Year string `json:"year"`
	Authors []people `json:"authors"`
	Abstract string `json:"abstract"`
	Link string `json:"link"`
	Keywords []string `json:"keywords"`
	Body string `json:"body"`
}
//For decoding our json file, we will use this array. Notice the type item and people.
var data []item

/*func init(){
	client,err := elastic.NewClient()
	if err!=nil{
		log.Fatal(err)
	}
}*/

func main() {
	file,err := os.Open("data.json")
	if err!=nil{
		log.Fatal(err)
	}	
	defer file.Close()
	jsonDeocder := 	json.NewDecoder(file)
	if err := jsonDeocder.Decode(&data); err!=nil{
		log.Fatal("Decode: ",err)
	}
	fmt.Println(data[0].Id)
}

//func bulkIndex(index string,type string,)