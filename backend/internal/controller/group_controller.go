package controller

import (
    "net/http"
    "strconv"
    "strings"
    "todo-app/backend/internal/middleware"
    "todo-app/backend/internal/model"
    "todo-app/backend/internal/service"
    "todo-app/backend/internal/util"
)

type GroupController struct { service *service.GroupService }
func NewGroupController(s *service.GroupService) *GroupController { return &GroupController{service:s} }
func (c *GroupController) List(w http.ResponseWriter, r *http.Request) { res,err:=c.service.List(r.Context(), middleware.UserID(r)); if err!=nil { util.Error(w,500,err.Error()); return }; util.JSON(w,200,res) }
func (c *GroupController) Create(w http.ResponseWriter, r *http.Request) { var req model.GroupRequest; if err:=util.Decode(r,&req); err!=nil { util.Error(w,400,"invalid json"); return }; res,err:=c.service.Create(r.Context(), middleware.UserID(r), req); if err!=nil { util.Error(w,400,err.Error()); return }; util.JSON(w,201,res) }
func (c *GroupController) Update(w http.ResponseWriter, r *http.Request) { id,ok:=pathID(w,r,"/api/groups/"); if !ok { return }; var req model.GroupRequest; if err:=util.Decode(r,&req); err!=nil { util.Error(w,400,"invalid json"); return }; if err:=c.service.Update(r.Context(), middleware.UserID(r), id, req); err!=nil { util.Error(w,400,err.Error()); return }; util.JSON(w,200,model.MessageResponse{Message:"group updated"}) }
func (c *GroupController) Delete(w http.ResponseWriter, r *http.Request) { id,ok:=pathID(w,r,"/api/groups/"); if !ok { return }; if err:=c.service.Delete(r.Context(), middleware.UserID(r), id); err!=nil { util.Error(w,400,err.Error()); return }; util.JSON(w,200,model.MessageResponse{Message:"group deleted"}) }
func pathID(w http.ResponseWriter, r *http.Request, prefix string) (int64,bool) { raw:=strings.TrimPrefix(r.URL.Path,prefix); raw=strings.Trim(raw,"/"); id,err:=strconv.ParseInt(raw,10,64); if err!=nil || id==0 { util.Error(w,400,"invalid id"); return 0,false }; return id,true }
