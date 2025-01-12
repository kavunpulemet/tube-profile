package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"tube-profile/internal/model"
	profileservice "tube-profile/internal/service"
	"tube-profile/internal/utils"
)

func CreateProfile(ctx utils.MyContext, service profileservice.ProfileService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("UserID").(string)

		log.Println("CreateProfile userID", userID)

		var profile model.Profile
		if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
			utils.NewErrorResponse(ctx, w, "invalid JSON payload", http.StatusBadRequest)
			return
		}

		profile.UserID = userID

		err := service.CreateProfile(ctx, profile)
		if err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := utils.WriteResponse(w, http.StatusOK, utils.StatusResponse{Status: "ok"}); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetProfile(ctx utils.MyContext, service profileservice.ProfileService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("UserID").(string)

		profile, err := service.GetProfile(ctx, userID)
		if err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = utils.WriteResponse(w, http.StatusOK, profile); err != nil {
			utils.NewErrorResponse(ctx, w, "internal server error", http.StatusInternalServerError)
			return
		}
	}
}

func UpdateProfile(ctx utils.MyContext, service profileservice.ProfileService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("UserID").(string)

		var profile model.Profile
		if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
			utils.NewErrorResponse(ctx, w, "invalid JSON payload", http.StatusBadRequest)
			return
		}

		profile.UserID = userID

		if err := service.UpdateProfile(ctx, profile); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := utils.WriteResponse(w, http.StatusOK, utils.StatusResponse{Status: "ok"}); err != nil {
			utils.NewErrorResponse(ctx, w, "internal server error", http.StatusInternalServerError)
			return
		}
	}
}
