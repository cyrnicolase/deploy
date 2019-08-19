package models

// ErrUserNotExist 用户不存在的错误
type ErrUserNotExist struct{}

// IsErrUserNotExist check is ErrUserNotExist error
func IsErrUserNotExist(err error) bool {
	_, ok := err.(ErrUserNotExist)
	return ok
}

func (ErrUserNotExist) Error() string {
	return "用户不存在"
}

// ErrProjectNotExist 项目不存在
type ErrProjectNotExist struct{}

// IsErrProjectNotExist return bool
func IsErrProjectNotExist(err error) bool {
	_, ok := err.(ErrProjectNotExist)

	return ok
}

func (ErrProjectNotExist) Error() string {
	return "项目不存在"
}
