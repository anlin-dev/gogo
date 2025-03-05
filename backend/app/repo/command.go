package repo

type CommandRepo struct{}

type ICommandRepo interface {
}

func NewCommandRepo() ICommandRepo {
	return &CommandRepo{}
}
