package response

const (
	StatusSuccess = "success"
	StatusError   = "error"
)

type Status struct {
	Status  string
	Message string
}
