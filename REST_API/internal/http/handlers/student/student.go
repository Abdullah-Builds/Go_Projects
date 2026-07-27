package student

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/Abdullah-Builds/Go_Projects/internal/storage"
	"github.com/Abdullah-Builds/Go_Projects/internal/types"
	"github.com/Abdullah-Builds/Go_Projects/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.CustomResponse(err))
			return
		}
		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.CustomResponse(err))
			return
		}

		//validate the body provided is correct
		validate := validator.New()
		v_err := validate.Struct(student)

		if v_err != nil {
			validate_error := v_err.(validator.ValidationErrors)
			response.WriteJSON(w, http.StatusBadRequest, response.ValidatorResponse(validate_error))
			return
		}

		LastId, err := storage.CreateStudent(
			student.Name,
			student.Age,
			student.Email,
		)

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.CustomResponse(err))
			return
		}
		response.WriteJSON(w, http.StatusAccepted, map[string]int64{"id": LastId})
	}
}
func GetByID(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid student ID", http.StatusBadRequest)
			return
		}

		student, err := storage.GetStudentByID(id)
		if err != nil {
			http.Error(w, "Student not found", http.StatusNotFound)
			return
		}

		response.WriteJSON(w, http.StatusOK, student)
	}
}

func GetAllStudents(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit := 10
		cursor := 0

		if l := r.URL.Query().Get("limit"); l != "" {
			if value, err := strconv.Atoi(l); err == nil && value > 0 {
				limit = value
			}
		}

		if c := r.URL.Query().Get("cursor"); c != "" {
			if value, err := strconv.Atoi(c); err == nil {
				cursor = value
			}
		}

		students, nextCursor, err := storage.GetStudents(limit, cursor)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.WriteJSON(w, http.StatusOK, map[string]any{
			"students":    students,
			"next_cursor": nextCursor,
		})
	}
}
