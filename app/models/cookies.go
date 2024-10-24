package models

type Cookie struct {
	Id   int
	Name string
	Cost string
	Img  string
}

var Cookies = []Cookie{
	{
		Id:   1,
		Name: "Печенье 1",
		Cost: "2500 ₽ за 1 кг",
		Img:  "/static/images/cookies/cookie1.jpg"},
	{Id: 2,
		Name: "Печенье 2",
		Cost: "2500 ₽ за 1 кг",
		Img:  "/static/images/cookies/cookie2.jpg"},
	{Id: 3,
		Name: "Печенье 3",
		Cost: "2500 ₽ за 1 кг",
		Img:  "/static/images/cookies/cookie3.jpg"},
	{Id: 4,
		Name: "Печенье 4",
		Cost: "2500 ₽ за 1 кг",
		Img:  "/static/images/cookies/cookie4.jpg"},
	{Id: 5,
		Name: "Печенье 5",
		Cost: "2500 ₽ за 1 кг",
		Img:  "/static/images/cookies/cookie5.jpg"},
	{Id: 6,
		Name: "Печенье 6",
		Cost: "2500 ₽ за 1 кг",
		Img:  "/static/images/cookies/cookie6.jpg"},
	{Id: 7,
		Name: "Печенье 7",
		Cost: "2500 ₽ за 1 кг",
		Img:  "/static/images/cookies/cookie7.jpg"},
	{Id: 8,
		Name: "Печенье 8",
		Cost: "2500 ₽ за 1 кг",
		Img:  "/static/images/cookies/cookie8.jpg"},
	{Id: 9,
		Name: "Печенье 9",
		Cost: "2500 ₽ за 1 кг",
		Img:  "/static/images/cookies/cookie9.jpg"},
	{Id: 10,
		Name: "Пироженное 1",
		Cost: "2500 ₽ за 1 кг",
		Img:  "/static/images/cookies/cookie10.jpg"},
}
