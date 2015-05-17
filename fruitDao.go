package main

import (
	"github.com/jinzhu/gorm"
)

// FruitDao is a repository to store Fruits
type FruitDao struct {
	DB *gorm.DB
}

//instance methods
func (t FruitDao) add(fruit *Fruit) *Fruit {
	t.DB.Create(fruit)
	return fruit
}

// AddAll adds all given fruits to store
func (t FruitDao) AddAll(fruits []*Fruit) []*Fruit {
	for _, f := range fruits {
		t.add(f)
	}
	return fruits
}

// FindAll finds fruits in store
func (t FruitDao) FindAll() *[]Fruit {
	var fruits []Fruit
	t.DB.Find(&fruits)
	return &fruits
}

// FindByName finds matching fruits by name
func (t FruitDao) FindByName(name string) *[]Fruit {
	var fruits []Fruit
	t.DB.Where("name like ?", "%"+name+"%").Find(&fruits)
	return &fruits
}

// FindByID finds fruit by its id
func (t FruitDao) FindByID(id *int) *Fruit {
	fruit := Fruit{}
	t.DB.Find(&fruit, *id)
	return &fruit
}

// Update updates fruit
func (t FruitDao) Update(fruit *Fruit) *Fruit {
	t.DB.Save(fruit)
	return fruit
}

// RemoveAll removes all fruits in store
func (t FruitDao) RemoveAll() {
	t.DB.Delete(Fruit{})
}
