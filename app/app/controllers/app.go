package controllers

import (
	"app/app/models"

	"github.com/revel/revel"
)

type App struct {
	TransactionalController
}

func (c App) Index() revel.Result {
	text := "HELL"
	return c.Render(text)
}

func (c App) Hype() revel.Result {
	type Result struct {
		Message string `json:"message"`
		Number  int    `json:"number"`
	}

	result := Result{
		Message: "hogehoge",
		Number:  10,
	}

	return c.RenderJSON(result)
}

func (c App) GenerateUser() revel.Result {
	first := c.Params.Query.Get("first")
	last := c.Params.Query.Get("last")

	user := models.User{
		FirstName: first,
		LastName:  last,
		Password:  "aaaa",
	}

	models.DB.Create(&user)

	return c.RenderJSON(user)
}

func (c App) User() revel.Result {
	// for i := 0; i < 1; i++ {
	// 	DbMap.Insert(&models.User{
	// 		Id:        0,
	// 		FirstName: fmt.Sprintf("first%d", i),
	// 		LastName:  fmt.Sprintf("last%d", i),
	// 		Password:  "gogo",
	// 	})
	// }

	var users []models.User
	models.DB.Limit(5).Find(&users)

	// var user models.User
	// models.DB.First(&user)

	// model.DB.Select(&users, "select * from users")

	for _, r := range users {
		user := r
		c.Log.Info("%s(%s)", user.FirstName, user.ID)
	}

	// users := []models.User{}

	// users = append(users, models.User{
	// 	Id:        1,
	// 	FirstName: "NAME",
	// 	LastName:  "SECOND",
	// 	Password:  "pass",
	// })

	return c.RenderJSON(users)
}
