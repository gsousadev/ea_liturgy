package mass

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StoreMassController(c *gin.Context) {

	var data MassRequest

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {

		mass, err := storeMassService(&data)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			response := gin.H{"message": fmt.Sprint("Missa Cadastrada com Sucesso ID:", mass.ID)}
			c.JSON(http.StatusOK, response)
		}

	}
}

func DeleteMassController(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	if err := deleteMassService(id); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		response := gin.H{"message": "Missa Excluida com Sucesso"}
		c.JSON(http.StatusOK, response)
	}
}
