package service

import (
	"context"
	pb "github.com/imtihon/3/Item_Service/generated/generated/items_service"
	"github.com/imtihon/3/Item_Service/storage/postgres"
	"log/slog"
)

type ItemService struct {
	pb.UnimplementedItemsServiceServer
	Items  postgres.ItemRepo
	Logger *slog.Logger
}

func NewItemService(Item postgres.ItemRepo) *ItemService {
	return &ItemService{Items: Item}
}

func (r *ItemService) CreateItem(ctx context.Context, in *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	return r.Items.CreateItem(ctx, in)
}
func (r *ItemService) UpdateItem(ctx context.Context, in *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {
	return r.Items.UpdateItem(ctx, in)
}
func (r *ItemService) DeleteItem(ctx context.Context, in *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	return r.Items.DeleteItem(ctx, in)
}
func (r *ItemService) GetAllItems(ctx context.Context, in *pb.GetAllItemsRequest) (*pb.GetAllItemsResponse, error) {
	return r.Items.GetAllItems(ctx, in)
}
func (r *ItemService) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	return r.Items.GetItem(ctx, req)
}

func (r *ItemService) SearchItemsAndFilter(ctx context.Context, in *pb.SearchItemsAndFilterRequest) (*pb.SearchItemsAndFilterResponse, error) {
	return r.Items.SearchItemsAndFilter(in)
}

func (r *ItemService) CreateChangeSwaps(ctx context.Context, in *pb.CreateSwapRequest) (*pb.CreateSwapResponse, error) {
	return r.Items.CreateChangeSwaps(in)
}
func (r *ItemService) UpdateAcceptSwap(ctx context.Context, in *pb.UpdateSwapRequest) (*pb.UpdateSwapResponse, error) {
	return r.Items.UpdateAcceptSwap(in)
}

func (r *ItemService) UpdateRejectSwap(ctx context.Context, in *pb.UpdateRejectSwapRequest) (*pb.UpdateRejectSwapResponse, error) {
	return r.Items.UpdateRejectSwap(in)
}
func (r *ItemService) GetChangedSwap(ctx context.Context, in *pb.GetChangeSwapRequest) (*pb.GetChangeSwapResponse, error) {
	return r.Items.GetChangeSwap(in)
}

//func (r *ItemService) CreateAddRecyclingCenter(ctx context.Context, in *pb.CreateRecyclingCenterRequest) (*pb.CreateRecyclingCenterResponse, error) {
//	return r.Items.CreateAddRecyclingCenter(in)
//}

//func (r *ItemService) SearchRecyclingCenter(ctx context.Context, in *pb.SearchRecyclingCenterRequest) (*pb.SearchRecyclingCenterResponse, error) {
//	return r.Items.SearchRecyclingCenter(in)
//}

func (r *ItemService) CreateRecyclingSubmissions(ctx context.Context, in *pb.CreateRecyclingSubmissionsRequest) (*pb.CreateRecyclingSubmissionsResponse, error) {
	return r.Items.CreateRecyclingSubmission(in)
}

func (r *ItemService) CreateRating(ctx context.Context, in *pb.CreateRatingRequest) (*pb.CreateRatingResponse, error) {
	return r.Items.CreateAddUserRating(in)
}
func (r *ItemService) GetUserRatings(ctx context.Context, in *pb.GetRatingRequest) (*pb.GetRatingResponse, error) {
	return r.Items.GetRatings(ctx, in)
}

func (r *ItemService) CreateItemCategory(ctx context.Context, in *pb.CreateItemCategoryManageRequest) (*pb.CreateItemCategoryManageResponse, error) {
	return r.Items.CreateItemCategory(in)
}
func (r *ItemService) GetStatistics(ctx context.Context, in *pb.GetStatisticsRequest) (*pb.GetStatisticsResponse, error) {
	return r.Items.GetStatistics(in)
}
func (r *ItemService) GetMonitoringUserActivity(ctx context.Context, in *pb.GetMonitoringUserActivityRequest) (*pb.GetMonitoringUserActivityResponse, error) {
	return r.Items.GetMonitoringUserActivity(in)
}

func (r *ItemService) CreateEcoChallenge(ctx context.Context, in *pb.CreateEcoChallengeRequest) (*pb.CreateEcoChallengeResponse, error) {
	return r.Items.CreateEcoChallenge(in)
}

func (r *ItemService) CreateParticipateChallenge(ctx context.Context, in *pb.CreateParticipateChallengeRequest) (*pb.CreateParticipateChallengeResponse, error) {
	return r.Items.ParticipateChallenge(in)
}
func (r *ItemService) UpdateEcoChallengeResult(ctx context.Context, in *pb.UpdateEcoChallengeResultRequest) (*pb.UpdateEcoChallengeResultResponse, error) {
	return r.Items.UpdateEcoChallengeResult(in)
}

func (r *ItemService) CreateEcoTip(ctx context.Context, in *pb.CreateEcoTipRequest) (*pb.CreateEcoTipResponse, error) {
	return r.Items.CreateAddEcoTips(in)
}
func (r *ItemService) GetAddEcoTips(ctx context.Context, in *pb.GetAddEcoTipsRequest) (*pb.GetAddEcoTipsResponse, error) {
	return r.Items.GetAddEcoTips(in)
}
