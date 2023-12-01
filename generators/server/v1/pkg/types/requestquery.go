package types

type RequestQuery struct {
	TypeName   string
	MaxResults int
}

type RequestResult struct {
	Items     []*Request
	NextToken string
}
