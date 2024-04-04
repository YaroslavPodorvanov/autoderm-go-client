# Autoderm Go Client
Unofficial API client for https://autoderm.ai/

## Documentation
- https://autoderm.ai/docs/quickstart.html
- https://autoderm.ai/docs/redoc.html
- https://autoderm.firstderm.com/documentation/examples/
- https://autoderm.firstderm.com/docs

## Requirements
- Go 1.21 or newer

## Installation

```sh
go get github.com/YaroslavPodorvanov/autoderm-go-client
```

## Usage
```go
package main


import (
    "fmt"
    "log"
    "os"

    autoderm "github.com/YaroslavPodorvanov/autoderm-go-client"
)

func main() {
    client := autoderm.NewClient("secret_dev_team")
    
    file, err := os.Open("skin.jpeg")
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()

    response, err := client.Query("skin.jpeg", file)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(response)
}
```
```json
{
  "success": true,
  "message": "completed",
  "id": "00000000-0000-0000-0000-000007b030ce",
  "predictions": [
    {
      "confidence": 0.20118004282871116,
      "icd": "D18.01",
      "name": "Angioma (Cherry angioma)",
      "classificationId": "3e51ccc4-d4aa-11e7-a562-0242ac120003",
      "readMoreUrl": "https://www.firstderm.com/angioma/"
    },
    {
      "confidence": 0.15057508063238972,
      "icd": "D22.9",
      "name": "Atypical Melanocytic Nevus",
      "classificationId": "b698ea32-8f39-4bb5-82e5-082a58f013d7",
      "readMoreUrl": "https://www.firstderm.com/mole-congenital-nevus/"
    },
    {
      "confidence": 0.07693463345465025,
      "icd": "D22.9",
      "name": "Nevus (Benign Mole)",
      "classificationId": "3e4fd5a6-d4aa-11e7-a562-0242ac120003",
      "readMoreUrl": "https://www.firstderm.com/mole-congenital-nevus/"
    },
    {
      "confidence": 0.05560834131992198,
      "icd": "D23.9",
      "name": "Dermatofibroma",
      "classificationId": "3e4fdae1-d4aa-11e7-a562-0242ac120003",
      "readMoreUrl": "https://www.firstderm.com/dermatofibroma/"
    },
    {
      "confidence": 0.05559970467420187,
      "icd": "L81.0",
      "name": "Postinflammatory Hyperpigmentation",
      "classificationId": "3e51595d-d4aa-11e7-a562-0242ac120003",
      "readMoreUrl": "https://www.firstderm.com/hyperpigmentation/"
    }
  ]
}
```
