package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Truy cập /players để xem danh sách cầu thủ")
	})

	r.GET("/players", func(c *gin.Context) {
		type Player struct {
			Name     string   `json:"name"`
			Age      int      `json:"age"`
			Position string   `json:"position"`
		}

		var player1 = Player{
			Name: "Lionel Messi",
			Age: 30,
			Position: "Striker",
		}

		var player2 = Player{
			Name: "Cristiano Ronaldo",
			Age: 30,
			Position: "Striker",
		}

		var player3 = Player{
			Name: "Luka Modric",
			Age: 30,
			Position: "Midfielder",
		}

		var player4 = Player{
			Name: "David de Gea",
			Age: 30,
			Position: "Goalkeeper",
		}

		var playerList = []Player{player1, player2, player3, player4}
		c.JSON(http.StatusOK, playerList)
	})

	r.Run(":8080")
}