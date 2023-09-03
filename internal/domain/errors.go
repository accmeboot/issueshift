package domain

type errKey string

var ErrKey = errKey("appError")

type Envelope map[string]any

type ErrNoRecord error
type ErrAlreadyExists error
type ErrEditConflict error
type ErrServer error
type ErrInvalidCredentials error
type ErrBadlyFormattedJson error
