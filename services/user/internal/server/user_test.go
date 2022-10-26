package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"pad/services/user/internal/store"
	"pad/services/user/services/user"
)

func Test_Register(t *testing.T) { // REQ: Unit tests
	var (
		dbMock = &store.UserMock{}
		srv    = NewUserServer(dbMock)
		ctx    = context.Background()
	)

	tests := []struct {
		name        string
		mock        func()
		req         *user.RegisterRequest
		expectedErr error
	}{
		{
			name: "Successful",
			mock: func() {
				dbMock.CreateCredentialsMock = func(ctx context.Context, credentials *store.Credentials) error {
					return nil
				}
			},
			req: &user.RegisterRequest{
				Name:     "test_name",
				Password: "test_pass",
			},
			expectedErr: nil,
		},
		{
			name: "Unsuccessful",
			mock: func() {
				dbMock.CreateCredentialsMock = func(ctx context.Context, credentials *store.Credentials) error {
					return fmt.Errorf("something went wrong")
				}
			},
			req: &user.RegisterRequest{
				Name:     "test_name",
				Password: "test_pass",
			},
			expectedErr: fmt.Errorf("something went wrong"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()

			_, actualErr := srv.Register(ctx, tc.req)
			assert.Equal(t, tc.expectedErr, actualErr)
		})
	}
}
