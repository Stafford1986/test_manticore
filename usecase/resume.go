package usecase

import (
	"context"
	"github.com/Stafford1986/test_manticore/pb"
	"log"
)

type ResumeUseCase struct {
	resumeRepo ResumeRepo
	logger     *log.Logger
}

type ResumeRepo interface {
	Insert(ctx context.Context, req *pb.ResumeEntity) error
	Upsert(ctx context.Context, req *pb.ResumeEntity) error
	Find(ctx context.Context, sr *pb.ResumeSearchEntity) (*pb.ResumeSearchResponse, error)
	GetSuggestions(ctx context.Context, req string) ([]string, error)
}

func NewResumeUseCase(repo ResumeRepo) *ResumeUseCase {
	lg := log.Default()
	return &ResumeUseCase{
		resumeRepo: repo,
		logger:     lg,
	}
}

func (uc *ResumeUseCase) Save(ctx context.Context, req *pb.ResumeEntity) error {
	err := uc.resumeRepo.Insert(ctx, req)
	if err != nil {
		uc.logger.Printf("failed to save resume into manticore_native: %v", err)

		return err
	}

	return nil
}

func (uc *ResumeUseCase) Update(ctx context.Context, req *pb.ResumeEntity) error {
	err := uc.resumeRepo.Upsert(ctx, req)
	if err != nil {
		uc.logger.Printf("failed to upsert resume into manticore_native: %v", err)

		return err
	}

	return nil
}

func (uc *ResumeUseCase) Search(ctx context.Context, req *pb.ResumeSearchEntity) (*pb.ResumeSearchResponse, error) {
	res, err := uc.resumeRepo.Find(ctx, req)
	if err != nil {
		uc.logger.Printf("failed find resume: %v", err)

		return nil, err
	}

	return res, nil
}

func (uc *ResumeUseCase) Suggests(ctx context.Context, req string) ([]string, error) {
	res, err := uc.resumeRepo.GetSuggestions(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
