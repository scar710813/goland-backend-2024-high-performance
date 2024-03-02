package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	usecase "github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/domain/use-case"
)

func ExtractHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil || userId < 1 || userId > 5 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User ID is invalid")
		return
	}

	dbValue := r.Context().Value("DB")
	if dbValue == nil {
		http.Error(w, "Conexão do banco de dados não encontrada no contexto", http.StatusInternalServerError)
		return
	}

	db, ok := dbValue.(*sql.DB)
	if !ok {
		http.Error(w, "Valor no contexto não pode ser convertido para *sql.DB", http.StatusInternalServerError)
		return
	}

	result, err := usecase.NewExtractUseCase(db, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
