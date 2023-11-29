package types

type Error struct {
	Code     int
	Message  string
	Metadata map[string]string
}
