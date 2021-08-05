package resolver

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/layer5io/meshery/handlers"
	graphqlModels "github.com/layer5io/meshery/internal/graphql/model"
	"github.com/layer5io/meshery/models"
)

// func (r *Resolver) subscribePerfResults(ctx context.Context, provider models.Provider, filter *graphqlModels.PageFilter) (<-chan *graphqlModels.PerfPageResult, error) {
// 	if r.performanceChannel == nil {
// 		r.performanceChannel = make(chan *graphqlModels.PerfPageResult)
// 		r.operatorSyncChannel = make(chan struct{})
// 	}

// 	go func() {
// 		r.Log.Info("Performance subscription started")

// 		tokenString := ctx.Value("token").(string)

// 		provider.FetchAllResults()

// 		for {
// 			select {
// 			case <-r.operatorSyncChannel:
// 				status, err := r.getOperatorStatus(ctx, provider)
// 				if err != nil {
// 					r.Log.Error(ErrOperatorSubscription(err))
// 					return
// 				}
// 				r.performanceChannel <- status
// 			case <-ctx.Done():
// 				r.Log.Info("Operator subscription flushed")
// 				return
// 			}
// 		}
// 	}()

// 	return r.performanceChannel, nil
// }

func (r *Resolver) getPerfResult(ctx context.Context, provider models.Provider, id *string) (*graphqlModels.MesheryResult, error) {
	if *id == "" {
		return nil, handlers.ErrQueryGet("*id")
	}

	resultID, err := uuid.FromString(*id)

	if err != nil {
		r.Log.Error(err)
		return nil, err
	}

	tokenString := ctx.Value("token").(string)

	bdr, err := provider.GetResult(tokenString, resultID)

	if err != nil {
		r.Log.Error(err)
		return nil, err
	}

	startTime := int(bdr.TestStartTime.Unix())
	serverBoardConfig := fmt.Sprintf("%v", bdr.ServerBoardConfig)
	serverMetrics := fmt.Sprintf("%v", bdr.ServerMetrics)
	// runnerResults := fmt.Sprintf("%v", bdr.Result)
	mesheryId := fmt.Sprintf("%v", bdr.ID)
	performanceProfile := fmt.Sprintf("%v", bdr.PerformanceProfileInfo.ID)

	return &graphqlModels.MesheryResult{
		MesheryID:          &mesheryId,
		Name:               &bdr.Name,
		Mesh:               &bdr.Mesh,
		PerformanceProfile: &performanceProfile,
		TestID:             &bdr.TestID,
		RunnerResults:      bdr.Result,
		ServerMetrics:      &serverMetrics,
		ServerBoardConfig:  &serverBoardConfig,
		TestStartTime:      &startTime,
		UserID:             &bdr.UserID,
		UpdatedAt:          &bdr.UpdatedAt,
		CreatedAt:          &bdr.CreatedAt,
	}, nil

}
