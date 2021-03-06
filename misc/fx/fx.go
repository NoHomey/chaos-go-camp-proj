package fx

import "context"

//IgnoreContext wraps life cycle hook wich does not use context.
func IgnoreContext(lc func() error) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		return lc()
	}
}
