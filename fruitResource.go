package main

import (
	"net/http"
	"net/url"
	"strconv"
)

type FruitResource struct {
	FruitDao FruitDao
}

func (t FruitResource) findFruits(u *url.URL, h http.Header, _ interface{}) (int, http.Header, *[]Fruit, error) {
	fruits := t.FruitDao.FindAll()
	return http.StatusOK, nil, fruits, nil
}

func (t FruitResource) searchFruits(u *url.URL, h http.Header, _ interface{}) (int, http.Header, *[]Fruit, error) {
	name := u.Query().Get("name")
	var fruits *[]Fruit
	if len(name) == 0 {
		fruits = t.FruitDao.FindAll()
	} else {
		fruits = t.FruitDao.FindByName(name)
	}
	return http.StatusOK, nil, fruits, nil
}

func (t FruitResource) updateFruitPrice(u *url.URL, h http.Header, price *float64) (int, http.Header, *Fruit, error) {
	idStr := u.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	fruit := t.FruitDao.FindById(&id)
	if fruit == nil {
		return http.StatusNotFound, nil, nil, nil
	}
	fruit.Price = *price
	t.FruitDao.Update(fruit)
	return http.StatusOK, nil, fruit, nil
}

func (t FruitResource) addOrOverride(u *url.URL, h http.Header, fruits []*Fruit) (int, http.Header, []*Fruit, error) {

	t.FruitDao.RemoveAll()
	t.FruitDao.AddAll(fruits)

	return http.StatusOK, nil, fruits, nil
}
