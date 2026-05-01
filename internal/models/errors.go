package models

import "errors"

var ErrInvalidID = errors.New("ID is invalid")
var ErrBlankContent = errors.New("Content is empty")
var ErrUserAlreadyExists = errors.New("User already exists")
var ErrUserNotFound = errors.New("User not found")
var ErrInvalidCredentials = errors.New("Invalid credentials")
var ErrInvalidRefreshToken = errors.New("Invalid refresh token")
var ErrInvalidAuthorizationHeader = errors.New("Invalid authorization header")
