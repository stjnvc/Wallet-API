package handler

import (
	"bytes"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/stjnvc/wallet-api/internal/api/v1/dto"
	"github.com/stjnvc/wallet-api/internal/api/v1/model"
	"github.com/stjnvc/wallet-api/internal/api/v1/service"
	"golang.org/x/crypto/bcrypt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {

	logrus.Info("Login request:", c.Request.Method, c.Request.Body)

	var loginUserDTO dto.LoginUserDTO

	if err := c.ShouldBindJSON(&loginUserDTO); err != nil {
		logrus.Error("Login request error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.Login(loginUserDTO.Username, loginUserDTO.Password)
	if err != nil {
		logrus.Error("Failed to login user:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUserDTO.Password)); err != nil {
		logrus.Error("Failed to verify user password:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	tokenString, err := token.SignedString([]byte("jwt-secret"))
	if err != nil {
		logrus.Error("Failed to generate JWT token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (h *AuthHandler) Register(c *gin.Context) {
	// Read the request body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Error("Failed to read request body:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	// Log the request body
	logrus.Info("Request body:", string(body))

	// Reset the request body so it can be read again
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	// Bind JSON data to CreateUserDto
	var createUserDTO dto.CreateUserDto

	if err := c.ShouldBindJSON(&createUserDTO); err != nil {
		if err == io.EOF {
			logrus.Error("Empty request body:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Empty request body"})
			return
		}

		logrus.Error("Failed to bind JSON data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON data"})
		return
	}

	// Validate CreateUserDto
	validate := validator.New()
	if err := validate.Struct(createUserDTO); err != nil {
		logrus.Error("Validation failed:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := model.NewUser(createUserDTO.Username, createUserDTO.Password, createUserDTO.Email)

	if err := h.authService.Register(newUser); err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"message": "User registered successfully",
			"user":    newUser,
		})
	} else {
		logrus.Error("Failed to create user:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to register user",
		})
	}
}
