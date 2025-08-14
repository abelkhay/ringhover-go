package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"ringhover-go/internal/daoerrors"
	"ringhover-go/internal/domain/models"
	"ringhover-go/internal/domain/resp"
	"ringhover-go/internal/http/endpoints"
	"ringhover-go/internal/mocks"
)

func setupRouter(h *Handler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	api := r.Group(endpoints.APIBase)
	{
		api.GET(endpoints.TaskSubtasks, h.GetSubtasks)
	}
	return r
}

func TestGetSubTasks_OK(t *testing.T) {
	svc := mocks.NewModelisationServiceInterface(t)
	h := &Handler{service: svc}
	r := setupRouter(h)

	taskID := uint64(7)
	subtasks := []models.Task{
		{Id: 11, Title: "Sous-tâche A", ParentTaskID: &taskID},
		{Id: 12, Title: "Sous-tâche B", ParentTaskID: &taskID},
	}

	svc.EXPECT().
		GetSubTasks(taskID).
		Return(subtasks, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/tasks/"+strconv.FormatUint(taskID, 10)+"/subtasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var subTasks []resp.TaskTree
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &subTasks))

	require.Len(t, subTasks, 2)
	require.Equal(t, "Sous-tâche A", subTasks[0].Title)
	require.Equal(t, "Sous-tâche B", subTasks[1].Title)
}

func TestGetSubTasks_InvalidID(t *testing.T) {
	svc := mocks.NewModelisationServiceInterface(t)
	h := &Handler{service: svc}
	r := setupRouter(h)

	req := httptest.NewRequest(http.MethodGet, "/api/tasks/abc/subtasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetSubTasks_NotFound(t *testing.T) {
	svc := mocks.NewModelisationServiceInterface(t)
	h := &Handler{service: svc}
	r := setupRouter(h)

	taskID := uint64(999)

	svc.EXPECT().
		GetSubTasks(taskID).
		Return(nil, daoerrors.ErrNotFound)

	req := httptest.NewRequest(http.MethodGet, "/api/tasks/"+strconv.FormatUint(taskID, 10)+"/subtasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetSubTasks_InternalError(t *testing.T) {
	svc := mocks.NewModelisationServiceInterface(t)
	h := &Handler{service: svc}
	r := setupRouter(h)

	taskID := uint64(7)

	svc.EXPECT().
		GetSubTasks(taskID).
		Return(nil, errors.New("boom"))

	req := httptest.NewRequest(http.MethodGet, "/api/tasks/"+strconv.FormatUint(taskID, 10)+"/subtasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
}
