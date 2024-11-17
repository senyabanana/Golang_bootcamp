package dbcomparer

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"json-xml-basics/internal/dbreader"
)

// captureOutput захватывает вывод функции f
func captureOutput(f func()) string {
	var buf bytes.Buffer
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = stdout
	io.Copy(&buf, r)

	return buf.String()
}

func TestCompareDB(t *testing.T) {
	oldRecipes := &dbreader.Recipes{
		Cakes: []dbreader.Cake{
			{
				Name: "Chocolate Cake",
				Time: "1h",
				Ingredients: []dbreader.Ingredient{
					{Name: "Flour", Count: "2", Unit: "cups"},
					{Name: "Sugar", Count: "1.5", Unit: "cups"},
				},
			},
			{
				Name: "Carrot Cake",
				Time: "1h 15m",
				Ingredients: []dbreader.Ingredient{
					{Name: "Carrot", Count: "3", Unit: "pieces"},
					{Name: "Flour", Count: "2", Unit: "cups"},
				},
			},
		},
	}

	newRecipes := &dbreader.Recipes{
		Cakes: []dbreader.Cake{
			{
				Name: "Chocolate Cake",
				Time: "1h 10m",
				Ingredients: []dbreader.Ingredient{
					{Name: "Flour", Count: "2", Unit: "cups"},
					{Name: "Sugar", Count: "1.5", Unit: "cups"},
					{Name: "Cocoa", Count: "0.5", Unit: "cups"},
				},
			},
			{
				Name: "Banana Cake",
				Time: "2h",
				Ingredients: []dbreader.Ingredient{
					{Name: "Banana", Count: "4", Unit: "pieces"},
					{Name: "Flour", Count: "2", Unit: "cups"},
				},
			},
		},
	}

	expectedOutput := `ADDED cake "Banana Cake"
REMOVED cake "Carrot Cake"
CHANGED cooking time for cake "Chocolate Cake" - "1h 10m" instead of "1h"
ADDED ingredient "Cocoa" for cake "Chocolate Cake"
`

	output := captureOutput(func() {
		CompareDB(oldRecipes, newRecipes)
	})

	assert.Equal(t, expectedOutput, output)
}
