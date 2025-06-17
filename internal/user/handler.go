package user

import (
	"net/http"
	"time"
	"user/vault/internal/auth"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	Repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{Repo: repo}
}

func (handler Handler) SignUp(c *gin.Context) {
	// get user data from request body
	var req SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		er := http.StatusBadRequest
		http.Error(c.Writer, "Invalid request", er)
		return
	}

	// generate userID
	userID := uuid.New().String() // Generate unique ID
	// generate hashed password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		er := http.StatusInternalServerError
		http.Error(c.Writer, "Internal server error", er)
		return
	}

	// create new user data
	user := &UserData{
		PK:       "USER#" + userID,
		SK:       "METADATA",
		GSI1PK:   "EMAIL#" + req.Email,
		UserID:   userID,
		UserName: req.Name,
		UserType: req.Type,
		Email:    req.Email,
		Password: hashedPassword,
	}

	// add new user data to db
	if err := handler.Repo.AddUser(c, *user); err != nil {
		er := http.StatusInternalServerError
		http.Error(c.Writer, "Internal server error", er)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (handler Handler) Login(c *gin.Context) {
	// get email, password from request form
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	user, err := handler.Repo.GetUserbyEmail(c, email)
	if err != nil {
		er := http.StatusNoContent
		http.Error(c.Writer, "User cannot be found", er)
		return
	}

	// check password
	match := auth.CheckPasswordHash(password, user.Password)
	if !match {
		er := http.StatusBadRequest
		http.Error(c.Writer, "Wrong Password", er)
		return
	}

	// create session token + csrf token for user
	sessionToken := auth.CreateToken(32)
	csrfToken := auth.CreateToken(32)
	// set session + csrf cookie
	c.SetCookie("session_token", sessionToken, int(time.Hour), "/", "localhost:8080", false, true)
	c.SetCookie("csrf_token", csrfToken, int(time.Hour), "/", "localhost:8080", false, false)
	session := SessionData{
		SessionToken: sessionToken,
		CSRFToken:    csrfToken,
	}
	// store session + csrf token in db
	// TODO: create seperate repo func to deal with session data?
	new_user := &UserData{
		PK:       "USER#" + user.UserID,
		SK:       "METADATA",
		GSI1PK:   "EMAIL#" + user.Email,
		UserID:   user.UserID,
		UserName: user.UserName,
		UserType: user.UserType,
		Email:    user.Email,
		Password: user.Password,
		Session:  session,
	}
	if err := handler.Repo.AddUser(c, *new_user); err != nil {
		er := http.StatusInternalServerError
		http.Error(c.Writer, "Internal server error", er)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login Success! Welcome user " + user.UserName})
}

func (handler Handler) Logout(c *gin.Context) {
	if err := auth.Authorize(c); err != nil {
		er := http.StatusUnauthorized
		http.Error(c.Writer, err.Error(), er)
		return
	}

	// clear cookie
	c.SetCookie("session_token", "", -1, "/", "localhost:8080", false, true)
	c.SetCookie("csrf_token", "", -1, "/", "localhost:8080", false, false)

	// clear token from db
	userID := c.Param("id")
	user, _ := handler.Repo.GetUserByID(c, userID)
	new_user := &UserData{
		PK:       "USER#" + user.UserID,
		SK:       "METADATA",
		GSI1PK:   "EMAIL#" + user.Email,
		UserID:   user.UserID,
		UserName: user.UserName,
		UserType: user.UserType,
		Email:    user.Email,
		Password: user.Password,
		Session: SessionData{
			SessionToken: "",
			CSRFToken:    "",
		},
	}
	if err := handler.Repo.AddUser(c, *new_user); err != nil {
		er := http.StatusInternalServerError
		http.Error(c.Writer, "Internal server error", er)
		return
	}

	// return 200
	c.IndentedJSON(http.StatusOK, gin.H{"message": ""})
}

func (handler Handler) GetProfile(c *gin.Context) {
	// retrieve user data
	userID := c.Param("id")
	user, err := handler.Repo.GetUserByID(c, userID)
	if err != nil {
		er := http.StatusNoContent
		http.Error(c.Writer, "User cannot be found", er)
	}

	c.JSON(http.StatusOK, user)
}

func (handler Handler) UpdateProfile(c *gin.Context) {
	userID := c.Param("id")
	// get user data from request body
	var req UserDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		er := http.StatusBadRequest
		http.Error(c.Writer, "Invalid request", er)
		return
	}

	// check user data in db
	_, err := handler.Repo.GetUserByID(c, userID)
	if err != nil {
		er := http.StatusNoContent
		http.Error(c.Writer, "User cannot be found", er)
	}

	// generate hashed password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		er := http.StatusInternalServerError
		http.Error(c.Writer, "Internal server error", er)
		return
	}
	// create new user data
	user := &UserData{
		PK:       "USER#" + userID,
		SK:       "METADATA",
		GSI1PK:   "EMAIL#" + req.Email,
		UserID:   userID,
		UserName: req.Name,
		UserType: req.Type,
		Email:    req.Email,
		Password: hashedPassword,
	}

	new_profile, err := handler.Repo.UpdateUser(c, *user)
	if err != nil {
		er := http.StatusInternalServerError
		http.Error(c.Writer, "User profile cannot be updated", er)
		return
	}

	c.JSON(http.StatusOK, new_profile)
}

func (handler Handler) GetGroups(c *gin.Context) {
	// retrieve group data from user data
	userID := c.Param("id")
	user, err := handler.Repo.GetUserByID(c, userID)
	if err != nil {
		er := http.StatusNoContent
		http.Error(c.Writer, "User cannot be found", er)
		return
	}

	c.JSON(http.StatusOK, user.Groups)
}

// TODO: create AddToGroup func
