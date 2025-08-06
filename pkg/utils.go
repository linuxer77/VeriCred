package pkg

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)
var secretkey []byte
type Claims struct{
	MetamaskAddress string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}

	secretkey = []byte(secret)
}

func CreateToken(metamaskAddress string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"metamask_address": metamaskAddress,
			"exp":              time.Now().Add(time.Hour * 24).Unix(),
		})
	// The better way is to initialize the struct directly:
	var C = &Claims{MetamaskAddress: metamaskAddress}
	// Alternatively, you can declare and then assign:
	// var C Claims
	// C.MetamaskAddress = metamaskAddress
	log.Println("Successfully set the Claims ", C)

	tokenString, err := token.SignedString(secretkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretkey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	
	var C = &Claims{}
	return C.MetamaskAddress, nil
}

func GenerateNonce() string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	randomNum := rng.Intn(899999) + 100000
	return fmt.Sprintf("login request #%d", randomNum)
}

func VerifySignature(address, message, sigHex string) (bool, error) {
	log.Println("VerifySignature gets called.")
	data := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + message)
    hash :=  crypto.Keccak256(data)
	
	log.Println("Hash: ", hash)

	sig := hexToBytes(sigHex)
	log.Println("Sig: ", sig)

	if sig[64] != 27 && sig[64] != 28 {
		return false, nil
	}
	sig[64] -= 27

	log.Println("sig after opeartions: ", sig)

	pubKey, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return false, err
	}

	log.Println("pubkey: ", pubKey)

	recoveredAddr := crypto.PubkeyToAddress(*pubKey).Hex()
	log.Println("recoveredAdrr: ", recoveredAddr)
	log.Println("Sent address: ", address)
	return strings.EqualFold(recoveredAddr, address), nil
}

func hexToBytes(hexStr string) []byte {
	b, _ := hex.DecodeString(strings.TrimPrefix(hexStr, "0x"))
	return b
}