package server

type bodyAssignedError struct{}

func (e bodyAssignedError) Error() string {
	return "body already assigned"
}

func (e bodyAssignedError) HasBodyError() {}
