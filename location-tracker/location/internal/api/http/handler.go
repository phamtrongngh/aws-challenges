package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"location/internal/dataaccess/broker"
	"location/internal/dataaccess/repo"
	"location/internal/model"
)

type handler struct {
	lp   broker.LocationProducer
	repo repo.LocationRepo
}

func NewHandler(lp broker.LocationProducer, r repo.LocationRepo) *handler {
	return &handler{lp, r}
}

// UploadLocation godoc
// @Summary Upload location
// @Description Upload location file
// @Tags locations
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "Location file"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /v1/location/upload [post]
func (s *handler) UploadLocation(w http.ResponseWriter, r *http.Request) {
	// load the uploaded file
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var location model.LocationUpdate

	// Parse request body
	if err := json.NewDecoder(file).Decode(&location); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Publish message
	if err := s.lp.Produce(r.Context(), &location); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateLocation godoc
// @Summary Update location
// @Description Update location
// @Tags locations
// @Accept  json
// @Produce  json
// @Param location body model.LocationUpdate true "Location update"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /v1/location [post]
func (s *handler) UpdateLocation(w http.ResponseWriter, r *http.Request) {
	var location model.LocationUpdate

	// Parse request body
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Publish message
	if err := s.lp.Produce(r.Context(), &location); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetLatestLocation godoc
// @Summary Get latest locations
// @Description Get the latest location for a device
// @Tags locations
// @Accept  json
// @Produce  json
// @Param device_id path string true "Device ID"
// @Param limit query int false "Limit"
// @Success 200 {array} model.LocationUpdateValue
// @Failure 500
// @Router /v1/location/{device_id} [get]
func (s *handler) GetLatestLocation(w http.ResponseWriter, r *http.Request) {
	deviceID := r.PathValue("device_id")
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if err != nil {
		limit = 1
	}

	locValues, err := s.repo.GetLatestLocation(r.Context(), deviceID, model.LocationFilter{
		Limit: &limit,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locValues)
	// w.WriteHeader(http.StatusOK)
}
