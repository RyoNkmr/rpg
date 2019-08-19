package usecase

import "github.com/google/wire"

var UsecaseSet = wire.NewSet(
	NewBattleUsecase,
	NewSystemUsecase,
	wire.Bind(new(BattleUsecase), new(*battleUsecase)),
	wire.Bind(new(SystemUsecase), new(*systemUsecase)),
)
