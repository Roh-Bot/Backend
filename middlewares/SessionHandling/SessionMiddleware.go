package SessionHandling

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

var sessions = map[string]session{}

type session struct {
	username int
	expiry   time.Time
}

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func Signin(context echo.Context, user_id int) error {
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	sessions[sessionToken] = session{
		username: user_id,
		expiry:   expiresAt,
	}

	context.SetCookie(&http.Cookie{
		Name:    "GMsession_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})
	return fmt.Errorf("hello")
}

func Welcome(context echo.Context) error {
	// We can obtain the session token from the requests cookies, which come with every request
	cookie, err := context.Cookie("GMsession_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return context.NoContent(http.StatusUnauthorized)

		}
		// For any other type of error, return a bad request status
		return context.NoContent(http.StatusBadRequest)

	}
	sessionToken := cookie.Value

	// We then get the session from our session map
	userSession, exists := sessions[sessionToken]
	if !exists {
		// If the session token is not present in session map, return an unauthorized error
		return context.NoContent(http.StatusUnauthorized)

	}
	// If the session is present, but has expired, we can delete the session, and return
	// an unauthorized status
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		return context.NoContent(http.StatusUnauthorized)

	}

	// If the session is valid, return the welcome message to the user
	return context.String(http.StatusOK, fmt.Sprintf("Welcome %s!", userSession.username))
}

func Refresh(context echo.Context) error {
	// (BEGIN) The code from this point is the same as the first part of the `Welcome` route
	cookie, err := context.Cookie("GMsession_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return context.NoContent(http.StatusUnauthorized)

		}
		return context.NoContent(http.StatusBadRequest)

	}
	sessionToken := cookie.Value

	userSession, exists := sessions[sessionToken]
	if !exists {
		return context.NoContent(http.StatusUnauthorized)

	}
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		return context.NoContent(http.StatusUnauthorized)

	}
	// (END) The code until this point is the same as the first part of the `Welcome` route

	// If the previous session is valid, create a new session token for the current user
	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	// Set the token in the session map, along with the user whom it represents
	sessions[newSessionToken] = session{
		username: userSession.username,
		expiry:   expiresAt,
	}

	// Delete the older session token
	delete(sessions, sessionToken)

	// Set the new token as the users `session_token` cookie
	context.SetCookie(&http.Cookie{
		Name:    "GMsession_token",
		Value:   newSessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
	return context.NoContent(201)
}

func Logout(context echo.Context) error {
	cookie, err := context.Cookie("GMsession_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return context.NoContent(http.StatusUnauthorized)

		}
		// For any other type of error, return a bad request status
		return context.NoContent(http.StatusBadRequest)

	}
	sessionToken := cookie.Value

	// remove the users session from the session map
	delete(sessions, sessionToken)

	// We need to let the client know that the cookie is expired
	// In the response, we set the session token to an empty
	// value and set its expiry as the current time
	context.SetCookie(&http.Cookie{
		Name:    "GMsession_token",
		Value:   "",
		Expires: time.Now(),
	})
	return context.NoContent(200)
}
