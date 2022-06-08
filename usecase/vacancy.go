package usecase

import (
	"context"
	"github.com/Stafford1986/test_manticore/pb"
	"log"
)

type VacancyUseCase struct {
	vacancyRepo VacancyRepo
	logger      *log.Logger
}

type VacancyRepo interface {
	Insert(ctx context.Context, req *pb.VacancyEntity) error
	Upsert(ctx context.Context, req *pb.VacancyEntity) error
	Find(ctx context.Context, sr *pb.VacancySearchEntity) (*pb.VacancySearchResponse, error)
	GetSuggestions(ctx context.Context, req string) ([]string, error)
}

func NewVacancyUseCase(repo VacancyRepo) *VacancyUseCase {
	lg := log.Default()
	return &VacancyUseCase{
		vacancyRepo: repo,
		logger:      lg,
	}
}

func (uc *VacancyUseCase) Save(ctx context.Context, req *pb.VacancyEntity) error {
	err := uc.vacancyRepo.Insert(ctx, req)
	if err != nil {
		uc.logger.Printf("failed to save vacancy into manticore_native: %v", err)

		return err
	}

	return nil
}

func (uc *VacancyUseCase) Update(ctx context.Context, req *pb.VacancyEntity) error {
	err := uc.vacancyRepo.Upsert(ctx, req)
	if err != nil {
		uc.logger.Printf("failed to upsert vacancy into manticore_native: %v", err)

		return err
	}

	return nil
}

func (uc *VacancyUseCase) Search(ctx context.Context, req *pb.VacancySearchEntity) (*pb.VacancySearchResponse, error) {
	res, err := uc.vacancyRepo.Find(ctx, req)
	if err != nil {
		uc.logger.Printf("failed find resume: %v", err)

		return nil, err
	}

	return res, nil
}

func (uc *VacancyUseCase) Suggests(ctx context.Context, req string) ([]string, error) {
	res, err := uc.vacancyRepo.GetSuggestions(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}
