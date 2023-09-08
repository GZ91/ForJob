package errorsapp

import "errors"

var ErrLinkAlreadyExists = errors.New("the link already exists")
var ErrAlredyBeenRegistered = errors.New("this service has already been registered")
var ErrInvalidLinkReceived = errors.New("an invalid link was received")
var ErrLineURLDeleted = errors.New("this entry has been deleted")
var ErrNotConfiguration = errors.New("not configuration")
var ErrNotFoundToken = errors.New("not found token")
var ErrNotFoundLink = errors.New("not found link")
