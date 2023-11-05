package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/awa-hc/backend/initializers/database"
	"github.com/awa-hc/backend/initializers/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Fullname        string `json:"fullname"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmpassword"`
		ImgProfile      string `json:"imgprofile"`
		Username        string `json:"username"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json invalid"})
	}
	if len(body.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 8 characters"})
		return
	}

	if body.ConfirmPassword != body.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password do not match"})
		return
	}
	if body.Fullname == "" || body.Email == "" || body.Password == "" || body.ConfirmPassword == "" || body.ImgProfile == "" || body.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty fields"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to hash the password"})
		return
	}

	user := models.Users{Fullname: body.Fullname, Email: body.Email, Password: string(hash), ImgProfile: body.ImgProfile, Username: body.Username}

	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create the user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully", "user": user})

}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseOK struct {
	Token string `json:"token"`
}

type LoginResponseFalse struct {
	Error string `json:"error"`
}

// @Summary Login Function
// @Description Verifica datos en la base de datos y devuelve un JWT en caso de éxito
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} LoginResponseOK
// @Success 400 {object} LoginResponseFalse
// @Router /auth/login [post]
// @Param data body LoginData true "datos del inicio de sesion"
func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json invalid"})
		return
	}

	var user models.Users
	database.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create token"})
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	id := user.(models.Users).ID

	c.JSON(http.StatusOK, gin.H{"user": user, "id": id})
}
