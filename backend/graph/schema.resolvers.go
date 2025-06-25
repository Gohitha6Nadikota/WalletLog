package graph

import (
	"backend/graph/generated"
	"backend/graph/model"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// Hello
func (r *queryResolver) Hello(ctx context.Context) (string, error) {
    return "Hello, World!", nil
}

// AddExpense
func (r *mutationResolver) AddExpense(ctx context.Context, input model.NewExpense) (*model.Expense, error) {
	userID := getUserIDFromContext(ctx)

	expense := &model.Expense{
		ID:          uuid.NewString(),
		Amount:      input.Amount,
		Category:    input.Category,
		Description: input.Description,
		Date:        input.Date,
		UserID:      userID,
	}

	if err := r.DB.Create(expense).Error; err != nil {
		return nil, err
	}

	return expense, nil
}

// DeleteExpense
func (r *mutationResolver) DeleteExpense(ctx context.Context, id string) (bool, error) {
	userID := getUserIDFromContext(ctx)

	var expense model.Expense
	if err := r.DB.First(&expense, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		return false, fmt.Errorf("expense not found or unauthorized")
	}

	if err := r.DB.Delete(&expense).Error; err != nil {
		return false, err
	}

	return true, nil
}

// UpdateExpense
func (r *mutationResolver) UpdateExpense(ctx context.Context, input model.UpdateExpenseInput) (*model.Expense, error) {
	userID := getUserIDFromContext(ctx)

	var expense model.Expense
	if err := r.DB.First(&expense, "id = ? AND user_id = ?", input.ID, userID).Error; err != nil {
		return nil, fmt.Errorf("expense not found or unauthorized")
	}

	if input.Amount != nil {
		expense.Amount = *input.Amount
	}
	if input.Category != nil {
		expense.Category = *input.Category
	}
	if input.Description != nil {
		expense.Description = *input.Description
	}
	if input.Date != nil {
		expense.Date = *input.Date
	}

	if err := r.DB.Save(&expense).Error; err != nil {
		return nil, err
	}

	return &expense, nil
}

// Register
func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthPayload, error) {
	var existing model.InternalUser

	if err := r.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
		return nil, fmt.Errorf("user already exists")
	}

	hashedPwd, err := hashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &model.InternalUser{
		ID:       uuid.NewString(),
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPwd,
	}

	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}

	token, _ := generateToken(user.ID)

	return &model.AuthPayload{
		Token: token,
		User: &model.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

// Login
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthPayload, error) {
	var user model.InternalUser
	if err := r.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !checkPasswordHash(input.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, _ := generateToken(user.ID)

	return &model.AuthPayload{
		Token: token,
		User: &model.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}


// GetExpenses
func (r *queryResolver) GetExpenses(ctx context.Context, date *string, category *string) ([]*model.Expense, error) {
	userID := getUserIDFromContext(ctx)
	var expenses []*model.Expense
	query := r.DB.Where("user_id = ?", userID)
	if date != nil && *date != "" {
		query = query.Where("date = ?", *date)
	}
	if category != nil && *category != "" {
		query = query.Where("category = ?", *category)
	}
	if err := query.Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}

// ExpenseSummary
func (r *queryResolver) ExpenseSummary(ctx context.Context, startDate string, endDate string) (*model.ExpenseSummary, error) {
	userID := getUserIDFromContext(ctx)

	var expenses []model.Expense
	if err := r.DB.Where("user_id = ? AND date BETWEEN ? AND ?", userID, startDate, endDate).Find(&expenses).Error; err != nil {
		return nil, err
	}

	var (
		totalAmount    float64
		totalCount     int
		categoryTotals = make(map[string]float64)
	)

	for _, exp := range expenses {
		totalAmount += exp.Amount
		totalCount++
		categoryTotals[exp.Category] += exp.Amount
	}

	var categorySummary []*model.CategorySummary
	for category, total := range categoryTotals {
		categorySummary = append(categorySummary, &model.CategorySummary{
			Category: category,
			Total:    total,
		})
	}

	return &model.ExpenseSummary{
		TotalAmount: totalAmount,
		TotalCount:  totalCount,
		ByCategory:  categorySummary,
	}, nil
}

// GetExpenseByID
func (r *queryResolver) GetExpenseByID(ctx context.Context, id string) (*model.Expense, error) {
	userID := getUserIDFromContext(ctx)

	var expense model.Expense
	if err := r.DB.First(&expense, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		return nil, fmt.Errorf("expense not found or unauthorized")
	}

	return &expense, nil
}

// Resolvers
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
