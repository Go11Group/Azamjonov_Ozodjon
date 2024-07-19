package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	pb "github.com/imtihon/3/Api_gateway/generated/generated/items_service"
	"net/http"
	"strconv"
)

// CreateItem adds a new item.
// @Summary Add a new item
// @Description Add a new item to the inventory
// @Tags items
// @Accept json
// @Produce json
// @Param item body items_service.CreateItemRequest true "Create Item Request"
// @Success 201 {object} items_service.CreateItemResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/items [post]
func (h *Handler) CreateItem(c *gin.Context) {
	var req pb.CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := h.Items.CreateItem(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, r)
}

// UpdateItem edits an item.
// @Summary Edit an item
// @Description Edit an existing item in the inventory
// @Tags items
// @Accept json
// @Produce json
// @Param item_id path string true "Item ID"
// @Param item body items_service.UpdateItemRequest true "Update Item Request"
// @Success 200 {object} items_service.UpdateItemResponse
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /api/v1/items/{item_id} [put]
func (h *Handler) UpdateItem(c *gin.Context) {
	var req pb.UpdateItemRequest

	// Extract the item ID from the URL parameters
	itemID := c.Param("item_id")
	if itemID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "item_id is required"})
		return
	}

	// Bind the JSON body to the request struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the item ID in the request
	req.Id = itemID

	// Call the UpdateItem method
	r, err := h.Items.UpdateItem(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the updated item
	c.JSON(http.StatusOK, r)
}

// DeleteItem deletes an item.
// @Summary Delete an item
// @Description Delete an existing item from the inventory
// @Tags items
// @Accept json
// @Produce json
// @Param item_id path string true "Item ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /api/v1/items/{item_id} [delete]
func (h *Handler) DeleteItem(c *gin.Context) {
	itemID := c.Param("item_id")

	req := &pb.DeleteItemRequest{
		Id: itemID,
	}

	r, err := h.Items.DeleteItem(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, r)
}

// GetAllItems gets a list of items.
// @Summary Get a list of items
// @Description Get a paginated list of items from the inventory
// @Tags items
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param limit query int true "Items per page"
// @Success 200 {object} items_service.GetItemsRequest
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/items [get]
func (h *Handler) GetAllItems(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	req := &pb.GetAllItemsRequest{
		Page:  int32(page),
		Limit: int32(limit),
	}

	r, err := h.Items.GetAllItems(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, r)
}

