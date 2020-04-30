package main

import (
	"fmt"
	"github.com/jsterling7/go_proto_tutorial/spec"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

const fname = "addressbook"


func main() {
	writeAddressBook()
	readAddressBook()
}


func writeAddressBook() {
	josh := spec.Person{
		Name:        "Joshua Sterling",
		Id:          1234,
		Email:       "email",
		Phones:      []*spec.Person_PhoneNumber{{
			Number: "123-456-7891",
			Type:   spec.Person_MOBILE,
		}},
		LastUpdated: nil,
	}

	bill := spec.Person{
		Name:        "Bill Wander",
		Id:          1234,
		Email:       "email",
		Phones:      []*spec.Person_PhoneNumber{{
			Number: "123-456-5410",
			Type:   spec.Person_MOBILE,
		}},
		LastUpdated: nil,
	}

	addressBook := spec.AddressBook{People: []*spec.Person{&josh, &bill}}


	// Write the address book to disk
	out, err := proto.Marshal(&addressBook)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}

func readAddressBook() {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	addressBook := &spec.AddressBook{}
	if err := proto.Unmarshal(in, addressBook); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Println(addressBook)
}