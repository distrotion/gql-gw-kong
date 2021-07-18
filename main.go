package main

import (
	"context"
	_ "encoding/json"
	"fmt"
	_ "fmt"
	"gogatewaydemo/jwt"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
)

type Tokenschema struct {
	Token string `json:"token"`
}

type Gqlquery struct {
	Query string `json:"query"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		//-----------------------------------------

		token := "Server is ready"

		//=========================================

		c.JSON(200, token)
	})

	r.POST("/encrypt", func(c *gin.Context) {

		var input Gqlquery
		c.ShouldBind(&input)
		//-----------------------------------------

		str := fmt.Sprintf("%v", input.Query)
		token, err := jwt.GenerateToken(str)
		if err != nil {
			return
		}

		//=========================================

		c.JSON(200, token)
	})

	r.POST("/decrypt", func(c *gin.Context) {

		var input Tokenschema
		c.ShouldBind(&input)

		//-----------------------------------------
		fmt.Println(input.Token)
		// str := fmt.Sprintf("%v", input)

		tokenStr := input.Token
		output, err := jwt.ParseToken(tokenStr)
		if err != nil {
			return
		}

		//=========================================

		c.JSON(200, output)
	})

	r.POST("/auth", func(c *gin.Context) {

		var input Gqlquery
		//c.ShouldBind(&input) BindJSON
		c.BindJSON(&input)

		//-----------------------------------------
		//fmt.Println(reflect.TypeOf(input.Query))
		fmt.Println(input.Query)
		// str := fmt.Sprintf("%v", input)

		graphqlClient := graphql.NewClient("http://34.101.204.44:9110/query")
		graphqlRequest := graphql.NewRequest(input.Query)
		var graphqlResponse interface{}
		if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
			panic(err)
		}
		fmt.Println(graphqlResponse)

		//=========================================

		c.JSON(200, graphqlResponse)
	})

	r.POST("/noti", func(c *gin.Context) {

		var input Gqlquery
		//c.ShouldBind(&input) BindJSON
		c.BindJSON(&input)

		//-----------------------------------------
		//fmt.Println(reflect.TypeOf(input.Query))
		fmt.Println(input.Query)
		// str := fmt.Sprintf("%v", input)

		graphqlClient := graphql.NewClient("http://34.101.190.173:9111/query")
		graphqlRequest := graphql.NewRequest(input.Query)
		var graphqlResponse interface{}
		if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
			panic(err)
		}
		fmt.Println(graphqlResponse)

		//=========================================

		c.JSON(200, graphqlResponse)
	})

	r.Run(":9210")
}
