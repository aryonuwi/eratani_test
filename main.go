package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aryonuwi/eratani_test/soaldua"
	"github.com/aryonuwi/eratani_test/soalempat"
	"github.com/aryonuwi/eratani_test/soalsatu"
	"github.com/aryonuwi/eratani_test/soaltiga/models"
	"github.com/aryonuwi/eratani_test/soaltiga/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Soal no 1
	intData := [11]int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1000}
	for _, val := range intData {
		result := soalsatu.Polycarp(val)
		fmt.Println(result)
	}

	fmt.Println("+============================================+")
	// Soal no 2
	testDataPalindrome := [3]interface{}{"kodokkodok", "jinggaaku", 123321}
	for _, val := range testDataPalindrome {
		palindrome := soaldua.Plindrome(val)
		fmt.Println(palindrome)
	}

	fmt.Println("+============================================+")
	// Soal no 4
	randomNumbers := []int{9, 2, 5, 1, 7, 3, 8, 4, 6}
	response := soalempat.Sorting(randomNumbers)
	fmt.Println(response)

	fmt.Println("+============================================+")
	// Soal no 3
	/**
	4 A
	"SELECT
		u.country,
		count(u.country) as count_country,
		sum(o.total_buy) as buying
	FROM
		users u
	INNER JOIN
		orders o
	ON
		o.id_user =u.id
	GROUP BY
		u.country
	ORDER BY
		buying DESC "

	4 B
	"SELECT
	u.credit_card_type,
		count(u.credit_card_type) as count_credit_card_type
	FROM
		users u
	group by
		u.credit_card_type
	order by
		count_credit_card_type DESC
	limit 1
	*/

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Load env is failed")
	}
	models.DatabaseConnection()
	app := fiber.New()
	routes.RoutesInit(app)

	app.Listen(":" + os.Getenv("ACCESS_PORT"))

}
