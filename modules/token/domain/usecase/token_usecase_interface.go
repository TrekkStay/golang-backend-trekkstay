package usecase

type TokenProvider interface {
	Generate(payload map[string]interface{}, expiry int) (map[string]interface{}, error)
}