// GetByIdItem gets an item by ID.
// @Summary Get an item by ID
// @Description Get an existing item from the inventory by its ID
// @Tags items
// @Accept json
// @Produce json
// @Param item_id path string true "Item ID"
// @Success 200 {object} items_service.GetItemRequest
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /api/v1/items/{item_id} [get]
func (h *Handler) GetByIdItem(c *gin.Context) {
	id := c.Param("item_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	req := &pb.GetItemRequest{Id: id}
	r, err := h.Items.GetItem(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, r)
}

// SearchItemsAndFilter searches and filters items.
// @Summary Search and filter items
// @Description Search and filter items based on query parameters
// @Tags items
// @Accept json
// @Produce json
// @Param query query string false "Search query"
// @Param category query string false "Category"
// @Param condition query string false "Condition"
// @Param page query int true "Page number"
// @Param limit query int true "Items per page"
// @Success 200 {object} items_service.SearchItemsAndFilterResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/items/search [get]
func (h *Handler) SearchItemsAndFilter(c *gin.Context) {
	req := pb.SearchItemsAndFilterRequest{
		Query:     c.Query("query"),
		Category:  c.Query("category"),
		Condition: c.Query("condition"),
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err == nil {
		req.Page = int32(page)
	} else {
		req.Page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err == nil {
		req.Limit = int32(limit)
	} else {
		req.Limit = 10
	}

	r, err := h.Items.SearchItemsAndFilter(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, r)
}

// CreateChangeSwaps creates a new swap request.
// @Summary Create a new swap request
// @Description Create a new swap request in the items service.
// @Tags swaps
// @Accept json
// @Produce json
// @Param swap body items_service.CreateSwapRequest true "Create Swap Request"
// @Success 201 {object} items_service.CreateSwapResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /swaps [post]
func (h *Handler) CreateChangeSwaps(c *gin.Context) {
	var req pb.CreateSwapRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := h.Items.CreateChangeSwaps(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, r)
}

// UpdateAcceptSwap accepts a swap request.
// @Summary Accept a swap request
// @Description Accept a swap request in the items service.
// @Tags swaps
// @Accept json
// @Produce json
// @Param swap body items_service.UpdateSwapRequest true "Update Swap Request"
// @Success 200 {object} items_service.UpdateSwapResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /swaps/accept [put]
func (h *Handler) UpdateAcceptSwap(c *gin.Context) {
	var req pb.UpdateSwapRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log the request for debugging
	fmt.Printf("Received UpdateAcceptSwap request: %+v\n", req)

	// Validate that the UUID fields are not empty
	if req.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Swap ID is required"})
		return
	}

	// Ensure h.Items is initialized and not nil
	if h.Items == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Items service is not initialized"})
		return
	}

	// Handle the request using h.Items
	r, err := h.Items.UpdateAcceptSwap(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, r)
}

// UpdateRejectSwap rejects a swap request.
// @Summary Reject a swap request
// @Description Reject a swap request in the items service.
// @Tags swaps
// @Accept json
// @Produce json
// @Param swap body items_service.UpdateRejectSwapRequest true "Update Reject Swap Request"
// @Success 200 {object} items_service.UpdateRejectSwapResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /swaps/reject [put]
func (h *Handler) UpdateRejectSwap(c *gin.Context) {
	var req pb.UpdateRejectSwapRequest

	// Extract swap_id from the URL
	swapId := c.Param("swap_id")
	if swapId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Swap ID is required"})
		return
	}

	// Bind the JSON body to the request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the swap_id from the URL parameter
	req.SwapId = swapId

	// Log the request for debugging
	fmt.Printf("Received UpdateRejectSwap request: %+v\n", req)

	// Handle the request using h.Items
	r, err := h.Items.UpdateRejectSwap(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, r)
}

// GetChangedSwap gets a swap request by ID.
// @Summary Get a swap request by ID
// @Description Get a swap request by ID from the items service.
// @Tags swaps
// @Accept json
// @Produce json
// @Param swap body items_service.GetChangeSwapRequest true "Get Change Swap Request"
// @Success 200 {object} items_service.GetChangeSwapResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /swaps/{id} [get]
func (h *Handler) GetChangedSwap(c *gin.Context) {
	var req pb.GetChangeSwapRequest

	// Extract query parameters
	req.Status = c.Query("status")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}
	req.Page = int32(page)

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}
	req.Limit = int32(limit)

	// Handle the request using h.Items
	r, err := h.Items.GetChangedSwap(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, r)
}

// CreateAddRecyclingCenter adds a new recycling center.
// @Summary Add a new recycling center
// @Description Add a new recycling center in the items service.
// @Tags recycling_centers
// @Accept json
// @Produce json
// @Param recycling_center body items_service.CreateRecyclingCenterRequest true "Create Recycling Center Request"
// @Success 201 {object} items_service.CreateRecyclingCenterResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /recycling_centers [post]
func (h *Handler) CreateAddRecyclingCenter(c *gin.Context) {
	var req pb.CreateRecyclingCenterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure the repository is initialized and not nil
	if h.Items == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Items service is not initialized"})
		return
	}

	// Create request object for repository
	repositoryRequest := &pb.CreateRecyclingCenterRequest{
		Name:              req.Name,
		Address:           req.Address,
		AcceptedMaterials: req.AcceptedMaterials,
		WorkingHours:      req.WorkingHours,
		ContactNumber:     req.ContactNumber,
	}

	// Call repository method to create recycling center
	r, err := h.Items.CreateAddRecyclingCenter(c, repositoryRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create recycling center: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, r)
}

// SearchRecyclingCenter searches for recycling centers.
// @Summary Search for recycling centers
// @Description Search for recycling centers in the items service.
// @Tags recycling_centers
// @Accept json
// @Produce json
// @Param recycling_center body items_service.SearchRecyclingCenterRequest true "Search Recycling Center Request"
// @Success 200 {object} items_service.SearchRecyclingCenterResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /recycling_centers/search [post]
func (h *Handler) SearchRecyclingCenter(c *gin.Context) {
	var req pb.SearchRecyclingCenterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := h.Items.SearchRecyclingCenter(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, r)
}

