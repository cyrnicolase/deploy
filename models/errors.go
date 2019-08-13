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
