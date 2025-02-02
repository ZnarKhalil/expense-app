package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ZnarKhalil/expense-app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExpenseHandler struct {
	DB *gorm.DB
}

func NewExpenseHandler(db *gorm.DB) *ExpenseHandler {
	return &ExpenseHandler{DB: db}
}

type CreateExpenseRequest struct {
	ExpenseCategoryID uint    `json:"expense_category_id" binding:"required"`
	Amount            float64 `json:"amount" binding:"required"`
	Date              string  `json:"date" binding:"required"` // Expected format: YYYY-MM-DD
	Note              string  `json:"note"`
}

func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	var req CreateExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format; use YYYY-MM-DD"})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}
	u := user.(models.User)

	expense := models.Expense{
		UserID:            u.ID,
		ExpenseCategoryID: req.ExpenseCategoryID,
		Amount:            req.Amount,
		Date:              parsedDate,
		Note:              req.Note,
	}

	if err := h.DB.Create(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, expense)
}

func (h *ExpenseHandler) GetExpenses(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}
	u := user.(models.User)

	var expenses []models.Expense
	if err := h.DB.Where("user_id = ?", u.ID).Find(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

type UpdateExpenseRequest struct {
	ExpenseCategoryID uint    `json:"expense_category_id"`
	Amount            float64 `json:"amount"`
	Date              string  `json:"date"` // Expected format: YYYY-MM-DD
	Note              string  `json:"note"`
}

func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {
	idParam := c.Param("id")
	expenseID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid expense id"})
		return
	}

	var req UpdateExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}
	u := user.(models.User)

	var expense models.Expense
	if err := h.DB.Where("id = ? AND user_id = ?", expenseID, u.ID).First(&expense).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "expense not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if req.ExpenseCategoryID != 0 {
		expense.ExpenseCategoryID = req.ExpenseCategoryID
	}
	if req.Amount != 0 {
		expense.Amount = req.Amount
	}
	if req.Date != "" {
		parsedDate, err := time.Parse("2006-01-02", req.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format; use YYYY-MM-DD"})
			return
		}
		expense.Date = parsedDate
	}
	if req.Note != "" {
		expense.Note = req.Note
	}

	if err := h.DB.Save(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expense)
}

func (h *ExpenseHandler) DeleteExpense(c *gin.Context) {
	idParam := c.Param("id")
	expenseID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid expense id"})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}
	u := user.(models.User)

	if err := h.DB.Where("id = ? AND user_id = ?", expenseID, u.ID).Delete(&models.Expense{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "expense deleted successfully"})
}
