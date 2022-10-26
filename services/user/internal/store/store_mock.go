package store

import (
	"context"
)

type UserMock struct {
	CreateCredentialsMock func(context.Context, *Credentials) error
	GetCredentialsMock    func(context.Context, string) (*Credentials, error)
}

func (u *UserMock) CreateCredentials(ctx context.Context, pr *Credentials) error {
	if u.CreateCredentialsMock != nil {
		return u.CreateCredentialsMock(ctx, pr)
	}

	return nil
}

func (u *UserMock) GetCredentials(ctx context.Context, pr string) (*Credentials, error) {
	if u.GetCredentialsMock != nil {
		return u.GetCredentialsMock(ctx, pr)
	}

	return nil, nil
}
