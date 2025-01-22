package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "auth/genproto/register"
)

func (h *Handler) CreateRegister(c *gin.Context) {
	req := pb.CreateRegisterRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.AuthUser.CreateRegister(context.Background(), &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Foydalanuvchi malumotlarni bazga yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

func (h *Handler) Update(c *gin.Context) {
	req := pb.UpdateRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.AuthUser.Update(context.Background(), &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Foydalanuvchi malumotlarni bazga yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

func (h *Handler) AddImage(c *gin.Context) {
	req := pb.AddImageRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.AuthUser.AddImage(context.Background(), &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Foydalanuvchi malumotlarni bazga yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

func (h *Handler) GetRegister(c *gin.Context) {
	req := pb.GetRegisterRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.AuthUser.GetRegister(context.Background(), &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Foydalanuvchi malumotlarni bazdan olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) Login(c *gin.Context) {
	req := pb.LoginRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req1 := pb.GetByEmailRequest{
		Email: req.Email,
	}

	resp, err := h.AuthUser.GetByEmail(context.Background(), &req1)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Foydalanuvchi malumotlarni bazdan olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if req.Password != resp.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Parol noto'g'ri"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdatePassword(c *gin.Context) {
	req := pb.UpdatePasswordRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.AuthUser.UpdatePassword(context.Background(), &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Foydalanuvchi malumotlarni bazdan olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}