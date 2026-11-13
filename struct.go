package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Model3D struct {
	Name     string
	Format   string
	Polygons int
	Category string
}

func main() {

	model := Model3D{
		Name:     "Low_Poly_Chair",
		Format:   "GLB",
		Polygons: 850,
		Category: "Furnitur",
	}
	table := Model3D{
		Name:     "Low_Poly_Table",
		Format:   "FBX",
		Polygons: 650,
		Category: "Furnitur",
	}
	crate := Model3D{
		Name:     "SiFi_crate",
		Format:   "GLB",
		Polygons: 750,
		Category: "Dekor",
	}
	robot := Model3D{
		Name:     "SiFI_Robot",
		Format:   "FBX",
		Polygons: 4500,
		Category: "Character",
	}
	testAsset := Model3D{
		Name:     "Test Model_03",
		Format:   "FBX",
		Polygons: 684,
		Category: "Test",
	}
	testAsset2 := Model3D{
		Name:     "Test_Model_04",
		Format:   "GLB",
		Polygons: 5785,
		Category: "Test",
	}
	models := []Model3D{model, table, crate}

	//d := deleteModel(models, "Unknown_Model")
	//for _, item := range d {
	//	fmt.Println("Loaded from deleteModels:")
	//	item.PrintInfo()
	//}
	loadedModels, err := loadModels()
	if err != nil {
		fmt.Println("Load Error:", err)
		return
	}
	models = loadedModels

	models, added := addModels(models, robot)
	fmt.Println("Was model added:", added)

	models, added = addModels(models, testAsset)
	fmt.Println("Was model added:", added)
	models, added = addModels(models, testAsset2)
	fmt.Println("Was model added:", added)

	err = saveModels(models)
	if err != nil {
		fmt.Println("Save Error:", err)
		return
	}

	result_Poly := findModelsByPoligon(models, 500, 1000)
	for _, item := range result_Poly {
		item.PrintInfo()
	}
	result_Format := findModelsByFormat(models, "FBX")
	for _, item := range result_Format {
		fmt.Println("______________")
		item.PrintInfo()
		fmt.Println("______________")
	}
	countFurnitur, foundFurnitur := countModelsByCategory(models, "Character")
	fmt.Println("Furnitur count", countFurnitur)
	fmt.Println("Found", foundFurnitur)
	average, found := getAveragePolygons(models)
	fmt.Println("*********************")
	fmt.Println("Average:", average)
	fmt.Println("Found:", found)
	fmt.Println("*********************")
	result := getModelsByCategoryaAndPolygons(models, "Furnitur", 1000)
	for _, item := range result {
		item.PrintInfo()
	}
	fmt.Println("////////////////////////")
	printModelStatistics(models)
	fmt.Println("////////////////////////")
	result = filterModels(models, func(item Model3D) bool {
		return item.Polygons < 1000 || item.Format == "FBX"
	})
	for _, item := range result {
		item.PrintInfo()
	}
}
func loadModels() ([]Model3D, error) {
	fileData, err := os.ReadFile("models.json")
	if err != nil {
		fmt.Println("Read error:", err)
		return nil, err
	}
	var loadedModels []Model3D
	err = json.Unmarshal(fileData, &loadedModels)
	if err != nil {
		fmt.Println("Error load:", err)
		return nil, err
	}
	return loadedModels, nil
}

func saveModels(models []Model3D) error {
	jsonData, err := json.MarshalIndent(models, "", " ")
	if err != nil {
		fmt.Println("Error Func Save :", err)
		return err
	}
	erre := os.WriteFile("models.json", jsonData, 0644)
	if erre != nil {
		fmt.Println("Os.WeiteFile error:", erre)
		return erre
	}
	return nil
}

