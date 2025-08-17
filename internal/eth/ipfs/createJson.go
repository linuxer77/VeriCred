package ipfs

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// func CreateJSONFileAndStoreToIPFS() {

// 	f, err := os.CreateTemp("/home/linuxer77/temp-files", "tempfile-*.json")
// 	if err != nil {
// 		fmt.Println("Error when calling CreateTemp:", err)
// 		return
// 	}

// 	defer f.Close()
// 	fmt.Println(f.Name())

// 	// bodyBytes, err := io.ReadAll(r.Body)

// 	// if err != nil {
// 	// 	fmt.Println("Error in calling ReadAll():", err)
// 	// 	return
// 	// }

// 	bodyBytes := []byte("Hell")

// 	_, err = f.Write(bodyBytes)
// 	if err != nil {
// 		fmt.Println("Error when writing to data to file:", err)
// 		return
// 	}

// }


func CreateJSONFileAndStoreToIPFS(w http.ResponseWriter, r *http.Request) {
	
	f, err := os.CreateTemp("/home/linuxer77/temp-files", "tempfile-*.json")
	if err != nil {
		fmt.Println("Error when calling CreateTemp:", err)
		return
	}

	defer f.Close()
	defer os.Remove(f.Name())

	bodyBytes, err := io.ReadAll(r.Body)

	fmt.Printf("Received request body: %s\n", bodyBytes)
	
	if err != nil {
		fmt.Println("Error in calling ReadAll():", err)
		return
	}

	_, err = f.Write(bodyBytes)
	if err != nil {
		fmt.Println("Error when writing to data to file:", err)
		return
	}

	msg, err := UploadToIPFS(f.Name())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(msg)
}