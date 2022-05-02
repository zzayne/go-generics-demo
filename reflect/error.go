package reflect

import "errors"

var (
	ErrShouldBePrt       = errors.New("data should be ptr")
	ErrFieldNotExist     = errors.New("filed not exist")
	ErrValueTypeNotMatch = errors.New("value type not match")
)
