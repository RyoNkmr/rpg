package usecase

import "github.com/google/wire"

var UsecaseSet = wire.NewSet(
	NewAttackUsecase,
	NewSystemUsecase,
	wire.Bind(new(AttackUsecase), new(*attackUsecase)),
	wire.Bind(new(SystemUsecase), new(*systemUsecase)),
)
