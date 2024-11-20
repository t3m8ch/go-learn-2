package products

import (
	"context"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
	"github.com/t3m8ch/go-learn-2/internal/api"
	"github.com/t3m8ch/go-learn-2/internal/db"
	"go.uber.org/zap"
)

func SetupRoutes(r *gin.Engine, q db.Querier, logger *zap.Logger) {
	products := r.Group("/products")

	products.GET("", func(c *gin.Context) {
		products, err := q.GetAllProducts(context.Background())

		if err != nil {
			c.JSON(500, api.InternalServerError)
			logger.Error(err.Error())
			return
		}

		c.JSON(200, products)
	})

	products.GET("/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)

		if err != nil {
			c.JSON(404, api.NotFoundError)
			return
		}

		product, err := q.GetProductById(context.Background(), id)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				c.JSON(404, api.NotFoundError)
			} else {
				c.JSON(500, api.InternalServerError)
				logger.Error(err.Error())
			}
			return
		}

		c.JSON(200, product)
	})

	products.POST("", func(c *gin.Context) {
		var json CreateUpdateProductScheme
		if err := c.ShouldBindBodyWithJSON(&json); err != nil {
			c.JSON(400, api.InvalidJsonError)
			return
		}

		price, err := decimal.NewFromString(json.Price)
		if err != nil {
			c.JSON(400, api.InvalidJsonError)
			return
		}

		product, err := q.CreateProduct(context.Background(), db.CreateProductParams{
			Title:       json.Title,
			Description: json.Description,
			Price:       price,
		})

		if err != nil {
			c.JSON(500, api.InternalServerError)
			logger.Error(err.Error())
			return
		}

		c.JSON(201, product)
	})

	products.PUT("/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)

		if err != nil {
			c.JSON(404, api.NotFoundError)
			return
		}

		var json CreateUpdateProductScheme
		if err := c.ShouldBindBodyWithJSON(&json); err != nil {
			c.JSON(400, api.InvalidJsonError)
			return
		}

		price, err := decimal.NewFromString(json.Price)
		if err != nil {
			c.JSON(400, api.InvalidJsonError)
			return
		}

		err = q.UpdateProduct(context.Background(), db.UpdateProductParams{
			ID:          id,
			Title:       json.Title,
			Description: json.Description,
			Price:       price,
		})

		if err != nil {
			c.JSON(500, api.InternalServerError)
			logger.Error(err.Error())
			return
		}

		c.JSON(200, "OK")
	})

	products.DELETE("/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)

		if err != nil {
			c.JSON(404, api.NotFoundError)
			return
		}

		err = q.DeleteProduct(context.Background(), id)
		if err != nil {
			c.JSON(500, api.InternalServerError)
			logger.Error(err.Error())
			return
		}

		c.JSON(200, "OK")
	})
}
