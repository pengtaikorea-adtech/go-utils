# github.com/pengtaikorea-adtech/go-utils/secure

PTK go-based common security utilities

## GenerateKey

 Generate random Key string, hex

 ```go
	var key string = GenerateKey()
 ```



## Sha256

 Build SHA256 hash checksum sequence string

 ```go
	var checksum string = Sha256(input string)
```


## AESEncrypt

 Crypt input text. return base64 (RawURL) encoded string

 ```go
	var crypted string = AESEncrypt(key string, plain string)
 ```

## AESDecrypt

 Decrypt. input base64 (RawURL) encoded string, returns plain text

 ```go
	var plain string = AESDecrypt(key string, crypted string)
 ```