package main

type Animal interface {
	MakeSound() string
	GetHabitat() string
	GetInfo() string
	WeeklyCost() float64
	Health() string
	ID() string
	Species() string
}

type BaseAnimal struct {
	IDValue       string
	Name          string
	SpeciesValue  string
	Age           int
	HealthStatus  string
	DailyFoodCost float64
}

type Zookeeper struct {
	EmployeeId      string
	Name            string
	Specialization  string
	AssignedAnimals []Animal
}
type Zoo struct {
	Name       string
	Animals    []Animal
	Zookeepers []*Zookeeper
}

func (z Zookeeper) FeedAnimal(a Animal) {
	println(z.Name, "fed", a.GetInfo())
}

func (z *Zookeeper) AssignAnimal(a Animal) {
	z.AssignedAnimals = append(z.AssignedAnimals, a)
}
func (z Zookeeper) CheckHealth(a Animal) {
	println(
		z.Name,
		"checked health of",
		a.GetInfo(),
		"- Status:",
		a.Health(),
	)
}
func (z Zookeeper) GetWorkload() int {
	return len(z.AssignedAnimals)
}

type Lion struct {
	BaseAnimal
	ManeColor string
	PrideSize int
}
type Elephant struct {
	BaseAnimal
	TuskLength float64
	Weight     float64
}

func (z *Zoo) AddAnimal(a Animal) {
	z.Animals = append(z.Animals, a)
}
func (z *Zoo) RemoveAnimalByID(id string) {
	filtered := make([]Animal, 0)
	for _, a := range z.Animals {

		if a.ID() != id {
			filtered = append(filtered, a)
		}
	}

	z.Animals = filtered
}
func (z *Zoo) AssignAnimalToKeeper(a Animal, keeper *Zookeeper) {
	keeper.AssignAnimal(a)
}
func (z Zoo) GetAnimalsByHabitat(habitat string) []Animal {
	result := make([]Animal, 0)

	for _, a := range z.Animals {
		if a.GetHabitat() == habitat {
			result = append(result, a)
		}
	}

	return result
}

func (z Zoo) GetAnimalsBySpecies(species string) []Animal {
	result := make([]Animal, 0)

	for _, a := range z.Animals {
		if a.Species() == species {
			result = append(result, a)
		}
	}

	return result
}
func (z Zoo) CalculateTotalWeeklyCost() float64 {
	total := 0.0

	for _, a := range z.Animals {
		total += a.WeeklyCost()
	}

	return total
}
func (z Zoo) DisplayAllAnimals() {
	println("===", z.Name, "- All Animals ===")

	for _, a := range z.Animals {
		println(a.GetInfo(), "- Habitat:", a.GetHabitat())
	}
}
func (z Zoo) GetZooStatistics() {
	totalAnimals := len(z.Animals)
	totalKeepers := len(z.Zookeepers)
	totalCost := z.CalculateTotalWeeklyCost()

	println("=== Zoo Statistics ===")
	println("Total Animals:", totalAnimals)
	println("Total Zookeepers:", totalKeepers)
	println("Total Weekly Maintenance:", totalCost)
}

func (b BaseAnimal) WeeklyCost() float64 {
	return b.DailyFoodCost * 7
}

func (b BaseAnimal) GetInfo() string {
	return b.IDValue + " - " + b.Name + " (" + b.SpeciesValue + ")"
}
func (b BaseAnimal) Health() string {
	return b.HealthStatus
}
func (b BaseAnimal) ID() string {
	return b.IDValue
}
func (b BaseAnimal) Species() string {
	return b.SpeciesValue
}

func (l Lion) MakeSound() string {
	return "Roar!"
}
func (l Lion) GetHabitat() string {
	return "Savanna"
}
func (e Elephant) MakeSound() string {
	return "Trumpet"
}
func (e Elephant) GetHabitat() string {
	return "Grassland"
}

func main() {
	lion := Lion{
		BaseAnimal: BaseAnimal{
			IDValue:       "A001",
			Name:          "Simba",
			SpeciesValue:  "African Lion",
			Age:           5,
			HealthStatus:  "Healthy",
			DailyFoodCost: 50,
		},
		ManeColor: "Golden",
		PrideSize: 3,
	}

	elephant := Elephant{
		BaseAnimal: BaseAnimal{
			IDValue:       "A002",
			Name:          "Dumbo",
			SpeciesValue:  "African Elephant",
			Age:           15,
			HealthStatus:  "Healthy",
			DailyFoodCost: 80,
		},
		TuskLength: 2.5,
		Weight:     5000,
	}

	zoo := Zoo{Name: "Safari World"}
	zoo.AddAnimal(lion)
	zoo.AddAnimal(elephant)

	for _, a := range zoo.Animals {
		println(a.GetInfo())
		println("Sound:", a.MakeSound())
		println("Habitat:", a.GetHabitat())
		println("Weekly Cost:", a.WeeklyCost())
		println("----")
	}
}
