package controller

import (
	"encoding/json"
	"net/http"

	"github.com/furkancosgun/expense-tracker-api/internal/common"
	"github.com/furkancosgun/expense-tracker-api/internal/dto"
	"github.com/furkancosgun/expense-tracker-api/internal/helper"
	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/furkancosgun/expense-tracker-api/internal/service"
	"github.com/google/uuid"
)

type ProjectController struct {
	service service.IProjectService
}

func NewProjectController(service service.IProjectService) *ProjectController {
	return &ProjectController{service: service}
}

func (controller *ProjectController) GetProjectReportByUserId(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(common.AUTH_USER_ID).(string)

	projects, err := controller.service.GetProjectReportByUserId(userId)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	helper.JsonWriteToResponse(w, projects, http.StatusOK)
}
func (controller *ProjectController) CreateProject(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(common.AUTH_USER_ID).(string)

	var dto dto.CreateProjectRequest

	//Json Decode
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	model := model.Project{
		ProjectId: uuid.New().String(),
		Name:      dto.Name,
		UserId:    userId,
	}

	err = controller.service.CreateProject(model)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
