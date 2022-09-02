package jsonwebtoken_test

/*
import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"{{ .ProjectName }}/src/shared/jsonwebtoken"
	"{{ .ProjectName }}/src/shared/test/assert"
)

const (
	USER_ID     = "55157b04-e41d-414a-93d6-f55d43cb8f05"
	PRIVATE_KEY = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQ0KTUlJRXZRSUJBREFOQmdrcWhraUc5dzBCQVFFRkFBU0NCS2N3Z2dTakFnRUFBb0lCQVFEWEkwWUZaSlpZN3psdA0KN0FUZXQzcUJudEwybjlZL3MraTJKMTc1elhYTHZtZ1NoV1R1K0JYd1JIRjIyenp3TjRUem9CVzVyelMwZlNZWg0KcTJqRHJLajlBdjJFMmJVSkxqVU5NTWhZRXdoWUNhemRFMFV3SDJtamdWTFJSQU1XMEJDRjhQZkpDemZISlV2Yw0KZSsvZU5ERkx0RmpLRUhQTkZweDBjY1FFaXlqK3JzWm9LNWg0a25IOXlqbjdERzdsTEhOa3ZKaE4xdUsxMGFGSg0Kdi9LSEp3bFF1VFh6ODZvR05YT1BMZmJZYXRlb2pMbzJSRitRN0FTNUFORmdXOGJhTTNmaTBpQ2oyenVXYWNIbw0KZjVnUEhsYnRPWWg0NkpMNUZPcVdPT0MrUWFRZWx6a2pQY2ErM1JXVDBXVnRGS3VuMi8yVmhmeUxySzJUaE9taQ0KNzUxbEQyQTNBZ01CQUFFQ2dnRUFIRXVQM0lxM3VYZXhNbk5keXBzUUdqYzA0T254eUV3VnVjdGcwdk1DWUtzbw0KemZMTWJSK2s1d1poYy9QdXhsdE8rYW5lNFo2WkRIbi9SbFlFZitQWUQvclNOQ1JjQ3BxUXNLenZWS0lnTVZBSQ0KYzhVaWp1czZ2aDMyc2Y1RGQ3OGl4VE1GcStQMVVKUm5yMVovbkZaNSswNTdkUXdINXZ2bjFPclhrSTlxODEvMA0KYTEzN0ZOeFJ6anFuZG5ZQURqNlNsWEVsbmw0Wjd4eHVubzR3SmRJVlREVGNmb2dDR2tBbzlHTmpRUDJnRkhBZw0KRkVDZEpXcDFZQlUrc1QxT0t5K1lPTVdMaUFQQWZGRXNUK2FMTEdiSktqUFRRM1Zlamtxd3BaeU52eTl4Nk5mVQ0Kb3hvekpUQ0wzUjhxcE13bUlsMWllT3hqNnh4TVFpTEVHNHBLZFc1UlNRS0JnUUQ2dlV4TUs3dzBLcktwRHMzSg0KWitXdUpjbmVwL1B6cXdtT0ozTDhvakpMOU1SUnNiN1N1Q3QrRUh3RWlzalQwNVd4bkp5b0M3YlozM3NNVXFRUA0KTC9NeWFnekJScHUyL1pVVzV4Z1BBSmxtOFJFem51UmtXZ1p5Q09wNk9JVEp0VkhIN1BlVm5BSUN5aVArTUcwbQ0KWFlLV2Nya3oyUzg5OGhtY1FpQ3gzclpwQ1FLQmdRRGJwc01Dbm5EYTg1WE0rUmRTc1cwUEhBNzdwL0d2alVYYw0KMTdNZmNjMUhreW9VRzlHblVROTk3NHVaZFlKRGNwZVhnWURUM2tLa2dBU0pVOEFqNEtzdCs2TDB1V241WjhxSg0KVXNsNElSa1FrUHdySFZQaHhIeDNrQnRoUWlaQmlVZmtvT3JQN2xjZFFJbEhzazE0MUxFVkdGbCtMZjNOeEZpQw0KZk5rakN1UVBQd0tCZ0VmaW40NHk5N2twQ0tHcFNkeFZPaWNNVkd6T1VBVnE0c2xyaSs3Yi9YbURTZ2wxNFJSKw0Ka3BHSTNsVmJDS3FhUFk5M0svNk5wVFdmZWFLVnlzMUQzUUIySVFRRVh6NCtRMUVXbmZJbkpOTzdoMGY2Wk5aVQ0KYlFhaWdiN2FsMDlRK1lwTTZNcHV1TERlRFNXaDhwa09OQk0zL3RyYmlFekZMUXg4ZE8wcHdiZ3BBb0dCQUlqRg0KUmhpVFgrSjJXb2pQY2Q1ajdHekVJL0EvbXhhYytTdVNoRTdJSmZLemlEZ05PbmJjMnJDb1FGekY5dDdZczl2Nw0KSDZUTmVPSEZkUTJ0d0s2V2J3Q3E5OFU2enVvbDNzK3paUkFRUy9NczFGaGtZcDdxSWphdzNOdXF2UGVCNitwSQ0KNk8yZ2swMzdxWCtqWHVvbVJqM0VjN1ZHSGd2S2Z2S0JteE5lN0xNeEFvR0FXV2lDVmM5bVZoYjBuajY3RExtSw0KWXdHMTM5aEFpdjNxcnZjZ1IvY3VES3hxcHRRdjJnaTRYODFLc0gwSzNHNkZGRFgxNG1yamZTU29ITjdVK2ZWbg0KN1FsREEwNW1sK0sxQi9rayt1RUpTbnJTTVRITWNnc2lwVnNlZi9rUi9hZ0xoTG9HUG56dG8vZVJEbHVwSWVaZw0KSHRaYTZFcmV5Y1diWEkvcmN3M1F6elU9DQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ0K"
	PUBLIC_KEY  = "LS0tLS1CRUdJTiBSU0EgUFVCTElDIEtFWS0tLS0tDQpNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQTF5TkdCV1NXV084NWJld0UzcmQ2DQpnWjdTOXAvV1A3UG90aWRlK2MxMXk3NW9Fb1ZrN3ZnVjhFUnhkdHM4OERlRTg2QVZ1YTgwdEgwbUdhdG93NnlvDQovUUw5aE5tMUNTNDFEVERJV0JNSVdBbXMzUk5GTUI5cG80RlMwVVFERnRBUWhmRDN5UXMzeHlWTDNIdnYzalF4DQpTN1JZeWhCenpSYWNkSEhFQklzby9xN0dhQ3VZZUpKeC9jbzUrd3h1NVN4elpMeVlUZGJpdGRHaFNiL3loeWNKDQpVTGsxOC9PcUJqVnpqeTMyMkdyWHFJeTZOa1Jma093RXVRRFJZRnZHMmpOMzR0SWdvOXM3bG1uQjZIK1lEeDVXDQo3VG1JZU9pUytSVHFsampndmtHa0hwYzVJejNHdnQwVms5RmxiUlNycDl2OWxZWDhpNnl0azRUcG91K2RaUTlnDQpOd0lEQVFBQg0KLS0tLS1FTkQgUlNBIFBVQkxJQyBLRVktLS0tLQ0K"
)

func TestGenerateTokenExpiresInOneDay(t *testing.T) {
	t.Skip()
	privateKey, err := getRsaPrivateKeyFrom(PRIVATE_KEY)
	assert.Nil(t, err)

	publicKey, err := getRsaPublicKeyFrom(PUBLIC_KEY)
	assert.Nil(t, err)

	jwt := jsonwebtoken.New(jsonwebtoken.Config{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Audience:   "web",
		Expiration: "86400",
	})
	token, err := jwt.GenerateToken(USER_ID)
	assert.Nil(t, err, fmt.Sprintf("error: %s", err))
	assert.True(t, len(token) > 0)

	id, err := jwt.ParseToken(token)
	assert.Nil(t, err)
	assert.Equal(t, id, USER_ID)
}

func TestGenerateExpiredToken(t *testing.T) {
	t.Skip()
	privateKey, err := getRsaPrivateKeyFrom(PRIVATE_KEY)
	assert.Nil(t, err)

	publicKey, err := getRsaPublicKeyFrom(PUBLIC_KEY)
	assert.Nil(t, err)

	jwt := jsonwebtoken.New(jsonwebtoken.Config{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Audience:   "web",
		Expiration: "0",
	})
	token, err := jwt.GenerateToken(USER_ID)
	assert.Nil(t, err)
	assert.True(t, len(token) > 0)

	id, err := jwt.ParseToken(token)
	assert.NotNil(t, err)
	assert.Equal(t, id, "")
}

func getRsaPrivateKeyFrom(val string) (*rsa.PrivateKey, error) {
	privkeyBytes, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return nil, err
	}
	r := strings.NewReader(string(privkeyBytes))
	pemBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(pemBytes)

	privkey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privkey.(*rsa.PrivateKey), nil
}

func getRsaPublicKeyFrom(val string) (*rsa.PublicKey, error) {
	pubkeyBytes, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return nil, err
	}
	r := strings.NewReader(string(pubkeyBytes))
	pemBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(pemBytes)
	pubkey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pubkey.(*rsa.PublicKey), nil
}
*/
