package premium

import "context"

type PremiumRepository interface {
	Update(ctx context.Context, userID int) error
	GetPremiumByID(ctx context.Context, userID int) (int, error)
}
