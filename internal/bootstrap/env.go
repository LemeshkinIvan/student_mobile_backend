package bootstrap

type AppEnvironment struct{}

func NewAppEnvironment() *AppEnvironment {
	return &AppEnvironment{}
}
