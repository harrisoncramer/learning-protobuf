package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/harrisoncramer/learning-protobuf/model"
	"google.golang.org/protobuf/proto"
)

func main() {

	/* Create some book data */
	book := &model.Book{
		Id:       1,
		Title:    "Harry Potter",
		Author:   "JK Rowling",
		Category: model.Category_Novel, /* The Category_Novel field is auto-generated for us */
	}

	/* Encode the data into a byte stream using the proto library */
	data, err := proto.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}

	/* Write the data to a file. You could also send it across the wire */
	err = ioutil.WriteFile("book.protobuf", data, 0600)

	/* Read the file */
	data, err = ioutil.ReadFile("book.protobuf")
	if err != nil {
		log.Fatal(err)
	}

	/* Decode the data back into our struct */
	bookFromFile := model.Book{}
	err = proto.Unmarshal(data, &bookFromFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("book from protobuf file %+v\n", bookFromFile)

}
