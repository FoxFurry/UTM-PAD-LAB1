/*
Package service

Connects cluster with web interface.
Initializes cluster, handles mapping of http models to cluster models, wraps errors
*/
package service

import (
	"context"
	"net/http"
	"strings"

	"pad/services/cache/internal/api/httperr"
	"pad/services/cache/internal/api/model"
	"pad/services/cache/internal/cluster"
	"pad/services/cache/internal/command"
)

// Service represents layer between web and cluster interfaces. It does only one thing: pass http model to cluster
type Service interface {
	Execute([]model.Command) ([]string, error)
}

type service struct {
	data cluster.Cluster
}

// New creates AND initializes a new service
// TODO: Unify 'New' methods behavior. Some of them do init, some don't
func New(ctx context.Context) Service {
	newCluster := cluster.New()
	newCluster.Initialize(ctx)

	return &service{
		data: newCluster,
	}
}

// Execute transforms http transaction into cluster transaction and passes it to cluster itself
// Wraps errors into http errors
//
// TODO: Rework mapping between string value of command and actual command. I don't like this explicit SWITCH
// TODO: Every command should have unified append behavior. Currently GET and SET have different appends
func (s *service) Execute(httpTrns []model.Command) ([]string, error) {
	var clusterTrns []command.Command

	for _, httpCommand := range httpTrns {
		switch strings.ToUpper(httpCommand.CmdType) {
		case "GET":
			clusterTrns = append(clusterTrns, command.Get(httpCommand.Key))
		case "SET":
			clusterTrns = append(clusterTrns, command.Set(httpCommand.Key, httpCommand.Value))
		default:
			return nil, httperr.New("unknown command", http.StatusBadRequest)
		}
	}

	result, err := s.data.Execute(clusterTrns) // Execute cluster transaction
	if err != nil {
		return nil, httperr.WrapHttp(err, "could not execute command", http.StatusInternalServerError)
	}

	return result, nil
}
