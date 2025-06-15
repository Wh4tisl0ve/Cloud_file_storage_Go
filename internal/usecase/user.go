package user

type (
	Login interface {
		Execute(username string, password string) (token string, err error)
	}

	Logout interface {
		Execute(id int) error
	}

	Register interface {
		Execute(username string, password string) error
	}
)
