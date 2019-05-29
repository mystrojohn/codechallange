package main

import (
	"fmt"
	"os"
	"crypto/rand"
    "crypto/rsa"
	"crypto/x509"
	"crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
	"encoding/pem"
	"encoding/asn1" 
   "bufio"
   "io/ioutil"

)



//Method to store the RSA keys in public key Format -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY-----
func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) {  
	asn1Bytes, err := asn1.Marshal(pubkey)  
   if err != nil {  
      fmt.Println("Fatal error ", err.Error())  
      os.Exit(1)  
   }  
   var pemkey = &pem.Block{  
      Type:  "PUBLIC KEY",  
      Bytes: asn1Bytes,  
   }  
   
   pemfile, err := os.Create(fileName)  
   if err != nil {  
      fmt.Println("Fatal error ", err.Error())  
      os.Exit(1)  
   }  
   defer pemfile.Close()  
   err = pem.Encode(pemfile, pemkey)  
   if err != nil {  
      fmt.Println("Fatal error ", err.Error())  
      os.Exit(1)  
   }  
}


//Load public key stored in PEM format and create RSA public Key using go x509.ParsePKCS1PublicKey
func loadPublicPemKey (fileName string)  (*rsa.PublicKey) {  
  
   publicKeyFile,err := os.Open(fileName)  
   if err != nil {  
      fmt.Println("Fatal error ", err.Error())  
      os.Exit(1)  
   }  
  
   pemfileinfo, _ := publicKeyFile.Stat()  
  
   size  := pemfileinfo.Size()  
   pembytes := make([]byte, size)  
   buffer := bufio.NewReader(publicKeyFile)  
   _, err = buffer.Read(pembytes)  
   data, _ := pem.Decode([]byte(pembytes))  
   publicKeyFile.Close()  
   
   publicKeyFileImported, err := x509.ParsePKCS1PublicKey(data.Bytes)  
   if err != nil {  
      fmt.Println("Fatal error ", err.Error())  
      os.Exit(1)  
   }  
   return publicKeyFileImported  
}

// Exists checks whether the named file or directory exists.
func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

func CryptoSha256(message string, secret string) string {
    key := []byte(secret)
    h := hmac.New(sha256.New, key)
    h.Write([]byte(message))
    return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func main() {
	args := os.Args[1]
	fileExists := false


	// Exists checks whether the file exists.
	if (!Exists("pubkey.pem")) {
		if len(os.Args) > 1 {
			args = os.Args[1]
		} else {
			fmt.Println("Please enter your email address! Goodbye....")  
			os.Exit(0)    
		} 
	
	} else {
		fileExists = true
		loadPublicPemKey("pubkey.pem")
	}


	sh256 := CryptoSha256(args, "secret")

	fmt.Println(sh256)
	
	// Generate John RSA keys Of 2048 
   johnPrivateKey, err := rsa.GenerateKey(rand.Reader, 4096)  
   if err != nil {  
	  //fmt.Println(err.Error)  
	  os.Exit(1)  
   }  
   // Extract Public Key from RSA Private Key  
   johnPublicKey := johnPrivateKey.PublicKey  
   //fmt.Println("John Private Key :  ", *johnPrivateKey)  
   fmt.Println("John Public key ", johnPublicKey)

   // save private key
   pkey := x509.MarshalPKCS1PrivateKey(johnPrivateKey)
   ioutil.WriteFile("private.key", pkey, 0777)
   fmt.Println("private key saved to private.key")
   
   // save public key
   pubkey, _ := x509.MarshalPKIXPublicKey(johnPublicKey)
   ioutil.WriteFile("public.key", pubkey, 0777)
   fmt.Println("public key saved to public.key")


//	   Once the RSA Key pair is generated store it in PEM format with pkcs encoding.
	
	// Method to Store in Public Key format 
	if (!fileExists) {
		savePublicPEMKey("pubkey.pem",johnPublicKey)
	}


   //fmt.Println(args)
   fmt.Println(`{`) 
	fmt.Println(`"message":"`, args, `",`)
   fmt.Println(`"signature":"`, sh256, `",`)
	fmt.Println(`"pubkey":"-----BEGIN PUBLIC KEY-----\n`,johnPublicKey,`\n-----END PUBLIC KEY-----\n"`)
	fmt.Println(`}`) 
}
