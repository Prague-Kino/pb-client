package models

type EnvVars struct {
	BaseURL        string
	AdminEmail     string
	AdminPassword  string
	AppEnvironment string
}

func (e EnvVars) IsDev() bool {
	return e.AppEnvironment == "dev"
}
