package baidutranslatefree

import (
	"fmt"
	"log"
	"testing"
)

func TestTranslate(t *testing.T) {
	translater, _ := NewBaiduTranslater()
	translater.SetFromTo("en", "zh")
	text, _ := translater.Translate("Hello,World!")
	fmt.Printf("%v\n", text)
	text, _ = translater.Translate("a cartoon of a woman with blue hair")
	fmt.Printf("%v\n", text)
}

func TestGetIDAndToken(t *testing.T) {
	baiduid, token, err := GetIDAndToken()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("BAIDUID:", baiduid)
		fmt.Println("Token:", token)
	}
}

func TestCalSign(t *testing.T) {
	signer, err := NewSigner()
	if err != nil {
		log.Fatalf("Error creating signer: %v", err)
	}

	signature, err := signer.CalSign("Hello,World!")
	if err != nil {
		log.Fatalf("Error calculating sign: %v", err)
	}
	//791002.586475
	log.Printf("Signature: %s", signature)
}
