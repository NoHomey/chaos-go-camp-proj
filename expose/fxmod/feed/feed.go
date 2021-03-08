package feed

import (
	routes "github.com/NoHomey/chaos-go-camp-proj/expose/routes/feed"
	service "github.com/NoHomey/chaos-go-camp-proj/service/feed"
	"go.uber.org/fx"
)

//Module bundles fx.Options for the feed Fx Module.
var Module = fx.Options(
	fx.Provide(service.Use),
	fx.Invoke(routes.Register),
)
