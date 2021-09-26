package response

type ResponseSource string

const (
	SourceAPIServer     ResponseSource = "API Server"
	SourceBackendServer ResponseSource = "Backend Server"
)
