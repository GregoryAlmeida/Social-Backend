package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	type User struct {
		Id       string
		Name     string
		Email    string
		Password string
		Profile  string
		About    string
	}

	type Posts struct {
		Profile string
		Name    string
		Message string
	}

	user := []User{
		{
			Id:       "1",
			Name:     "Gregory",
			Email:    "gregorydealmeida901@gmail.com",
			Password: "123",
		},
		{
			Id:       "2",
			Name:     "Carlos Lucas",
			Email:    "carloslucas@gmail.com",
			Password: "1234",
		},
	}

	RUser := router.Group("/user")
	{
		//Logar em conta
		RUser.GET("/", func(ctx *gin.Context) {
			email := ctx.Query("email")
			password := ctx.Query("password")

			for _, v := range user {
				if v.Email == email && v.Password == password {
					ctx.JSON(http.StatusOK, gin.H{
						"Message": "OK",
						"Id":      v.Id,
					})
					return
				}
			}

			ctx.JSON(http.StatusOK, gin.H{
				"Message": "Not Found",
			})
		})

		//Cadastrar Usuario Novo
		RUser.POST("/cadastrar", func(ctx *gin.Context) {
			var newUser User
			ctx.ShouldBindJSON(&newUser)
			user = append(user, newUser)
			ctx.JSON(http.StatusOK, gin.H{
				"Message": "OK",
			})

		})

		//Verificar se possui usuario atraves do email
		RUser.GET("/:email", func(ctx *gin.Context) {
			email := ctx.Param("email")
			emailVerify := false

			for _, v := range user {
				if v.Email == email {
					emailVerify = true
				}
			}

			if !emailVerify {
				ctx.JSON(http.StatusOK, gin.H{"Message": "OK"})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"Message": "ERRO"})
			}
		})

		//GET User Completo
		RUser.GET("/usuario/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")

			for i, v := range user {
				if v.Id == id {
					ctx.JSON(http.StatusOK, user[i])
				}
			}
		})

		//GET todos usuarios
		RUser.GET("/allusers", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, user)
		})

		RUser.POST("/atualizarDados/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			var newData User

			ctx.ShouldBindJSON(&newData)

			for i, v := range user {
				if v.Id == id {
					user[i] = newData
				}
			}

			ctx.JSON(http.StatusOK, gin.H{
				"Message": "OK",
			})
		})
	}

	posts := []Posts{
		{
			Profile: "",
			Name:    "Admin Gregs",
			Message: "Bem vindo!",
		},
		{
			Profile: "",
			Name:    "Admin Gregs",
			Message: "Você pode iniciar suas postagens aqui...",
		},
	}
	//Rotas de postagens
	RUser = router.Group("/postagens")
	{
		//Get para pegar todas as postagens
		RUser.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, posts)
		})

		//Rota post de Postagens
		RUser.POST("/postar", func(ctx *gin.Context) {
			var newPost Posts
			ctx.ShouldBindJSON(&newPost)

			posts = append(posts, newPost)
		})
	}

}
