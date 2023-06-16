package main

import (
	"testing"

	"github.com/GA-Marketing/service-viability/helpers"
)

func TestCryptAndVerify(t *testing.T) {
	stringToEncrypt := "umastring234"
	encripted := string(helpers.Encrypt(stringToEncrypt))

	if !helpers.EncryptVerify(stringToEncrypt, encripted) {
		t.Error("Verificação da string falhou")
	}
}

func TestCryptAndVerifyWithFail(t *testing.T) {
	stringToEncrypt := "umastring123"
	encripted := string(helpers.Encrypt(stringToEncrypt))

	if helpers.EncryptVerify(encripted, stringToEncrypt+"k") {
		t.Errorf("String %s != de %s", encripted, helpers.Encrypt(stringToEncrypt+"k"))
	}
}