// CreateRecyclingSubmission creates a new recycling submission.
// @Summary Create a new recycling submission
// @Description Create a new recycling submission in the items service.
// @Tags recycling_submissions
// @Accept json
// @Produce json
// @Param recycling_submission body items_service.CreateRecyclingSubmissionsRequest true "Create Recycling Submissions Request"
// @Success 201 {object} items_service.CreateRecyclingSubmissionsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /recycling_submissions [post]
func (h *Handler) CreateRecyclingSubmission(c *gin.Context) {
	var req pb.CreateRecyclingSubmissionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := h.Items.CreateRecyclingSubmissions(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, r)
}

// CreateAddUserRating adds a new user rating.
// @Summary Add a new user rating
// @Description Add a new user rating in the items service.
// @Tags user_ratings
// @Accept json
// @Produce json
// @Param rating body items_service.CreateRatingRequest true "Create Rating Request"
// @Success 201 {object} items_service.CreateRatingResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user_ratings [post]
func (h *Handler) CreateAddUserRating(c *gin.Context) {
	var req pb.CreateRatingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := h.Items.CreateRating(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, r)
}

// GetUserRatings gets user ratings.
// @Summary Get user ratings
// @Description Get user ratings from the items service.
// @Tags user_ratings
// @Accept json
// @Produce json
// @Param rating body items_service.GetRatingRequest true "Get Rating Request"
// @Success 200 {object} items_service.GetRatingResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user_ratings [get]
func (h *Handler) GetUserRatings(c *gin.Context) {
	// Parse request parameters
	var req pb.GetRatingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call repository method to fetch ratings
	r, err := h.Items.GetUserRatings(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate average rating
	var totalRatings int32
	var totalRatingSum float64
	for _, rating := range r.Ratings {
		totalRatings++
		totalRatingSum += rating.Rating
	}
	averageRating := float64(0.0)
	if totalRatings > 0 {
		averageRating = totalRatingSum / float64(totalRatings)
	}

	// Prepare response
	response := map[string]interface{}{
		"ratings":        r.Ratings,
		"average_rating": averageRating,
		"total_ratings":  totalRatings,
		"page":           req.Page,
		"limit":          req.Limit,
	}

	c.JSON(http.StatusOK, response)
}

// CreateItemCategory creates a new item category.
// @Summary Create a new item category
// @Description Create a new item category in the items service.
// @Tags item_categories
// @Accept json
// @Produce json
// @Param item_category body items_service.CreateItemCategoryManageRequest true "Create Item Category Request"
// @Success 201 {object} items_service.CreateItemCategoryManageResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /item_categories [post]
func (h *Handler) CreateItemCategory(c *gin.Context) {
	var req pb.CreateItemCategoryManageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := h.Items.CreateItemCategory(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, r)
}

// GetStatistics gets statistics.
// @Summary Get statistics
// @Description Get statistics from the items service.
// @Tags statistics
// @Accept json
// @Produce json
// @Param statistics body items_service.GetStatisticsRequest true "Get Statistics Request"
// @Success 200 {object} items_service.GetStatisticsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /statistics [post]
//func (h *Handler) GetStatistics(c *gin.Context) {
//	var req pb.GetStatisticsRequest
//
//	start_date := c.Query("start_date")
//	if len(start_date) != 0 {
//		req.StartDate = start_date
//	}
//
//	end_date := c.Query("end_date")
//	if len(end_date) != 0 {
//		req.EndDate = end_date
//	}
//
//	r, err := h.Items.GetStatistics(context.Background(), &req)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, r)
//}

// GetMonitoringUserActivity gets user activity monitoring data.
// @Summary Get user activity monitoring data
// @Description Get user activity monitoring data from the items service.
// @Tags user_activity
// @Accept json
// @Produce json
// @Param activity body items_service.GetMonitoringUserActivityRequest true "Get Monitoring User Activity Request"
// @Success 200 {object} items_service.GetMonitoringUserActivityResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user_activity [post]
func (h *Handler) GetMonitoringUserActivity(c *gin.Context) {
	// Extract user_id from URL path
	userID := c.Param("user_id")

	// Extract query parameters
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	// Prepare request for gRPC call
	req := &pb.GetMonitoringUserActivityRequest{
		UserId:    userID,
		StartDate: startDate,
		EndDate:   endDate,
	}

	// Call gRPC service method
	r, err := h.Items.GetMonitoringUserActivity(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return JSON response
	c.JSON(http.StatusOK, r)
}

// CreateEcoChallenges creates a new eco challenge.
// @Summary Create a new eco challenge
// @Description Create a new eco challenge in the items service.
// @Tags eco_challenges
// @Accept json
// @Produce json
// @Param eco_challenge body items_service.CreateEcoChallengeRequest true "Create Eco Challenge Request"
// @Success 201 {object} items_service.CreateEcoChallengeResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /eco_challenges [post]
func (h *Handler) CreateEcoChallenges(c *gin.Context) {
	var req pb.CreateEcoChallengeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := h.Items.CreateEcoChallenge(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, r)
}

// ParticipateChallenge participates in an eco challenge.
// @Summary Participate in an eco challenge
// @Description Participate in an eco challenge in the items service.
// @Tags eco_challenges
// @Accept json
// @Produce json
// @Param participate body items_service.CreateParticipateChallengeRequest true "Create Participate Challenge Request"
// @Success 201 {object} items_service.CreateParticipateChallengeResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /eco_challenges/participate [post]
func (h *Handler) ParticipateChallenge(c *gin.Context) {
	// Extract challenge_id from URL path
	challengeID := c.Param("challenge_id")

	// Parse JSON request body into struct
	var req pb.CreateParticipateChallengeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the challenge_id in the request
	req.ChallengeId = challengeID

	// Call gRPC service method
	r, err := h.Items.CreateParticipateChallenge(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return JSON response with status 201 Created
	c.JSON(http.StatusCreated, r)
}

// UpdateEcoChallengeResult updates the result of an eco challenge.
// @Summary Update the result of an eco challenge
// @Description Update the result of an eco challenge in the items service.
// @Tags eco_challenges
// @Accept json
// @Produce json
// @Param result body items_service.UpdateEcoChallengeResultRequest true "Update Eco Challenge Result Request"
// @Success 200 {object} items_service.UpdateEcoChallengeResultResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /eco_challenges/result [put]
func (h *Handler) UpdateEcoChallengeResult(c *gin.Context) {
	var req pb.UpdateEcoChallengeResultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := h.Items.UpdateEcoChallengeResult(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, r)
}

// CreateAddEcoTips adds new eco tips.
// @Summary Add new eco tips
// @Description Add new eco tips in the items service.
// @Tags eco_tips
// @Accept json
// @Produce json
// @Param eco_tips body items_service.CreateEcoTipRequest true "Create Eco Tip Request"
// @Success 201 {object} items_service.CreateEcoTipResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /eco_tips [post]
func (h *Handler) CreateAddEcoTips(c *gin.Context) {
	var req pb.CreateEcoTipRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r, err := h.Items.CreateEcoTip(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, r)
}

// GetAddEcoTips gets eco tips.
// @Summary Get eco tips
// @Description Get eco tips from the items service.
// @Tags eco_tips
// @Accept json
// @Produce json
// @Param eco_tips body items_service.GetAddEcoTipsRequest true "Get Add Eco Tips Request"
// @Success 200 {object} items_service.GetAddEcoTipsResponse
// @Failure 400 {object}  string
// @Failure 500 {object} string
// @Router /eco_tips [get]
func (h *Handler) GetAddEcoTips(c *gin.Context) {
	var req pb.GetAddEcoTipsRequest

	limit := c.Query("limit")
	if len(limit) != 0 {
		atoi, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		req.Limit = int32(atoi)
	}

	page := c.Query("page")
	if len(page) != 0 {
		atoi, err := strconv.Atoi(page)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		req.Page = int32(atoi)
	}
	r, err := h.Items.GetAddEcoTips(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, r)
}
