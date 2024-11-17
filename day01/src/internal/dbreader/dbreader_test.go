package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadDB(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		wantErr  bool
	}{
		{
			name:     "Read JSON file",
			filepath: "../../testdata/testrecipes.json",
			wantErr:  false,
		},
		{
			name:     "Read XML file",
			filepath: "../../testdata/testrecipes.xml",
			wantErr:  false,
		},
		{
			name:     "Unsupported file format",
			filepath: "../../testdata/recipes.txt",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := ReadDB(tt.filepath)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestJSONReader(t *testing.T) {
	file, err := os.Open("../../testdata/testrecipes.json")
	if err != nil {
		t.Fatalf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	reader := JSONReader{}
	recipes, err := reader.Read(file)
	assert.NoError(t, err)

	expected := &Recipes{
		Cakes: []Cake{
			{
				Name: "Chocolate Cake",
				Time: "1h",
				Ingredients: []Ingredient{
					{Name: "Flour", Count: "2", Unit: "cups"},
					{Name: "Sugar", Count: "1.5", Unit: "cups"},
				},
			},
			{
				Name: "Carrot Cake",
				Time: "1h 15m",
				Ingredients: []Ingredient{
					{Name: "Carrot", Count: "3", Unit: "pieces"},
					{Name: "Flour", Count: "2", Unit: "cups"},
				},
			},
		},
	}
	assert.Equal(t, expected, recipes)

	jsonData, err := reader.Convert(recipes)
	assert.NoError(t, err)

	var jsonRecipes Recipes
	err = xml.Unmarshal(jsonData, &jsonRecipes)
	assert.NoError(t, err)
	assert.Equal(t, expected, &jsonRecipes)
}

func TestXMLReader(t *testing.T) {
	file, err := os.Open("../../testdata/testrecipes.xml")
	if err != nil {
		t.Fatalf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	reader := XMLReader{}
	recipes, err := reader.Read(file)
	assert.NoError(t, err)

	expected := &Recipes{
		Cakes: []Cake{
			{
				Name: "Chocolate Cake",
				Time: "1h",
				Ingredients: []Ingredient{
					{Name: "Flour", Count: "2", Unit: "cups"},
					{Name: "Sugar", Count: "1.5", Unit: "cups"},
				},
			},
			{
				Name: "Carrot Cake",
				Time: "1h 15m",
				Ingredients: []Ingredient{
					{Name: "Carrot", Count: "3", Unit: "pieces"},
					{Name: "Flour", Count: "2", Unit: "cups"},
				},
			},
		},
	}
	assert.Equal(t, expected, recipes)

	xmlData, err := reader.Convert(recipes)
	assert.NoError(t, err)

	var xmlRecipes Recipes
	err = json.Unmarshal(xmlData, &xmlRecipes)
	assert.NoError(t, err)
	assert.Equal(t, expected, &xmlRecipes)
}
