package handlers

import (
	"backend-trainee-assignment-2023/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary CreateSegment
// @Tags segment
// @Description create segment
// @ID create-segment
// @Accept  json
// @Produce  json
// @Param input body models.SegmentRequest true "slug"
// @Success 200 {object} models.CreateSegmentResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /segment [post]
func (h *Handler) CreateSegment(ctx echo.Context) error {
	var req models.SegmentRequest
	if err := ctx.Bind(&req); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	id, err := h.service.Segment.Create(req.Slug)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, models.CreateSegmentResponse{Id: id})
}

// @Summary DeleteSegment
// @Tags segment
// @Description delete segment
// @ID delete-segment
// @Accept  json
// @Produce  json
// @Param input body models.SegmentRequest true "slug"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /segment [delete]
func (h *Handler) DeleteSegment(ctx echo.Context) error {
	var req models.SegmentRequest
	if err := ctx.Bind(&req); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	if err := h.service.Segment.Delete(req.Slug); err != nil {
		return NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
	})
}
