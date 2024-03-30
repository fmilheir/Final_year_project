package controller

import (
	"log"
	"os"
	"strconv"
	"time"
	"net/http"
	"bytes"
	"encoding/json"


	"github.com/dgrijalva/jwt-go"
	"github.com/fmilheir/final_year_project/backend/database"
	"github.com/fmilheir/final_year_project/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	
)

func init() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Register_admin(c *fiber.Ctx) error {
	println("Register_admin!!!!!!!!!")
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Create the company
	company := models.Company{
		Name: data["companyName"],
	}
	database.DB.Db.Create(&company)
	// Hash the password
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.Admin{
		Username: data["username"],
		Email:    data["email"],
		Password: password,
		CompanyID: company.ID,
	}
	// Create a new record
	database.DB.Db.Create(&user)
	return c.JSON(data)
}

func Register_user(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Hash the password
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Username: data["username"],
		Email:    data["email"],
		Password: password,
		CompanyID: 1,
	}
	// Create a new record
	database.DB.Db.Create(&user)
	return c.JSON(data)
}


func Login_admin(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user models.Admin
	database.DB.Db.Where("email = ?", data["email"]).First(&user)
	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}
	// Compare the password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}
	
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{
		"message": "Success",
	})
}


func Login_user(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Db.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}
	// Compare the password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User
	database.DB.Db.Where("id = ?", claims.Issuer).First(&user)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
func synchronousChatbot(c *fiber.Ctx) error {
    type RequestBody struct {
        Message string `json:"message"`
    }

    var body RequestBody
    if err := c.BodyParser(&body); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    flaskRequestBody, err := json.Marshal(body)
    if err != nil {
        log.Printf("Error encoding request body: %v\n", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error encoding request body"})
    }

    req, err := http.NewRequest("POST", "http://chatbot:5000/webhook", bytes.NewBuffer(flaskRequestBody))
    if err != nil {
        log.Printf("Error creating request to chatbot service: %v\n", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating request to chatbot service"})
    }

    req.Header.Add("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Printf("Error contacting chatbot service: %v\n", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error contacting chatbot service"})
    }
    defer resp.Body.Close()

    var chatbotResp map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&chatbotResp); err != nil {
        log.Printf("Error decoding response from chatbot service: %v\n", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error decoding response from chatbot service"})
    }

    // Return the chatbot's response directly to the client
	println("chatbotResp, ", resp.Body)
    return c.Status(fiber.StatusOK).JSON(chatbotResp)
}

func MyHandler(c *fiber.Ctx) error {
    return synchronousChatbot(c)
}
