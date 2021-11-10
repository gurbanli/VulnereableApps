package controller

import (
	"SQLInjection/model"
	"SQLInjection/repository"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ProductController struct{

}

func (pc *ProductController)GetProduct(c *gin.Context){
	sess := sessions.Default(c)
	var product model.Product
	if sess.Get("id") != nil {

		productId := c.Param("id")
		db := repository.Database

		rows, err := db.Query(fmt.Sprintf("SELECT * FROM products WHERE id='%s'", productId))
		if err != nil{
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
			})
			return
		}else{
				for rows.Next(){
				err = rows.Scan(&product.Id, &product.ProductName, &product.ProductType, &product.Count)
				if err != nil{
					log.Println(err)
					c.JSON(http.StatusBadRequest, gin.H{
						"message": "Bad Request",
					})
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"message":"Successfully retrieved!",
					"product": product,
				})}
		}
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Please Login",
		})
	}
}

