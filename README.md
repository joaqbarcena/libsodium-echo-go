# libsodium-echo-go
A Libsodium encryptor wrapper in Go and more


## Installation 

#### Unix-like or Windows
- First you need to [install libsodium](https://libsodium.gitbook.io/doc/installation)

#### MacOS
- In case of running MacOS you could download from homebrew
```bash
brew install libsodium
```

Then clone it and run
```bash
git clone https://github.com/joaqbarcena/libsodium-echo-go && cd libsodium-echo-go
go run src/api/main.go
```
It runs over the 8080 port, so be sure to keep it clean

## Usage

There are 2 endpoints

#### /encrypt
This uses a `public_key` and `text` parameters (expected as query params) to encrypt the given text with your public key
Example
```curl
curl --location --request GET 'localhost:8080/encrypt?public_key=RlRwWuxn8Pm5caD6fk02HXtCkDFMgXoNfAmZX1hbJm0=&text=que%20miras'
```
returns a json
```json
{
    "encrypted": "PdzZk3XC/j9giK/gX9LPf8LmKIW8r8LI0YB07x2A7BcdSE2cZDPMXCrxHaUvGg7gvv+Oox22ozlk"
}
```

#### /native/mock
This uses a `public_key` parameter (expected as query param) to encrypt any ocurrence of text between `<<` `>>` delimiters from [response.json](src/api/resources/response.json)

Example
```curl
curl --location --request GET 'localhost:8080/native/mock?public_key=RlRwWuxn8Pm5caD6fk02HXtCkDFMgXoNfAmZX1hbJm0='
```
returns a json
```json
{
    "plain_text": "hola che gil",
    "encrypted_custom_object": {
        "encrypted_text": "OWZhu2CdBAg486ZCkiCY3GINYqXXu/Fdw/SB2F0tbFRM6DDqo3Jw2yU1YCScw1LcRuZlFLNgwqQXcNQZ"
    }
}
```
You can change the content of `response.json`, and give any mock json, where any occurrence of text inside `<<` `>>` delimiters wil be encrypted
Example `response.json`
```json
{
    "myencrypted_custom_answer": "<<Im a encrypted Text !>>",
    "another_encrypted_custom_answer": "<<Im another encrypted Text !>>",
    "common_text": "Im a encrypted Text !"
}
```
will result in 
```json
{
    "myencrypted_custom_answer": "ORJAoV/WanwLvpxHQPgYu92KpAFVIfYPSuRiq0j+SkCXnv8AUkI/emygizuiYIPVaA0N52MLX0HiorVl3RzjaM4cexwT",
    "another_encrypted_custom_answer": "g1s1u64fhudeqI1q0dyZVzVyHdazaZWOwUkfM6KLHFEV/56sKESYfB4qr3AbBvvr/mm6/h/tvWGrZnb4oO9s4STVUb9nvEY4oo9Q",
    "common_text": "Im a encrypted Text !"
}
```
Observe that **any** ocurrence will encrypt.

#### /keypairgen
Endpoint to obtain a secret key and its corresponding public key.

```curl
curl --location --request GET 'localhost:8080/keypairgen'
```

Response

```json
{
    "secret_key": "94alpEfoJ34wPGl3j8XQvpVl3a0S+mge/tsr1w76cBQ=",
    "public_key": "HQO6x27IpgnM/kExXLA62+c+rGguR5R8de7BxRum2G4="
}
```

You can optionally pass a seed as a query param, in order to obtain the key pair deterministically.
This seed should be base64 encoded.

```curl
curl --location --request GET 'localhost:8080/keypairgen?seed=2y3J9Te/e4BDaakrsRaQHKcVBMQ9lFfajmXNceH6Hiw='
```

Response

```json
{
    "secret_key": "3eqLv3JxMq+3eQ90syLHtjXaZtuuwlDTzUJ0qwIyB6o=",
    "public_key": "hbIsylPj0a/yiRrw9sipHFRATERqqdITu2CfyxYO/DU="
}
```
