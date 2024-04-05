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
go get github.com/dochq/autoderm-go-client
```

## Usage
```go
package main


import (
    "fmt"
    "log"
    "os"

    autoderm "github.com/dochq/autoderm-go-client"
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

## Classifications
- Source https://autoderm.ai/v1/classifications
- Using PostgreSQL to create a table with values:
```sql
SELECT classification ->> 'id'   AS id,
       classification ->> 'name' AS name,
       classification ->> 'icd'  AS icd
FROM JSON_ARRAY_ELEMENTS('[
  {
    "id": "3e5001af-d4aa-11e7-a562-0242ac120003",
    "icd": "L82.9",
    "name": "Seborrhoeic Keratosis"
  },
  {
    "id": "3e50559e-d4aa-11e7-a562-0242ac120003",
    "icd": "L92.0",
    "name": "Granuloma Annulare"
  },
  {
    "id": "feffe040-e0d4-11e7-889a-cd86047b183d",
    "icd": "R23.1",
    "name": "Livedo Reticularis"
  },
  {
    "id": "3e4fb2bb-d4aa-11e7-a562-0242ac120003",
    "icd": "L72.0",
    "name": "Milia"
  },
  {
    "id": "3e4fba39-d4aa-11e7-a562-0242ac120003",
    "icd": "B07.9",
    "name": "Verruca Vulgaris / Wart(s)"
  },
  {
    "id": "3e51ced2-d4aa-11e7-a562-0242ac120003",
    "icd": "L70.9",
    "name": "Acne, Unspecified"
  },
  {
    "id": "3e4faeb8-d4aa-11e7-a562-0242ac120003",
    "icd": "L25.9",
    "name": "Contact Dermatitis"
  },
  {
    "id": "3e4fdfcd-d4aa-11e7-a562-0242ac120003",
    "icd": "L73.1",
    "name": "Pseudofolliculitis Barbae"
  },
  {
    "id": "3e4f975b-d4aa-11e7-a562-0242ac120003",
    "icd": "L57.0",
    "name": "Actinic Keratosis"
  },
  {
    "id": "3e519003-d4aa-11e7-a562-0242ac120003",
    "icd": "N48.1",
    "name": "Balanitis"
  },
  {
    "id": "3e50220f-d4aa-11e7-a562-0242ac120003",
    "icd": "L30.1",
    "name": "Dyshidrotic Eczema"
  },
  {
    "id": "3e4f9cbe-d4aa-11e7-a562-0242ac120003",
    "icd": "B02.9",
    "name": "Herpes Zoster"
  },
  {
    "id": "3e505f7b-d4aa-11e7-a562-0242ac120003",
    "icd": "L30.4",
    "name": "Intertrigo"
  },
  {
    "id": "3e4fad25-d4aa-11e7-a562-0242ac120003",
    "icd": "L91.0",
    "name": "Keloid"
  },
  {
    "id": "3e4fd079-d4aa-11e7-a562-0242ac120003",
    "icd": "L30.9",
    "name": "Dermatitis, UNS"
  },
  {
    "id": "3e50a7d2-d4aa-11e7-a562-0242ac120003",
    "icd": "B09",
    "name": "Viral Exanthema"
  },
  {
    "id": "3e4fc498-d4aa-11e7-a562-0242ac120003",
    "icd": "B36.0",
    "name": "Pityriasis Versicolor"
  },
  {
    "id": "3e507f48-d4aa-11e7-a562-0242ac120003",
    "icd": "L01.0",
    "name": "Impetigo"
  },
  {
    "id": "3e511f5f-d4aa-11e7-a562-0242ac120003",
    "icd": "D29.0",
    "name": "Pearly Penile Papules"
  },
  {
    "id": "3e5027c1-d4aa-11e7-a562-0242ac120003",
    "icd": "L98.0",
    "name": "Pyogenic Granuloma"
  },
  {
    "id": "3e4fa42d-d4aa-11e7-a562-0242ac120003",
    "icd": "L81.1",
    "name": "Melasma (Cloasma)"
  },
  {
    "id": "3e4fb82d-d4aa-11e7-a562-0242ac120003",
    "icd": "L50.9",
    "name": "Urticaria"
  },
  {
    "id": "3e4fc6ff-d4aa-11e7-a562-0242ac120003",
    "icd": "L40.9",
    "name": "Psoriasis"
  },
  {
    "id": "3e501db1-d4aa-11e7-a562-0242ac120003",
    "icd": "L21.9",
    "name": "Seborrhoeic Dermatitis"
  },
  {
    "id": "607570e0-11f1-11e8-9c1e-73fd0e230291",
    "icd": "T69.1",
    "name": "Chilblains (pernio)"
  },
  {
    "id": "3e4fbf6b-d4aa-11e7-a562-0242ac120003",
    "icd": "D22.9",
    "name": "Dermal Nevus"
  },
  {
    "id": "3e4fd5a6-d4aa-11e7-a562-0242ac120003",
    "icd": "D22.9",
    "name": "Nevus (Benign Mole)"
  },
  {
    "id": "3e4fa102-d4aa-11e7-a562-0242ac120003",
    "icd": "I83.1",
    "name": "Eczema Hypostaticum"
  },
  {
    "id": "3e4f935a-d4aa-11e7-a562-0242ac120003",
    "icd": "L02.91",
    "name": "Abscess"
  },
  {
    "id": "3e4fa9b7-d4aa-11e7-a562-0242ac120003",
    "icd": "A60.0",
    "name": "Genital Herpes"
  },
  {
    "id": "3e50c7c4-d4aa-11e7-a562-0242ac120003",
    "icd": "L20.8",
    "name": "Neurodermatitis (Lichen Simplex Chronicus)"
  },
  {
    "id": "3e509f1c-d4aa-11e7-a562-0242ac120003",
    "icd": "L95.9",
    "name": "Vasculitis"
  },
  {
    "id": "3e4fdd6d-d4aa-11e7-a562-0242ac120003",
    "icd": "C43.9",
    "name": "Malignant Melanoma"
  },
  {
    "id": "3e50f0e3-d4aa-11e7-a562-0242ac120003",
    "icd": "R23.3",
    "name": "Petechia Or Purpura"
  },
  {
    "id": "3e4fb3b9-d4aa-11e7-a562-0242ac120003",
    "icd": "L71.9",
    "name": "Rosacea"
  },
  {
    "id": "3e511e50-d4aa-11e7-a562-0242ac120003",
    "icd": "L90.0",
    "name": "Lichen Sclerosus et Atrophicus"
  },
  {
    "id": "3e5154d1-d4aa-11e7-a562-0242ac120003",
    "icd": "B35.4",
    "name": "Tinea Corporis (Ringworm)"
  },
  {
    "id": "3e4fd005-d4aa-11e7-a562-0242ac120003",
    "icd": "N48.89",
    "name": "Sebaceous Glands (Fordyce Spots)"
  },
  {
    "id": "3e500346-d4aa-11e7-a562-0242ac120003",
    "icd": "L70.0",
    "name": "Acne Comedonica (Comedonal Acne)"
  },
  {
    "id": "3e4fca47-d4aa-11e7-a562-0242ac120003",
    "icd": "A51.0",
    "name": "Syphilis"
  },
  {
    "id": "3e4f96d1-d4aa-11e7-a562-0242ac120003",
    "icd": "L70.0",
    "name": "Acne Vulgaris"
  },
  {
    "id": "3e4f9bc2-d4aa-11e7-a562-0242ac120003",
    "icd": "A69.2",
    "name": "Borrelia"
  },
  {
    "id": "3e505aed-d4aa-11e7-a562-0242ac120003",
    "icd": "B00.9",
    "name": "Herpes Simplex"
  },
  {
    "id": "3e506d38-d4aa-11e7-a562-0242ac120003",
    "icd": "L43.0",
    "name": "Lichen Ruber Planus"
  },
  {
    "id": "3e4fc125-d4aa-11e7-a562-0242ac120003",
    "icd": "L71.0",
    "name": "Perioral Dermatitis"
  },
  {
    "id": "3e5075a7-d4aa-11e7-a562-0242ac120003",
    "icd": "B86.0",
    "name": "Scabies"
  },
  {
    "id": "3e51595d-d4aa-11e7-a562-0242ac120003",
    "icd": "L81.0",
    "name": "Postinflammatory Hyperpigmentation"
  },
  {
    "id": "3e510e8f-d4aa-11e7-a562-0242ac120003",
    "icd": "A63.0",
    "name": "Genital Warts (Condyloma)"
  },
  {
    "id": "3e4f9992-d4aa-11e7-a562-0242ac120003",
    "icd": "L20.9",
    "name": "Atopic Dermatitis"
  },
  {
    "id": "3e517bcf-d4aa-11e7-a562-0242ac120003",
    "icd": "L02.92",
    "name": "Furuncle (Deep Folliculitis)"
  },
  {
    "id": "3e4fa213-d4aa-11e7-a562-0242ac120003",
    "icd": "L30.0",
    "name": "Eczema Nummularis"
  },
  {
    "id": "3e518ec2-d4aa-11e7-a562-0242ac120003",
    "icd": "C44.92",
    "name": "Squamous Cell Carcinoma"
  },
  {
    "id": "3e51ccc4-d4aa-11e7-a562-0242ac120003",
    "icd": "D18.01",
    "name": "Angioma (Cherry angioma)"
  },
  {
    "id": "3e519623-d4aa-11e7-a562-0242ac120003",
    "icd": "T14.0",
    "name": "Hematoma/Ecchymosis (Bruise)"
  },
  {
    "id": "3e4fdae1-d4aa-11e7-a562-0242ac120003",
    "icd": "D23.9",
    "name": "Dermatofibroma"
  },
  {
    "id": "3e4f9eae-d4aa-11e7-a562-0242ac120003",
    "icd": "L72.0",
    "name": "Epidermal / Benign Cyst"
  },
  {
    "id": "3e51108f-d4aa-11e7-a562-0242ac120003",
    "icd": "L73.9",
    "name": "Folliculitis"
  },
  {
    "id": "b698ea32-8f39-4bb5-82e5-082a58f013d7",
    "icd": "D22.9",
    "name": "Atypical Melanocytic Lesion (Atypical Nevus)"
  },
  {
    "id": "3e50698e-d4aa-11e7-a562-0242ac120003",
    "icd": "L81.4",
    "name": "Lentigo"
  },
  {
    "id": "3e51b0b9-d4aa-11e7-a562-0242ac120003",
    "icd": "L28.1",
    "name": "Prurigo Nodularis"
  },
  {
    "id": "639ce6b0-e32d-11e7-bf93-29090e80497e",
    "icd": "B35.9",
    "name": "Tinea UNS"
  },
  {
    "id": "3e506501-d4aa-11e7-a562-0242ac120003",
    "icd": "L85.8",
    "name": "Keratosis Pilaris"
  },
  {
    "id": "3e50f536-d4aa-11e7-a562-0242ac120003",
    "icd": "L30.5",
    "name": "Pityriasis Alba"
  },
  {
    "id": "3e4fc43c-d4aa-11e7-a562-0242ac120003",
    "icd": "L42.9",
    "name": "Pityriasis Rosea"
  },
  {
    "id": "3e51d2b2-d4aa-11e7-a562-0242ac120003",
    "icd": "L70.8",
    "name": "Acne Scar(s)"
  },
  {
    "id": "3e500636-d4aa-11e7-a562-0242ac120003",
    "icd": "D21.9",
    "name": "Skin Tag"
  },
  {
    "id": "3e4f94a3-d4aa-11e7-a562-0242ac120003",
    "icd": "L70.0",
    "name": "Acne Nodulocystica"
  },
  {
    "id": "3e501f49-d4aa-11e7-a562-0242ac120003",
    "icd": "A46",
    "name": "Erysipelas"
  },
  {
    "id": "3e4fac86-d4aa-11e7-a562-0242ac120003",
    "icd": "T63.4",
    "name": "Insect Bite"
  },
  {
    "id": "3e510891-d4aa-11e7-a562-0242ac120003",
    "icd": "B08.1",
    "name": "Molluscum Contagiosum"
  },
  {
    "id": "3e4fd7d8-d4aa-11e7-a562-0242ac120003",
    "icd": "C44.91",
    "name": "Basal Cell Carcinoma"
  },
  {
    "id": "3e51ce16-d4aa-11e7-a562-0242ac120003",
    "icd": "L90.5",
    "name": "Scar"
  },
  {
    "id": "3e50a9c5-d4aa-11e7-a562-0242ac120003",
    "icd": "L80.9",
    "name": "Vitiligo"
  }
]') AS classification;
```
