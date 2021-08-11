package taskName

type Task interface {
	Run()
	BindParameters(map[string]string)
	GetName() string
}
