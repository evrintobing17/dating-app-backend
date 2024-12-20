package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/evrintobing17/dating-app-go/internal/module/swipe"
	"github.com/evrintobing17/dating-app-go/internal/repository"
)

type SwipeUseCase struct {
	repo  swipe.SwipeRepository
	cache *repository.RedisClient
}

func NewSwipeUsecase(repo swipe.SwipeRepository, cache *repository.RedisClient) swipe.SwipeUsecase {
	return &SwipeUseCase{repo: repo, cache: cache}
}

func (s *SwipeUseCase) Swipe(ctx context.Context, userID, profileID int, action string, isPremium bool) error {
	if action != "like" && action != "pass" {
		return errors.New("invalid action")
	}

	// Redis key for daily swipes
	redisKey := fmt.Sprintf("user:%d:swipes:%s", userID, time.Now().Format("2006-01-02"))

	// Enforce daily swipe limit for non-premium users
	if !isPremium {
		swipeCount := s.cache.Client.SCard(ctx, redisKey).Val()
		fmt.Println(swipeCount, "apa")
		if swipeCount >= 10 {
			return errors.New("daily swipe limit reached")
		}
	}

	// Check if profile was already swiped today
	alreadySwiped, err := s.cache.IsMemberOfSet(redisKey, profileID)
	if err != nil {
		return err
	}

	if alreadySwiped {
		return errors.New("profile already swiped today")
	}

	err = s.repo.CreateSwipe(userID, profileID, action, time.Now())
	if err != nil {
		return err
	}

	// Add profile to Redis set
	if err := s.cache.AddToSet(redisKey, profileID, 24*time.Hour); err != nil {
		return err
	}

	// Record the swipe in the database for persistence
	return err
}
