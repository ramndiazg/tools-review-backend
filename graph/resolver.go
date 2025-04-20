package graph

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"
	"tools-review-backend/ent"
	"tools-review-backend/graph/generated"
	"tools-review-backend/graph/model"

	"github.com/google/uuid"
)

type Resolver struct{
	client *ent.Client
}

func NewResolver(client *ent.Client) *Resolver {
	return &Resolver{client: client}
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*ent.User, error) {
	return r.client.User.
		Create().
		SetName(input.Name).
		SetUsername(input.Username).
		SetEmail(input.Email).
		SetPasswordHash(input.Password).
		Save(ctx)
}

// CreateTool is the resolver for the createTool field.
func (r *mutationResolver) CreateTool(ctx context.Context, input model.CreateToolInput) (*ent.Tool, error) {
	return r.client.Tool.
		Create().
		SetName(input.Name).
		SetDescription(*input.Description).
		SetCategory(*input.Category).
		SetWebsite(*input.Website).
		SetImageURL(*input.ImageURL).
		Save(ctx)
}

// CreateReview is the resolver for the createReview field.
func (r *mutationResolver) CreateReview(ctx context.Context, input model.CreateReviewInput) (*ent.Review, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, err
	}

	toolID, err := uuid.Parse(input.ToolID)
	if err != nil {
		return nil, err
	}

	return r.client.Review.
		Create().
		SetRating(input.Rating).
		SetComment(input.Comment).
		SetUserID(userID).
		SetToolID(toolID).
		Save(ctx)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*ent.User, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return r.client.User.Get(ctx, userID)
}

// Tools is the resolver for the tools field.
func (r *queryResolver) Tools(ctx context.Context) ([]*ent.Tool, error) {
	return r.client.Tool.Query().All(ctx)
}

// Tool is the resolver for the tool field.
func (r *queryResolver) Tool(ctx context.Context, id string) (*ent.Tool, error) {
	toolID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return r.client.Tool.Get(ctx, toolID)
}

// Reviews is the resolver for the reviews field.
func (r *queryResolver) Reviews(ctx context.Context) ([]*ent.Review, error) {
	return r.client.Review.Query().All(ctx)
}

// Review is the resolver for the review field.
func (r *queryResolver) Review(ctx context.Context, id string) (*ent.Review, error) {
	reviewID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return r.client.Review.Get(ctx, reviewID)
}

// ID is the resolver for the id field.
func (r *reviewResolver) ID(ctx context.Context, obj *ent.Review) (string, error) {
	return obj.ID.String(), nil
}

// User is the resolver for the user field.
func (r *reviewResolver) User(ctx context.Context, obj *ent.Review) (*ent.User, error) {
	return obj.QueryUser().Only(ctx)
}

// Tool is the resolver for the tool field.
func (r *reviewResolver) Tool(ctx context.Context, obj *ent.Review) (*ent.Tool, error) {
	return obj.QueryTool().Only(ctx)
}

// ID is the resolver for the id field.
func (r *toolResolver) ID(ctx context.Context, obj *ent.Tool) (string, error) {
	return obj.ID.String(), nil
}

// Reviews is the resolver for the reviews field.
func (r *toolResolver) Reviews(ctx context.Context, obj *ent.Tool) ([]*ent.Review, error) {
	return obj.QueryReviews().All(ctx)
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	return obj.ID.String(), nil
}

// Reviews is the resolver for the reviews field.
func (r *userResolver) Reviews(ctx context.Context, obj *ent.User) ([]*ent.Review, error) {
	return obj.QueryReviews().All(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Review returns generated.ReviewResolver implementation.
func (r *Resolver) Review() generated.ReviewResolver { return &reviewResolver{r} }

// Tool returns generated.ToolResolver implementation.
func (r *Resolver) Tool() generated.ToolResolver { return &toolResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type reviewResolver struct{ *Resolver }
type toolResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
