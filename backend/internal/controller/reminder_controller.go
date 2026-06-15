package controller

import (
    "net/http"
    "todo-app/backend/internal/middleware"
    "todo-app/backend/internal/model"
    "todo-app/backend/internal/service"
    "todo-app/backend/internal/util"
)

type ReminderController struct { service *service.ReminderService }
func NewReminderController(s *service.ReminderService) *ReminderController { return &ReminderController{service:s} }
func (c *ReminderController) List(w http.ResponseWriter, r *http.Request) { res,err:=c.service.List(r.Context(), middleware.UserID(r)); if err!=nil { util.Error(w,500,err.Error()); return }; util.JSON(w,200,res) }
func (c *ReminderController) Create(w http.ResponseWriter, r *http.Request) { var req model.ReminderRequest; if err:=util.Decode(r,&req); err!=nil { util.Error(w,400,"invalid json"); return }; res,err:=c.service.Create(r.Context(), middleware.UserID(r), req); if err!=nil { util.Error(w,400,err.Error()); return }; util.JSON(w,201,res) }
func (c *ReminderController) Delete(w http.ResponseWriter, r *http.Request) { id,ok:=pathID(w,r,"/api/reminders/"); if !ok { return }; if err:=c.service.Delete(r.Context(), middleware.UserID(r), id); err!=nil { util.Error(w,400,err.Error()); return }; util.JSON(w,200,model.MessageResponse{Message:"reminder deleted"}) }
