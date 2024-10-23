package baidutranslatefree

import (
	"fmt"
	"testing"
)

func TestSigner_CalSign(t *testing.T) {
	signer, err := NewSigner()
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	sign, err := signer.CalSign("a cartoon of a woman with blue hair")
	if err != nil || sign != "889305.619240" {
		t.Error(err)
	}
	sign, err = signer.CalSign("Hello,Word!")
	if err != nil || sign != "524688.844449" {
		t.Error(err)
	}

}
