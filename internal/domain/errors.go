package domain

import "errors"

var ErrNoRecord = errors.New("the requested resource could not be found")
var ErrEditConflict = errors.New("edit conflict")
var ErrServer = errors.New("the server encountered a problem and could not process your request")