func findModels(models []Model3D, name string) (Model3D, bool) {
	for _, item := range models {
		if item.Name == name {
			return item, true
		}
	}
	return Model3D{}, false
}
func (m Model3D) PrintInfo() {
	fmt.Println("Name:", m.Name)
	fmt.Println("Format:", m.Format)
	fmt.Println("Polygons:", m.Polygons)
	fmt.Println("Category:", m.Category)
}
func (m Model3D) IsLowPoly() bool {
	if m.Polygons < 1000 {
		return true
	} else {
		return false
	}
}
func addModels(models []Model3D, newModel Model3D) ([]Model3D, bool) {
	_, found := findModels(models, newModel.Name)
	if found {
		fmt.Println("Model Found")
		return models, false
	} else {
		models = append(models, newModel)
		return models, true
	}
}
func findModelsIsByCategory(models []Model3D, category string) []Model3D {
	result := []Model3D{}
	for _, item := range models {
		if strings.ToLower(item.Category) == strings.ToLower(category) {
			result = append(result, item)
		}

	}
	if len(result) == 0 {
		fmt.Println("Category is not found")
	}

	return result

}
func deleteModel(models []Model3D, name string) []Model3D {
	for i, item := range models {
		if item.Name == name {
			models = append(models[:i], models[i+1:]...)
			return models
		}
	}
	return models
}
func findModelsByPoligon(models []Model3D, minPolygons int, maxPolygons int) []Model3D {
	result := []Model3D{}
	for _, item := range models {
		if item.Polygons >= minPolygons && item.Polygons <= maxPolygons {
			result = append(result, item)
		}
	}
	return result
}
func findModelsByFormat(models []Model3D, format string) []Model3D {
	result := []Model3D{}
	for _, item := range models {
		if strings.ToLower(item.Format) == strings.ToLower(format) {
			result = append(result, item)
		}
	}
	return result
}
func countModelsByCategory(models []Model3D, category string) (int, bool) {
	result := 0
	for _, item := range models {
		if strings.ToLower(item.Category) == strings.ToLower(category) {
			result++
		}
	}
	if result != 0 {
		return result, true
	}
	return 0, false

}
func getLargestModel(models []Model3D) (Model3D, bool) {
	if len(models) == 0 {
		return Model3D{}, false
	}
	largestModel := models[0]
	for _, item := range models {
		if largestModel.Polygons < item.Polygons {
			largestModel = item
		}
	}

	return largestModel, true

}
func getAveragePolygons(models []Model3D) (float64, bool) {
	if len(models) == 0 {
		return 0, false
	}
	totalPolygons := 0
	for _, item := range models {
		totalPolygons += item.Polygons
	}

	average := float64(totalPolygons) / float64(len(models))
	return float64(average), true
}
func getTotalPolygons(models []Model3D) int {
	if len(models) == 0 {
		fmt.Println("Slice is empty")
		return 0
	}
	totalPolygons := 0
	for _, item := range models {
		totalPolygons += item.Polygons
	}
	return totalPolygons
}
func gatSmallestModel(models []Model3D) (Model3D, bool) {
	if len(models) == 0 {
		return Model3D{}, false
	}
	smallestModel := models[0]
	for _, item := range models {
		if smallestModel.Polygons > item.Polygons {
			smallestModel = item
		}
	}
	return smallestModel, true
}
func getModelsByCategoryaAndPolygons(models []Model3D, category string, maxPolygons int) []Model3D {
	result := []Model3D{}
	for _, item := range models {
		if item.Category == category && item.Polygons <= maxPolygons {
			result = append(result, item)
		}
	}
	return result
}
func printModelStatistics(models []Model3D) {
	fmt.Println("Total Models:", len(models))
	total := getTotalPolygons(models)
	fmt.Println("Total Polygons:", total)
	average, found := getAveragePolygons(models)
	fmt.Println("Average:", average, "Found:", found)
	largModel, found := getLargestModel(models)
	if found {
		fmt.Println("Model is found:", found)
		fmt.Println("Largest model is:", largModel.Name, largModel.Polygons)
	}
	smallestModel, found := gatSmallestModel(models)
	if found {
		fmt.Println("Model is found:", found)
		fmt.Println("Smallest model is:", smallestModel.Name, smallestModel.Polygons)
	}

}
func filterModels(models []Model3D, check func(Model3D) bool) []Model3D {
	result := []Model3D{}
	for _, item := range models {
		if check(item) {
			result = append(result, item)
		}
	}
	return result
}
