package router

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"vote.app/m/db"
	"vote.app/m/graph/model"
)

func VoteSelect(ctx context.Context, input model.VoteSelect) (*model.Result, error) {
	dsn := db.GetGormDB()
	var user db.User
	var vote db.Vote
	err := dsn.First(&user, db.User{STUDENT_NUMBER: int64(input.StudentNumber)})
	if err.Error != nil {
		// Handle error if user not found
		errList := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "ไม่มีผู้ใช้ที่ตรงกับข้อมูลที่ส่งมา",
			Extensions: map[string]interface{}{
				"code": "NOT_FOUND",
			},
		}
		return &model.Result{Success: false}, errList
	}
	if user.IsVoted {
		errList := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "คุณได้ทำการลงคะแนนไปแล้ว",
			Extensions: map[string]interface{}{
				"code": "FORBIDDEN",
			},
		}
		return &model.Result{Success: false}, errList
	}
	err = dsn.First(&vote, input.ID)
	if err.Error != nil {
		errList := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "เกิดปัญหาในการลงคะแนน",
			Extensions: map[string]interface{}{
				"code": "์NOT_FOUND",
			},
		}
		return &model.Result{Success: false}, errList
	}
	vote.Score = vote.Score + 1
	err = dsn.Updates(&vote)
	if err.Error != nil {
		errList := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "เกิดปัญหาในการลงคะแนน",
			Extensions: map[string]interface{}{
				"code": "INTERNAL_SERVER_ERROR",
			},
		}
		return &model.Result{Success: false}, errList
	}
	user.IsVoted = true
	err = dsn.Updates(&user)
	if err.Error != nil {
		errList := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "เกิดปัญหาในการลงคะแนน",
			Extensions: map[string]interface{}{
				"code": "INTERNAL_SERVER_ERROR",
			},
		}
		return &model.Result{Success: false}, errList
	}

	return &model.Result{Success: true}, nil
}
func VoteList(ctx context.Context) ([]*model.VoteList, error) {
	dsn := db.GetGormDB()
	var VoteList []db.Vote
	err := dsn.Find(&VoteList)
	if err.Error != nil {
		errList := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "เกิดปัญหาในการดึงข้อมูลผู้ลงสมัคร",
			Extensions: map[string]interface{}{
				"code": "INTERNAL_SERVER_ERROR",
			},
		}
		return nil, errList
	}
	var modelVoteList []*model.VoteList // Initialize as an empty slice of pointers
	for _, vote := range VoteList {
		modelVote := &model.VoteList{ // Create a pointer to the struct
			ID:      int(vote.ID),
			Name:    vote.Name,
			Number:  int(vote.Number),
			Details: vote.Details,
			LogoURL: vote.LogoUrl,
			Score:   int(vote.Score),
		}
		modelVoteList = append(modelVoteList, modelVote)
	}

	return modelVoteList, nil
}

func CreateVote(ctx context.Context, input model.CreateVote) (*model.Result, error) {
	dsn := db.GetGormDB()
	vote := db.Vote{
		Name:    input.Name,
		Details: input.Details,
		LogoUrl: input.LogoURL,
		Score:   0,
	}
	err := dsn.Create(&vote)
	if err.Error != nil {
		errList := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "เกิดปัญหาในการลงคะแนน",
			Extensions: map[string]interface{}{
				"code": "INTERNAL_SERVER_ERROR",
			},
		}
		return &model.Result{Success: false}, errList
	}
	return &model.Result{Success: true}, nil
}

func EditVote(ctx context.Context, input model.EditVote) (*model.Result, error) {
	dsn := db.GetGormDB()
	vote := db.Vote{
		Name:    input.Name,
		Details: input.Details,
		LogoUrl: input.LogoURL,
	}
	err := dsn.Updates(&vote)
	if err.Error != nil {
		errList := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "เกิดปัญหาในการลงคะแนน",
			Extensions: map[string]interface{}{
				"code": "INTERNAL_SERVER_ERROR",
			},
		}
		return &model.Result{Success: false}, errList
	}
	return &model.Result{Success: true}, nil
}

func DeleteVote(ctx context.Context, input model.DeleteVote) (*model.Result, error) {
	dsn := db.GetGormDB()
	err := dsn.Delete(&db.Vote{}, input.ID)
	if err.Error != nil {
		errList := &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: "เกิดปัญหาในการลงคะแนน",
			Extensions: map[string]interface{}{
				"code": "INTERNAL_SERVER_ERROR",
			},
		}
		return &model.Result{Success: false}, errList
	}
	return &model.Result{Success: true}, nil
}
