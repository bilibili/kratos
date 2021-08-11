package ratelimit

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/sra/ratelimit/bbr"
)

type Option func(*options)

// WithLimiter set Limiter implementation,
// default bbr limiter
func WithLimiter(limiter Limiter) Option {
	return func(o *options) {
		o.limiter = limiter
	}
}

// WithErrorCode set error code when limiter triggered,
// default error code 429
func WithErrorCode(code int) Option {
	return func(o *options) {
		o.errCode = code
	}
}

func WithErrorReason(reason string) Option {
	return func(o *options) {
		o.errReason = reason
	}
}

func WithErrorMessage(message string) Option {
	return func(o *options) {
		o.errMessage = message
	}
}

type options struct {
	limiter    Limiter
	errCode    int
	errReason  string
	errMessage string
}

// Limiter limit interface.
type Limiter interface {
	Allow(ctx context.Context) (done func(), err error)
}

// RateLimiter middleware
func RateLimiter(opts ...Option) middleware.Middleware {
	options := &options{
		limiter:    bbr.NewLimiter(),
		errCode:    429,
		errReason:  "rate limit exceeded",
		errMessage: "service unavailable due to rate limit exceeded",
	}
	for _, o := range opts {
		o(options)
	}

	// errLimitExceed is returned when the rate limiter is
	// triggered and the request is rejected due to limit exceeded.
	errLimitExceed := errors.New(options.errCode, options.errReason, options.errMessage)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			done, e := options.limiter.Allow(ctx)
			if e != nil {
				// blocked
				return nil, errLimitExceed
			}
			// passed
			reply, err = handler(ctx, req)
			done()
			return
		}
	}
}
