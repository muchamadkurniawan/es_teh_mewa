package service

import "context"

type DashboardService interface {
	Index(ctx context.Context)
}
