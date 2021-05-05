package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api/internal/comment"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

type Response struct {
	Message string
}

func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I'm alive"}); err != nil {
			panic(err)
		}
	})

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to Parse Uint from ID")
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		fmt.Fprintf(w, "Error Returning Comment by ID")
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}

}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(w, "Failed To Return All Comments")
	}
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Fprintf(w, "Failed to decode json body")
	}
	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse uint from ID")
	}

	comment, err = h.Service.UpdateComment(uint(commentID), comment)

	if err != nil {
		fmt.Fprintf(w, "failed to update comment")
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Fprintf(w, "Failed to decode json")
	}

	comment, err := h.Service.PostComment(comment)

	if err != nil {
		fmt.Fprintf(w, "Failed to post new comment")
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse uint from ID")
	}

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		fmt.Fprintf(w, "failed to delete comment by ID")
	}

	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(Response{Message: "I'm alive"}); err != nil {
		panic(err)
	}
}
