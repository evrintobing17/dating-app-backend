package premium

import (
	"context"
)

type PremiumUsecase interface {
	UpgradeToPremium(ctx context.Context, UserID int) error
}
