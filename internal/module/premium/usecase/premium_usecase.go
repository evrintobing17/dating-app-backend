package usecase

import (
	"context"
	"errors"

	"github.com/evrintobing17/dating-app-go/internal/module/auth"
	"github.com/evrintobing17/dating-app-go/internal/module/premium"
)

type PremiumUseCase struct {
	repo     premium.PremiumRepository
	userRepo auth.AuthRepository
}

func NewPremiumUsecase(repo premium.PremiumRepository, userRepo auth.AuthRepository) premium.PremiumUsecase {
	return &PremiumUseCase{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (s *PremiumUseCase) UpgradeToPremium(ctx context.Context, userID int) error {
	count, err := s.repo.GetPremiumByID(ctx, userID)
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("user already premium")
	}
	err = s.repo.Update(ctx, userID)
	go s.userRepo.UpdateUser(ctx, userID)
	return err
}
