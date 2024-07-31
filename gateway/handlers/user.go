package handlers

import (
	"log"
	"net/http"

	common "github.com/tuananh9201/go-eco/common"
	pb "github.com/tuananh9201/go-eco/common/api"
)

type userHandler struct {
	client pb.UserServiceClient
}

func NewUserHandler(client pb.UserServiceClient) *userHandler {
	return &userHandler{client: client}
}
func (h *userHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/users", h.CreateNewUserHandler)
	mux.HandleFunc("GET /api/users", h.GetUsersHandler)
}

func (h *userHandler) CreateNewUserHandler(w http.ResponseWriter, r *http.Request) {
	var user pb.CreateUserRequest
	if err := common.ReadJSON(r, &user); err != nil {
		common.WriteError(w, http.StatusBadRequest, "gateway: ReadJSON: "+err.Error())
		return
	}
	log.Println("gateway: CreateNewUserHandler")
	o, err := h.client.CreateUser(r.Context(), &user)
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, "gateway: CreateOrder: "+err.Error())
		return
	}
	common.WriteJSON(w, http.StatusCreated, o)
}

func (h *userHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	userRequest := &pb.GetListUserRequest{}
	log.Println("gateway: GetUsersHandler")
	o, err := h.client.GetListUser(r.Context(), userRequest)
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, "gateway: GetUsersHandler: "+err.Error())
		return
	}
	common.WriteJSON(w, http.StatusOK, o)
}
