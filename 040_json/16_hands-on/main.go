package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type status struct {
	Code    int    `json:"Code"`
	Descrip string `json:"Descrip"`
}
type statuses []status

func main() {
	rcvd := `[{"Code":200,"Descrip":"StatusOK"},{"Code":301,"Descrip":"StatusMovedPermanently"},{"Code":302,"Descrip":"StatusFound"},{"Code":303,"Descrip":"StatusSeeOther"},{"Code":307,"Descrip":"StatusTemporaryRedirect"},{"Code":400,"Descrip":"StatusBadRequest"},{"Code":401,"Descrip":"StatusUnauthorized"},{"Code":402,"Descrip":"StatusPaymentRequired"},{"Code":403,"Descrip":"StatusForbidden"},{"Code":404,"Descrip":"StatusNotFound"},{"Code":405,"Descrip":"StatusMethodNotAllowed"},{"Code":418,"Descrip":"StatusTeapot"},{"Code":500,"Descrip":"StatusInternalServerError"}]`
	var data statuses
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(data)
	for _, v := range data {
		fmt.Println(v.Code, "=>", v.Descrip)
	}
}