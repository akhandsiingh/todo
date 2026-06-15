package controller

import (
    "net/http"
    "strconv"
    "todo-app/backend/internal/middleware"
    "todo-app/backend/internal/model"
    "todo-app/backend/internal/service"
    "todo-app/backend/internal/util"
)

type TaskController struct { service *service.TaskService }
func NewTaskController(s *service.TaskService) *TaskController { return &TaskController{service:s} }
func (c *TaskController) List(w http.ResponseWriter, r *http.Request) { limit,offset:=util.Pagination(r); gid,_:=strconv.ParseInt(r.URL.Query().Get("group_id"),10,64); res,err:=c.service.List(r.Context(), middleware.UserID(r), r.URL.Query().Get("status"), gid, limit, offset); if err!=nil { util.Error(w,500,err.Error()); return }; util.JSON(w,200,res) }
func (c *TaskController) Create(w http.ResponseWriter, r *http.Request) { var req model.TaskRequest; if err:=util.Decode(r,&req); err!=nil { util.Error(w,400,"invalid json"); return }; res,err:=c.service.Create(r.Context(), middleware.UserID(r), req); if err!=nil { util.Error(w,400,err.Error()); return }; util.JSON(w,201,res) }
func (c *TaskController) Update(w http.ResponseWriter, r *http.Request) { id,ok:=pathID(w,r,"/api/tasks/"); if !ok { return }; var req model.TaskRequest; if err:=util.Decode(r,&req); err!=nil { util.Error(w,400,"invalid json"); return }; res,err:=c.service.Update(r.Context(), middleware.UserID(r), id, req); if err!=nil { util.Error(w,400,err.Error()); return }; util.JSON(w,200,res) }
func (c *TaskController) Delete(w http.ResponseWriter, r *http.Request) { id,ok:=pathID(w,r,"/api/tasks/"); if !ok { return }; if err:=c.service.Delete(r.Context(), middleware.UserID(r), id); err!=nil { util.Error(w,400,err.Error()); return }; util.JSON(w,200,model.MessageResponse{Message:"task deleted"}) }
