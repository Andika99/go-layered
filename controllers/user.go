package controllers

import (
	"strconv"
	"net/http"
	"go-layered/models"
	"go-layered/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Interface
type UserController interface {
	GetUsers(c *gin.Context)
	GetUserById(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUserById(c *gin.Context)
}

// Struct
type userController struct {
	repo repository.UserRepository
}

// Constructor
func NewUserController(repo repository.UserRepository) UserController {
	return &userController{
		repo: repo,
	}
}

// Methods
func (ctrl *userController) GetUsers(c *gin.Context) {
	users, err := ctrl.repo.FindAll()
	
	if (errorHandler(err, c)) { return }
	
	c.JSON(http.StatusOK, users)
}

func (ctrl *userController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	
	user, err := ctrl.repo.FindById(id)
	
	if (errorHandler(err, c)) { return }
	
	c.JSON(http.StatusOK, user)
}

func (ctrl *userController) CreateUser(c *gin.Context) {
	var input models.User
	
	err := c.ShouldBindJSON(&input)
	if (err != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	
	str_id := strconv.FormatUint(uint64(input.ID), 10)
	
	_, err = ctrl.repo.FindById(str_id)
	if (err == nil) {
		c.JSON(http.StatusConflict, gin.H{"message": "User id already exist"})
		return
	}
	
	user, err := ctrl.repo.Create(input)
	if (errorHandler(err, c)) { return }
	
	c.JSON(http.StatusCreated, user)
}

func (ctrl *userController) UpdateUser(c *gin.Context) {
	var input models.User
	
	err := c.ShouldBindJSON(&input)
	if (err != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	str_id := strconv.FormatUint(uint64(input.ID), 10)
	
	_, err = ctrl.repo.FindById(str_id)
	if (errorHandler(err, c)) { return }
	
	user, err := ctrl.repo.Update(input)
	if (errorHandler(err, c)) { return }
	
	c.JSON(http.StatusOK, user)	
}

func (ctrl *userController) DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	
	_, err := ctrl.repo.FindById(id)
	if (errorHandler(err, c)) { return }
	
	err = ctrl.repo.Delete(id)	
	if (errorHandler(err, c)) { return }
	
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (ctrl *userController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// Functions
func errorHandler(err error, c *gin.Context) bool {
	if (err != nil) {
		if (err == gorm.ErrRecordNotFound) {
			//c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			c.JSON(http.StatusNotFound, gin.H{"message": "User id not exist"})
			return true
		}
		
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return true
	}
	
	return false	
}
