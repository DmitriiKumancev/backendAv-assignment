package handlers

import (
	"backend-trainee-assignment-2023/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary ManageUserToSegments
// @Tags users-segment
// @Description add and remove segments from user
// @ID manage-user-to-segments
// @Accept  json
// @Produce  json
// @Param input body models.ManageUserToSegmentsRequest true "slugs to add and remove, user id"
// @Success 200 {object} models.ManageUserToSegmentsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users-segments [post]
func (h *Handler) ManageUserToSegments(ctx echo.Context) error {
	var req models.ManageUserToSegmentsRequest
	if err := ctx.Bind(&req); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	resp, err := h.service.ManageUserToSegments(req.SlugsToAdd, req.SlugsToRemove, req.UserId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, models.ManageUserToSegmentsResponse{
		SlugsHaveBeenAdded:   resp.SlugsHaveBeenAdded,
		SlugsHaveBeenRemoved: resp.SlugsHaveBeenRemoved,
		UserId:               resp.UserId,
	})
}

// @Summary GetUserSegments
// @Tags users-segment
// @Description get all user segments
// @ID get-user-segments
// @Accept  json
// @Produce  json
// @Param input body models.GetUserSegmentsRequest true "user id"
// @Success 200 {object} models.GetUserSegmentsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users-segments [get]
func (h *Handler) GetUserSegments(ctx echo.Context) error {
	var req models.GetUserSegmentsRequest
	if err := ctx.Bind(&req); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	slugs, err := h.service.User.GetUserSegments(req.UserId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, models.GetUserSegmentsResponse{Slugs: slugs})
}
