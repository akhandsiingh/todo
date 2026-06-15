package controller

import (
    "net/http"
    "todo-app/backend/internal/middleware"
    "todo-app/backend/internal/model"
    "todo-app/backend/internal/service"
    "todo-app/backend/internal/util"
)

type AuthController struct { service *service.AuthService }
func NewAuthController(s *service.AuthService) *AuthController { return &AuthController{service:s} }
func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) { var req model.RegisterRequest; if err:=util.Decode(r,&req); err!=nil { util.Error(w,400,"invalid json"); return }; res,err:=c.service.Register(r.Context(),req); if err!=nil { util.Error(w,400,err.Error()); return }; util.JSON(w,201,res) }
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) { var req model.LoginRequest; if err:=util.Decode(r,&req); err!=nil { util.Error(w,400,"invalid json"); return }; res,err:=c.service.Login(r.Context(),req); if err!=nil { util.Error(w,401,err.Error()); return }; util.JSON(w,200,res) }
func (c *AuthController) Me(w http.ResponseWriter, r *http.Request) { res,err:=c.service.Me(r.Context(), middleware.UserID(r)); if err!=nil { util.Error(w,404,"user not found"); return }; util.JSON(w,200,res) }
