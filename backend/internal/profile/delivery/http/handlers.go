package http

import (
	"encoding/json"
	"errors"
	"main/backend/internal/models"
	"main/backend/internal/profile"
	"main/backend/internal/session"
	"main/backend/internal/util"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type handlers struct {
	sessionRepository session.Repository
	profileRepository profile.Repository
}

func NewAuthHandlers(sR session.Repository, pR profile.Repository) *handlers {
	return &handlers{sR, pR}
}

func (h *handlers) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found, ok := r.Context().Value("user").(*models.User)

		if !ok {
			util.HandleError(w, errors.New("no user"))
			return
		}

		fields := &models.UserGeneralUpdate{}
		if err := json.NewDecoder(r.Body).Decode(fields); err != nil {
			util.HandleError(w, err)
			return
		}
		fields.Sanitize()
		if err := fields.Validate(); err != nil {
			util.HandleError(w, err)
			return
		}
		fields.UUID = found.UUID

		if fields.Avatar == found.Avatar && fields.DisplayName == found.Avatar && fields.About == found.Avatar {
			w.WriteHeader(http.StatusNotModified)
			return
		}

		if err := h.profileRepository.Update(r.Context(), fields); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) PasswordUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found, ok := r.Context().Value("user").(*models.User)

		if !ok {
			util.HandleError(w, errors.New("no user"))
			return
		}

		fields := &models.UserPasswordUpdate{}
		if err := json.NewDecoder(r.Body).Decode(fields); err != nil {
			util.HandleError(w, err)
			return
		}
		if err := fields.Validate(); err != nil {
			util.HandleError(w, err)
			return
		}

		if err := fields.ComparePasswords(found.Password); err != nil {
			if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				http.Error(w, "wrong credentials", http.StatusBadRequest)
				return
			}
			util.HandleError(w, err)
			return
		}
		if err := fields.HashPassword(); err != nil {
			util.HandleError(w, err)
			return
		}
		fields.UUID = found.UUID

		if err := h.profileRepository.PasswordUpdate(r.Context(), fields); err != nil {
			util.HandleError(w, err)
			return
		}

		session := &models.Session{UserID: found.UUID}
		if err := h.sessionRepository.InvalidateSessions(r.Context(), session); err != nil {
			util.HandleError(w, err)
			return
		}
		sessionId, err := h.sessionRepository.CreateSession(r.Context(), session, 30005)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    sessionId,
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}

		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) PrivacyUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found := r.Context().Value("user").(*models.User)

		fields := &models.UserPrivacyUpdate{}
		if err := json.NewDecoder(r.Body).Decode(fields); err != nil {
			util.HandleError(w, err)
			return
		}
		fields.UUID = found.UUID
		if fields.Privacy == "" || found.Privacy == fields.Privacy {
			w.WriteHeader(http.StatusNotModified)
			return
		}

		if err := h.profileRepository.PrivacyUpdate(r.Context(), fields); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) DeleteAvatar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found, ok := r.Context().Value("user").(*models.User)

		if !ok {
			util.HandleError(w, errors.New("no user"))
			return
		}

		if found.Avatar == "" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		avatar, err := h.profileRepository.DeleteAvatar(r.Context(), found)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.Write([]byte(avatar))
	}
}

func (h *handlers) CreateComment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found := r.Context().Value("user").(*models.User)
		userId := chi.URLParam(r, "user_id")

		uuid, err := uuid.Parse(userId)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		comment := &models.Comment{}
		if err := json.NewDecoder(r.Body).Decode(comment); err != nil {
			util.HandleError(w, err)
			return
		}
		comment.Sanitize()

		if comment.Text == "" {
			http.Error(w, "empty comment", http.StatusBadRequest)
			return
		}

		comment.ProfileID = uuid.String()
		comment.Commentator = found.UUID
		if err := h.profileRepository.CreateComment(r.Context(), comment); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (h *handlers) GetComments() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "user_id")
		uuid, err := uuid.Parse(userId)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		pageLimit := r.URL.Query().Get("pageLimit")
		pageLimitInteger := models.COMMENTS_PAGE_LIMIT
		if pageLimit != "" {
			var err error
			pageLimitInteger, err = strconv.Atoi(pageLimit)
			if err != nil {
				util.HandleError(w, err)
				return
			}
			if pageLimitInteger > models.COMMENTS_PAGE_LIMIT {
				pageLimitInteger = models.COMMENTS_PAGE_LIMIT
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

		comments, err := h.profileRepository.GetComments(r.Context(), &models.User{UUID: uuid}, pageLimitInteger, pageOffsetInteger)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(comments); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) FriendInvite() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found := r.Context().Value("user").(*models.User)
		userId := chi.URLParam(r, "user_id")
		uuid, err := uuid.Parse(userId)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		err = h.profileRepository.FriendInvite(r.Context(), &models.User{UUID: uuid}, found)
		if err != nil {
			if errors.Is(err, models.ErrCannotFriendSelf) ||
				errors.Is(err, models.ErrAlreadyFriends) ||
				errors.Is(err, models.ErrFriendInvitePending) ||
				errors.Is(err, models.ErrAlreadyTriedToFriendRejected) {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				util.HandleError(w, err)
			}
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (h *handlers) HandleFriendInvite(status string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found := r.Context().Value("user").(*models.User)
		userId := chi.URLParam(r, "user_id")
		uuid, err := uuid.Parse(userId)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		err = h.profileRepository.HandleFriendInvite(r.Context(), found, &models.User{UUID: uuid}, status)
		if err != nil {
			if errors.Is(err, models.ErrNoFriendRequest) {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				util.HandleError(w, err)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) DeleteFriend() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found := r.Context().Value("user").(*models.User)
		userId := chi.URLParam(r, "user_id")
		uuid, err := uuid.Parse(userId)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		deleted, err := h.profileRepository.DeleteFriend(r.Context(), found, &models.User{UUID: uuid})
		if err != nil {
			util.HandleError(w, err)
			return
		}

		if !deleted {
			w.WriteHeader(http.StatusNotModified)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) FriendStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found := r.Context().Value("user").(*models.User)
		userId := chi.URLParam(r, "user_id")
		uuid, err := uuid.Parse(userId)
		if err != nil {
			util.HandleError(w, err)
			return
		}
		user := &models.User{UUID: uuid}

		isFriend, err := h.profileRepository.IsFriend(r.Context(), user, found)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		hasIncomingInvite, err := h.profileRepository.HasIncomingInvite(r.Context(), user, found)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		hasOutgoingInvite, err := h.profileRepository.HasOutgoingInvite(r.Context(), user, found)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		status := &models.FriendStatus{IsFriend: isFriend, HasIncomingInvite: hasIncomingInvite, HasOutgoingInvite: hasOutgoingInvite}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(status); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) DeleteComment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found := r.Context().Value("user").(*models.User)
		userId := chi.URLParam(r, "user_id")
		uuid, err := uuid.Parse(userId)
		if err != nil {
			util.HandleError(w, err)
			return
		}
		commentId := chi.URLParam(r, "comment_id")
		commentIdInteger, err := strconv.Atoi(commentId)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		deleted, err := h.profileRepository.DeleteComment(r.Context(), &models.User{UUID: uuid}, found, &models.Comment{ID: commentIdInteger})
		if err != nil {
			util.HandleError(w, err)
			return
		}

		if !deleted {
			http.Error(w, "no comment has been deleted, permissions issue?", http.StatusConflict)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
