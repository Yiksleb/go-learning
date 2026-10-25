package main

import "fmt"

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
	models := []Model3D{model, table, crate}

	foundModels, found := findModels(models, "Low_Poly_Table")
	if found {
		foundModels.PrintInfo()
	} else {
		fmt.Println("Model not found")
	}
	models = addModels(models, robot)
	for _, item := range models {
		item.PrintInfo()
	}
	fmt.Println(robot.IsLowPoly())
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
	fmt.Println("Format", m.Format)
	fmt.Println("Polygons:", m.Polygons)
}
func (m Model3D) IsLowPoly() bool {
	if m.Polygons < 1000 {
		return true
	} else {
		return false
	}
}
func addModels(models []Model3D, newModel Model3D) []Model3D {
	models = append(models, newModel)
	return models
}
