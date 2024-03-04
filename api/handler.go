package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/luquxSentinel/parkingspace/service"
	"github.com/luquxSentinel/parkingspace/types"
)

type Handler struct {
	authService    service.AuthenticationService
	parkingService service.ParkingService
	bookingService service.BookingService
	paymentService service.PaymentService
}

func NewHandler(authService service.AuthenticationService,
	parkingService service.ParkingService,
	bookingService service.BookingService,
	paymentService service.PaymentService) *Handler {

	return &Handler{
		authService:    authService,
		parkingService: parkingService,
		bookingService: bookingService,
		paymentService: paymentService,
	}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {

	// request data
	reqData := new(types.SignUpData)

	// read request data
	err := h.read(r.Body, reqData)
	if err != nil {
		http.Error(w, "invalid request data", http.StatusBadRequest)
		return
	}

	// call authentication service service
	err = h.authService.SignUp(r.Context(), reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.write(w, "successfully created user", http.StatusOK)
	if err != nil {
		http.Error(w, "response error", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	// request data
	reqData := new(types.SignInData)

	// decode request data
	err := h.read(r.Body, reqData)
	if err != nil {
		http.Error(w, "invalid request data", http.StatusBadRequest)
		return
	}

	// call authentication service
	user, token, err := h.authService.SignIn(r.Context(), reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("authorization", *token)

	// write response
	err = h.write(w, user.Response(), http.StatusOK)
	if err != nil {
		http.Error(w, "response error", http.StatusInternalServerError)
	}
}

func (h *Handler) NearBy(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) Exit(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) Enter(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) Renew(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) Pay(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) read(r io.Reader, v any) error {
	// decode json data
	return json.NewDecoder(r).Decode(v)
}

func (h *Handler) write(w http.ResponseWriter, v any, statusCode int) error {
	//	encode type to json && write to io.Writer
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(map[string]any{"data": v})
}
