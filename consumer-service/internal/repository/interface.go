package repository

type ConsumerVerifyParams struct {
	FirstName string
	LastName  string
	Phone     string
}

type Consumer interface {
	Verify(params *ConsumerVerifyParams) error
}
