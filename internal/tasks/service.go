package tasks

import (
	"workmate/pkg/cerror"
	"workmate/pkg/utils"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) FindOne(name string) (*Task, *cerror.CError) {
	return s.repo.GetOne(name)
}

func (s *Service) Create(data TaskCreateReq) (*Task, *cerror.CError) {
	return s.repo.Create(data.Name, data.Desc)
}

func (s *Service) Cancel(name string) *cerror.CError {
	return s.repo.Cancel(name)
}

func (s *Service) Listing() TaskListing {

	data := s.repo.FindMany()

	meta := utils.Reduce(data, func(el *Task, ac *TaskMeta) *TaskMeta {
		switch el.Status {
		case TaskFailedStatus:
			ac.Failed++
		case TaskSuccessStatus:
			ac.Success++
		case TaskInProcessStatus:
			ac.Process++
		}
		return ac
	}, &TaskMeta{})

	return TaskListing{
		Data: data,
		Meta: meta,
	}

}
