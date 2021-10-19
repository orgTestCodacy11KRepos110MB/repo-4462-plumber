package awssns

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/batchcorp/plumber-schemas/build/go/protos/args"
	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"

	"github.com/batchcorp/plumber/backends/awssns/types"
	types2 "github.com/batchcorp/plumber/types"
)

const BackendName = "AWSSNS"

var ErrMissingTopicARN = errors.New("--topic cannot be empty")

type AWSSNS struct {
	// Base connection options / non-backend-specific options
	connOpts *opts.ConnectionOptions

	// Backend-specific args
	connArgs *args.AWSSNSConn

	Service  types.ISNSAPI
	QueueURL string
	log      *logrus.Entry
}

func New(connOpts *opts.ConnectionOptions) (*AWSSNS, error) {
	if err := validateBaseConnOpts(connOpts); err != nil {
		return nil, errors.Wrap(err, "unable to validate options")
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &AWSSNS{
		connOpts: connOpts,
		connArgs: connOpts.GetAwssns(),
		Service:  sns.New(sess),
		log:      logrus.WithField("backend", "aws-sns"),
	}, nil

}

func (a *AWSSNS) Name() string {
	return BackendName
}

func (a *AWSSNS) Close(_ context.Context) error {
	return nil
}

func (a *AWSSNS) Test(_ context.Context) error {
	return types2.NotImplementedErr
}

func validateBaseConnOpts(connOpts *opts.ConnectionOptions) error {
	if connOpts == nil {
		return errors.New("connection config cannot be nil")
	}

	if connOpts.Conn == nil {
		return errors.New("connection object in connection config cannot be nil")
	}

	if connOpts.GetAwssns() == nil {
		return errors.New("connection config args cannot be nil")
	}

	return nil
}