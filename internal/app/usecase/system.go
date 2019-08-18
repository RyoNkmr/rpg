package usecase

import (
	"time"

	"github.com/RyoNkmr/rpg/internal/app/usecase/output"
)

type systemUsecase struct {
	root   output.RootPresenter
	status output.StatusPresenter
	main   output.MainPresenter
}

type SystemUsecase interface {
	Run() error
	Hr()
	AddLine(string)
	AddLines(systems []string, delay time.Duration)
	AddLineBetweenHr(string)
}

func NewSystemUsecase(root output.RootPresenter, status output.StatusPresenter, main output.MainPresenter) *systemUsecase {
	return &systemUsecase{root, status, main}
}

func (u *systemUsecase) Run() error {
	go u.status.Update()
	return u.root.Run()
}

func (u *systemUsecase) Hr() {
	u.main.Hr()
}

func (u *systemUsecase) AddLine(m string) {
	u.main.AddLine(m)
}

func (u *systemUsecase) AddLines(systems []string, delay time.Duration) {
	u.main.AddLines(systems, delay)
}

func (u *systemUsecase) AddLineBetweenHr(m string) {
	u.main.Hr()
	u.main.AddLine(m)
	u.main.Hr()
}
