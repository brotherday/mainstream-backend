package db

const (
	version = "Development"
)

func boolPtr(b bool) *bool {
	return &b
}
