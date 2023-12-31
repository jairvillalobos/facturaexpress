package handlers

import (
	"facturaexpress/common"
	"facturaexpress/data"
	"facturaexpress/helpers"
	"facturaexpress/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register maneja el registro de usuarios.
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponseInit(common.ErrJSONBindingFailed, "Error al procesar los datos del usuario."))
		return
	}

	db := data.GetInstance()

	if err := helpers.CheckUsernameEmail(db, user.Username, user.Email); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := helpers.CheckRoleExists(db, common.USER); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponseInit(common.ErrJSONBindingFailed, "Error al hashear la contraseña."))
		return
	}

	userID, err := helpers.SaveUser(db, user.Username, hashedPassword, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := helpers.SaveUserRole(db, userID); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario registrado con éxito."})
}
