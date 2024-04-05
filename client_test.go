package autoderm_go_client

import (
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

const (
	fixture = `{"success":true,"message":"completed","id":"00000000-0000-0000-0000-000007b030ce","predictions":[{"confidence":0.20118004282871116,"icd":"D18.01","name":"Angioma (Cherry angioma)","classificationId":"3e51ccc4-d4aa-11e7-a562-0242ac120003","readMoreUrl":"https://www.firstderm.com/angioma/"},{"confidence":0.15057508063238972,"icd":"D22.9","name":"Atypical Melanocytic Nevus","classificationId":"b698ea32-8f39-4bb5-82e5-082a58f013d7","readMoreUrl":"https://www.firstderm.com/mole-congenital-nevus/"},{"confidence":0.07693463345465025,"icd":"D22.9","name":"Nevus (Benign Mole)","classificationId":"3e4fd5a6-d4aa-11e7-a562-0242ac120003","readMoreUrl":"https://www.firstderm.com/mole-congenital-nevus/"},{"confidence":0.05560834131992198,"icd":"D23.9","name":"Dermatofibroma","classificationId":"3e4fdae1-d4aa-11e7-a562-0242ac120003","readMoreUrl":"https://www.firstderm.com/dermatofibroma/"},{"confidence":0.05559970467420187,"icd":"L81.0","name":"Postinflammatory Hyperpigmentation","classificationId":"3e51595d-d4aa-11e7-a562-0242ac120003","readMoreUrl":"https://www.firstderm.com/hyperpigmentation/"}]}`
)

func TestClient(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodPost,
		"https://autoderm.ai/v1/query?model=autoderm_v2_2&language=en&save_image=false",
		httpmock.NewStringResponder(
			200,
			fixture,
		),
	)

	client := NewClient("secret_dev_team")

	file, err := os.Open("skin.jpeg")
	require.NoError(t, err)
	defer file.Close()

	response, err := client.Query("skin.jpeg", file, false)
	require.NoError(t, err)
	require.Equal(t, &QueryResponse{
		Success: true,
		Message: "completed",
		ID:      uuid.MustParse("00000000-0000-0000-0000-000007b030ce"),
		Predictions: []Prediction{
			{
				Confidence:       0.20118004282871116,
				Icd:              "D18.01",
				Name:             "Angioma (Cherry angioma)",
				ClassificationID: uuid.MustParse("3e51ccc4-d4aa-11e7-a562-0242ac120003"),
				ReadMoreURL:      "https://www.firstderm.com/angioma/",
			},
			{
				Confidence:       0.15057508063238972,
				Icd:              "D22.9",
				Name:             "Atypical Melanocytic Nevus",
				ClassificationID: uuid.MustParse("b698ea32-8f39-4bb5-82e5-082a58f013d7"),
				ReadMoreURL:      "https://www.firstderm.com/mole-congenital-nevus/",
			},
			{
				Confidence:       0.07693463345465025,
				Icd:              "D22.9",
				Name:             "Nevus (Benign Mole)",
				ClassificationID: uuid.MustParse("3e4fd5a6-d4aa-11e7-a562-0242ac120003"),
				ReadMoreURL:      "https://www.firstderm.com/mole-congenital-nevus/",
			},
			{
				Confidence:       0.05560834131992198,
				Icd:              "D23.9",
				Name:             "Dermatofibroma",
				ClassificationID: uuid.MustParse("3e4fdae1-d4aa-11e7-a562-0242ac120003"),
				ReadMoreURL:      "https://www.firstderm.com/dermatofibroma/",
			},
			{
				Confidence:       0.05559970467420187,
				Icd:              "L81.0",
				Name:             "Postinflammatory Hyperpigmentation",
				ClassificationID: uuid.MustParse("3e51595d-d4aa-11e7-a562-0242ac120003"),
				ReadMoreURL:      "https://www.firstderm.com/hyperpigmentation/",
			},
		},
	}, response)
}
