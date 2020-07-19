package mock

import (
	"github.com/stretchr/testify/mock"
	"net/http"
)

type RestClientMock struct {
	mock.Mock
}

func (r *RestClientMock) Do(req *http.Request) (*http.Response, error) {
	args := r.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}
