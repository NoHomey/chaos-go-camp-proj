package fx

import (
	"context"
	"time"

	"go.uber.org/fx"
)

//IgnoreContext wraps life cycle hook wich does not use context.
func IgnoreContext(lc func() error) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		return lc()
	}
}

//CronJob returns fx.Hook wich acts as a cron job.
func CronJob(job func(), td time.Duration) fx.Hook {
	done := make(chan struct{})
	return fx.Hook{
		OnStart: IgnoreContext(func() error {
			ticker := time.NewTicker(td)
			go func() {
				for {
					select {
					case <-done:
						return
					case <-ticker.C:
						job()
					}
				}
			}()
			return nil
		}),
		OnStop: IgnoreContext(func() error {
			done <- struct{}{}
			return nil
		}),
	}
}
