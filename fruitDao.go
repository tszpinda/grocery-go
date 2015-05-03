package main

import (
	"github.com/jinzhu/gorm"
)

type FruitDao struct {
	DB *gorm.DB
}

//instance methods
func (t FruitDao) add(fruit *Fruit) *Fruit {
	t.DB.Create(fruit)
	return fruit
}

func (t FruitDao) AddAll(fruits []*Fruit) []*Fruit {
	for _, f := range fruits {
		t.add(f)
	}
	return fruits
}

func (t FruitDao) FindAll() *[]Fruit {
	fruits := make([]Fruit, 0)
	t.DB.Find(&fruits)
	return &fruits
}

func (t FruitDao) FindByName(name string) *[]Fruit {
	fruits := make([]Fruit, 0)
	t.DB.Where("name like ?", "%"+name+"%").Find(&fruits)
	return &fruits
}

func (t FruitDao) FindById(id *int) *Fruit {
	fruit := Fruit{}
	t.DB.Find(&fruit, *id)
	return &fruit
}

func (t FruitDao) Update(fruit *Fruit) *Fruit {
	t.DB.Save(fruit)
	return fruit
}

func (t FruitDao) RemoveAll() {
	t.DB.Delete(Fruit{})
}
