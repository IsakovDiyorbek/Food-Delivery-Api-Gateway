package handler

import (
	pb "github.com/Food_Delivery/Food-Delivery-Api-Gateway/genproto"
	"github.com/gin-gonic/gin"
)



// @Summary Create a new product
// @Description Create a new product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param borrower body pb.CreateProductRequest true "Borrower"
// @Success 200 {object} pb.ProductEmpty
// @Failure 400 {string} string "Bad Request"
// @Router /product [post]
func (h *Handler) CreateProduct(c *gin.Context) {
	product := pb.CreateProductRequest{}
	err := c.BindJSON(&product)

	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	_, err = h.Product.CreateProduct(c, &product)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "Success Create Product")
}
