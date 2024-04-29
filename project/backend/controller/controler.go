package controller

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"fmt"
	"strings"
	"io"
	"crypto/tls"	


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
	if data["companyName"] == "" || data["firstName"] == "" || data["email"] == "" || data["password"] == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Bad request please fill in all fields",
		})
	}

	// Create the company
	company := models.Company{
		Name: data["companyName"],
	}
	database.DB.Db.Create(&company)
	// Hash the password
	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.Admin{
		FirstName:  data["FirstName"],
		LastName:   data["LastName"],
		Email:     data["email"],
		Password:  password,
		CompanyID: company.ID,
	}
	// Create a new record
	database.DB.Db.Create(&user)
	return c.JSON(user)
}

func Register_user(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["firstName"] == "" || data["lastName"] == "" || data["email"] == "" || data["password"] == "" || data["role"] == "" || data["companyID"] == ""{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Bad request please fill in all fields",
		})
	}

	// Hash the password
	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	fmt.Println(data)
	companyIDStr := data["companyID"]
	companyID, err := strconv.Atoi(companyIDStr)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid company ID",
		})
	}

	user := models.User{
		FirstName:  data["firstName"],
		LastName:   data["lastName"],
		Email:     data["email"],
		Password:  password,
		Role: 	data["role"],
		CompanyID: uint(companyID),
	}
	// Create a new record
	database.DB.Db.Create(&user)
	return c.JSON(user)
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
		HTTPOnly: false,
		MaxAge:  86400,
	})
	return c.JSON(fiber.Map{
		"token": token,
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
	fmt.Print(user)
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
		"token": token,
		"message": "Success",
	})
}

func User(c *fiber.Ctx) error {
		// Get the Authorization header from the request
		authorizationHeader := c.Get("Authorization")
	
		// Check if Authorization header exists and has the expected format
		if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "Invalid Authorization header",
			})
		}
	
		// Extract the JWT token from the Authorization header
		token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	
		// Parse and validate the JWT token
		claims := &jwt.StandardClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"message": "Invalid token",
			})
		}
	
		// Extract user ID from token claims
		userID := claims.Issuer
		fmt.Println(userID)
	
		// Query the database to find the user by ID
		var user models.Admin
		if err := database.DB.Db.First(&user, userID).Error; err != nil {
			c.Status(fiber.StatusNotFound)
			return c.JSON(fiber.Map{
				"message": "User not found",
			})
		}
	
		// Return the user's first name along with other necessary information
		return c.JSON(fiber.Map{
			"ID": user.ID,
			"role": user.Role,
			"firstName": user.FirstName,
			"email": user.Email, // You can include other user information as needed
			// Add more fields as necessary
		})
}	


func Logout(c *fiber.Ctx) error {
		// Clear the JWT token cookie by setting an empty value and expiration in the past
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour), // Set expiration to past time
			HTTPOnly: true,
		})
	
		// Return success message
		return c.JSON(fiber.Map{
			"message": "Logout successful",
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
    return c.Status(fiber.StatusOK).JSON(chatbotResp)
}

func GetCompanyID(c *fiber.Ctx) error {

	// get user id from body
	var data map[string]string
	var admin models.Admin
	var user models.User


	if err := c.BodyParser(&data); err != nil {
		return err
	}
	userID := data["userID"]
	
	if err := database.DB.Db.First(&admin, userID).Error; err != nil {
		if err := database.DB.Db.First(&user, userID).Error; err != nil {
			c.Status(fiber.StatusNotFound)
			return c.JSON(fiber.Map{
				"message": "User not found",
			})
		}
	}

	// Return the user's company ID
	var companyID uint
	if admin.ID != 0 {
		companyID = admin.CompanyID
	} else {
		companyID = user.CompanyID
	}
	return c.JSON(fiber.Map{
		"companyID": companyID,
		"message":   "Company ID retrieved",
	})
}

func GetCompanyUsers(c *fiber.Ctx) error {
	// Get the company ID from the request body
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	companyIDStr := data["companyid"]
	companyID, err := strconv.Atoi(companyIDStr)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid company ID",
			"companyID": companyIDStr,
		})
	}

	// Query the database to find all users in the company
	var users []models.User
	if err := database.DB.Db.Where("company_id = ?", companyID).Find(&users).Error; err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "error retrieving users",
		})
	}
	if len(users) == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "No users found",
		})
	}

	// Return the list of users
	return c.JSON(users)
}

 func DeleteUser(c *fiber.Ctx) error {
	// Get the user ID from the request body
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	userIDStr := data["userID"]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid user ID",
		})
	}
	
	// Query the database to find the user by ID
	var user models.User
	if err := database.DB.Db.First(&user, userID).Error; err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Delete the user
	database.DB.Db.Delete(&user)
	return c.JSON(fiber.Map{
		"message": "User deleted",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	// Get the user ID from the request body
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	userIDStr := data["userID"]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid user ID",
		})
	}

	// Query the database to find the user by ID
	var user models.User
	if err := database.DB.Db.First(&user, userID).Error; err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Update the user's information
	if data["firstName"] != "" {
		user.FirstName = data["firstName"]
	}
	if data["lastName"] != "" {
		user.LastName = data["lastName"]
	}
	if data["email"] != "" {
		user.Email = data["email"]
	}
	if data["role"] != "" {
		user.Role = data["role"]
	}
	if data["password"] != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = password
	}

	// Save the updated user information
	database.DB.Db.Save(&user)
	return c.JSON(user)
}

func MyHandler(c *fiber.Ctx) error {
    return synchronousChatbot(c)
}

//ticekting system calling api in localhost 5000
func CreateTicket(c *fiber.Ctx) error {
  // Read the request body
  requestBody := c.Request().Body()

  // Pass incident information to the API in another container
  response, err := PassIncidentInfoToAPI(requestBody)
  if err != nil {
	  return c.Status(http.StatusInternalServerError).SendString(err.Error())
  }

  // Respond with the response from the API
  return c.SendString(response)
}

func PassIncidentInfoToAPI(data []byte) (string, error) {

	tlsConfig := &tls.Config{
        InsecureSkipVerify: true,
    }
    transport := &http.Transport{
        TLSClientConfig: tlsConfig,
    }
    client := &http.Client{
        Transport: transport,
    }

	reader := bytes.NewReader(data)
	fmt.Println(reader)
	// Make a POST request to the API in another container:
 	resp, err := client.Post("https://ticketing:3030/tmf-api/Incident/v4/incident", "application/json", reader)
    if err != nil {
		fmt.Println(err)
        return "", err
    }
    defer resp.Body.Close()
	fmt.Println("passed 1")

    // Read the response body
    responseBody := new(bytes.Buffer)
    _, err = io.Copy(responseBody, resp.Body)
    if err != nil {
		fmt.Println(err)
        return "", err
    }
	fmt.Println("passed", resp.StatusCode)

    // Check if the request was successful (HTTP status code 200)
    if resp.StatusCode != http.StatusOK {
        return "woaahhh", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    return responseBody.String(), nil
}

	
