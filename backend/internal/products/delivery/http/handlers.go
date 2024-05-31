package http

import (
	"encoding/json"
	"errors"
	"main/backend/internal/models"
	"main/backend/internal/products"
	"main/backend/internal/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type handlers struct {
	productsRepository products.Repository
}

func NewAuthHandlers(pR products.Repository) *handlers {
	return &handlers{pR}
}

func (h *handlers) GetTier() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			currencyCode = "USD"
			rows         []*models.TierRow
			err          error
		)

		found, ok := r.Context().Value("user").(*models.User)

		if ok {
			currencyCode = found.CurrencyCode
		}

		genre := r.URL.Query().Get("genre")
		count := r.URL.Query().Get("count")

		if count == "" {
			count = "14"
		}

		if genre == "" {
			rows, err = h.productsRepository.GetTier(r.Context(), currencyCode, count)
		} else {
			rows, err = h.productsRepository.GetTierByGenre(r.Context(), currencyCode, genre, count)
		}

		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(rows)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) GetFeatured() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currencyCode := "USD"
		found, ok := r.Context().Value("user").(*models.User)

		if ok {
			currencyCode = found.CurrencyCode
		}

		rows, err := h.productsRepository.GetFeatured(r.Context(), currencyCode)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(rows)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) GetOwned() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found, ok := r.Context().Value("user").(*models.User)

		if !ok {
			util.HandleError(w, errors.New("no user"))
			return
		}

		ownedJson, err := h.productsRepository.GetOwnedIDs(r.Context(), found)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(ownedJson); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) FindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productId := chi.URLParam(r, "product_id")

		i, err := strconv.Atoi(productId)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		currencyCode := "USD"
		found, ok := r.Context().Value("user").(*models.User)

		if ok {
			currencyCode = found.CurrencyCode
		}

		product, err := h.productsRepository.FindByID(r.Context(), &models.Product{ID: i}, currencyCode, r.URL.Query().Get("lang"))
		if err != nil {
			util.HandleError(w, err)
			return
		}

		if product == nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(product); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) Search() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		term := r.URL.Query().Get("term")
		priceRange := r.URL.Query().Get("priceRange")
		priceRangeSplit := strings.Split(priceRange, ",")
		priceRangeArray := []float32{0, 540000.00}
		if priceRange != "" && len(priceRangeSplit) == 2 {
			min, err := strconv.ParseFloat(priceRangeSplit[0], 32)
			if err != nil {
				util.HandleError(w, err)
				return
			}
			max, err := strconv.ParseFloat(priceRangeSplit[1], 32)
			if err != nil {
				util.HandleError(w, err)
				return
			}
			if min > max {
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode([]*models.SearchProduct{}); err != nil {
					util.HandleError(w, err)
					return
				}

				w.WriteHeader(http.StatusOK)
			}
			priceRangeArray[0] = float32(min)
			priceRangeArray[1] = float32(max)
		}

		specials := r.URL.Query().Get("specials")

		if specials == "1" {
			specials = "0"
		} else {
			specials = "-1"
		}

		genres := r.URL.Query().Get("genres")
		genresArray := []string{}
		if genres != "" {
			genresArray = strings.Split(genres, ",")
		}

		pageLimit := r.URL.Query().Get("pageLimit")
		pageLimitInteger := models.PRODUCTS_PAGE_LIMIT
		if pageLimit != "" {
			var err error
			pageLimitInteger, err = strconv.Atoi(pageLimit)
			if err != nil {
				util.HandleError(w, err)
				return
			}
			if pageLimitInteger > 15 {
				pageLimitInteger = 15
			}
		}

		pageOffset := r.URL.Query().Get("pageOffset")
		pageOffsetInteger := 0
		if pageOffset != "" {
			var err error
			pageOffsetInteger, err = strconv.Atoi(pageOffset)
			if err != nil {
				util.HandleError(w, err)
				return
			}
			if pageOffsetInteger < 0 {
				pageOffsetInteger = 0
			}
		}

		currencyCode := "USD"
		found, ok := r.Context().Value("user").(*models.User)

		if ok {
			currencyCode = found.CurrencyCode
		}

		products, err := h.productsRepository.Search(r.Context(), currencyCode, term, priceRangeArray, specials, genresArray, pageLimitInteger, pageOffsetInteger)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(products); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) Currencies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currencies, err := h.productsRepository.Currencies(r.Context())
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(currencies); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		product := &models.PublishProduct{}
		if err := json.NewDecoder(r.Body).Decode(product); err != nil {
			util.HandleError(w, err)
			return
		}

		if err := h.productsRepository.CreateProduct(r.Context(), product); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
