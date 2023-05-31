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

// each session contains the username of the user and the time at which it expires
type session struct {
	username string
	expiry   time.Time
}

// we'll use this method later to determine if the session has expired
func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

// Create a struct that models the structure of a user in the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func Signin(context echo.Context) error {
	var creds Credentials
	// Get the JSON body and decode into credentials
	if err := context.Bind(&creds); err != nil {
		fmt.Println("Binding Error")
		return context.String(http.StatusBadRequest, "Invalid Input")

	}

	// Get the expected password from our in memory map
	expectedPassword, ok := users[creds.Username]

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if !ok || expectedPassword != creds.Password {
		return context.NoContent(http.StatusUnauthorized)

	}

	// Create a new random session token
	// we use the "github.com/google/uuid" library to generate UUIDs
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	// Set the token in the session map, along with the session information
	sessions[sessionToken] = session{
		username: creds.Username, //DbChange := userid
		expiry:   expiresAt,
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds

	context.SetCookie(&http.Cookie{
		Name:    "GMsession_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})
	return context.String(200, "Signin Succesful")
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
