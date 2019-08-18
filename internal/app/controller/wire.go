package controller

import "github.com/google/wire"

var ControllerSet = wire.NewSet(
	NewRootController,
	wire.Bind(new(RootController), new(*rootController)),
)
