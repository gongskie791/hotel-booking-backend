package handler

import (
	"booking-system/lvl1/internal/model"
	"booking-system/lvl1/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HotelHandler struct {
	hotelService *service.HotelService
}

func NewHotelHandler(hotelService *service.HotelService) *HotelHandler {
	return &HotelHandler{hotelService: hotelService}
}

func (h *HotelHandler) CreateHotel(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req model.CreateHotelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.UserID = userID.(string)

	if err := h.hotelService.CreateHotel(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Hotel created successfully"})
}

func (h *HotelHandler) GetHotelByID(c *gin.Context) {
	id := c.Param("id")
	hotel, err := h.hotelService.GetHotelByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}
	c.JSON(http.StatusOK, hotel)
}

func (h *HotelHandler) GetAllHotels(c *gin.Context) {
	hotels, err := h.hotelService.GetAllHotels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, hotels)
}

func (h *HotelHandler) UpdateHotel(c *gin.Context) {
	id := c.Param("id")
	var req model.UpdateHotelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.hotelService.UpdateHotel(id, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Hotel updated successfully"})
}
