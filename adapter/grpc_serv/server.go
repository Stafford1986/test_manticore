package grpc_serv

import (
	"context"
	"fmt"
	"github.com/Stafford1986/test_manticore/pb"
)

type Server struct {
	resumeUseCase  ResumeUseCase
	vacancyUseCase VacancyUseCase
}

type ResumeUseCase interface {
	Save(ctx context.Context, req *pb.ResumeEntity) error
	Update(ctx context.Context, req *pb.ResumeEntity) error
	Search(ctx context.Context, req *pb.ResumeSearchEntity) ([]uint32, error)
	Suggests(ctx context.Context, req string) ([]string, error)
}

type VacancyUseCase interface {
	Save(ctx context.Context, req *pb.VacancyEntity) error
	SaveBulk(ctx context.Context, req *pb.VacancyList) error
	Update(ctx context.Context, req *pb.VacancyEntity) error
	Search(ctx context.Context, req *pb.VacancySearchEntity) ([]uint32, error)
	Suggests(ctx context.Context, req string) ([]string, error)
}

func New(resumeUc ResumeUseCase, vacancyUc VacancyUseCase) *Server {
	return &Server{
		resumeUseCase:  resumeUc,
		vacancyUseCase: vacancyUc,
	}
}
func (s *Server) VacancySearch(ctx context.Context, req *pb.VacancySearchEntity) (*pb.VacancySearchResponse, error) {
	res, err := s.vacancyUseCase.Search(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.VacancySearchResponse{
		Items: res,
	}, nil
}
func (s *Server) VacancyIndexUpdate(ctx context.Context, req *pb.VacancyEntity) (*pb.CommentedResponse, error) {
	err := s.vacancyUseCase.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.CommentedResponse{Result: true}, nil
}
func (s *Server) VacancyIndexCreate(ctx context.Context, req *pb.VacancyEntity) (*pb.CommentedResponse, error) {
	items := make([]*pb.VacancyEntity, 0, 10000)
	for i := 0; i < 10000; i++ {
		reqC := *req
		reqC.Id += uint32(i)
		reqC.Name = fmt.Sprintf("%s %s %d", req.Name, "some", i)
		items = append(items, &reqC)
	}
	l := pb.VacancyList{
		Items: items,
	}
	err := s.vacancyUseCase.SaveBulk(ctx, &l)
	if err != nil {
		return nil, err
	}
	return &pb.CommentedResponse{Result: true}, nil
}
func (s *Server) ResumeSearch(ctx context.Context, req *pb.ResumeSearchEntity) (*pb.ResumeSearchResponse, error) {
	res, err := s.resumeUseCase.Search(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ResumeSearchResponse{Items: res}, nil
}
func (s *Server) ResumeIndexUpdate(ctx context.Context, req *pb.ResumeEntity) (*pb.CommentedResponse, error) {
	err := s.resumeUseCase.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.CommentedResponse{Result: true}, nil
}
func (s *Server) ResumeIndexCreate(ctx context.Context, req *pb.ResumeEntity) (*pb.CommentedResponse, error) {
	err := s.resumeUseCase.Save(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.CommentedResponse{Result: true}, nil
}
