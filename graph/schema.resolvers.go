package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	generated1 "github/jjmengze/test/graph/generated"
	"github/jjmengze/test/graph/model"
	"time"
)

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	first := "nil"
	last := "last"
	return &model.User{
		FirstName: &first,
		LastName:  &last,
	}, nil
}

func (r *queryResolver) Animal(ctx context.Context, name string) (model.Animal, error) {
	//if name == "bird" {
	return &model.Bird{
		Name:           "bird",
		WingSpanLength: 10,
	}, nil
	//} else {
	//	return &model.Monkey{
	//		Name:          "Monkey",
	//		ArmSpanLengt 100,
	//	}, nil
	//}
}

func (r *queryResolver) Animals(ctx context.Context) ([]model.Animal, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Now(ctx context.Context) (*time.Time, error) {
	now := time.Now()
	return &now, nil
}

func (r *queryResolver) IsFriday(ctx context.Context, date time.Time) (*bool, error) {
	isFriday := false
	if date.Weekday() == time.Friday {
		isFriday = true
		return &isFriday, nil
	}
	return &isFriday, nil
}

func (r *queryResolver) Today(ctx context.Context) (*string, error) {
	today := time.Now().String()
	return &today, nil
}

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
