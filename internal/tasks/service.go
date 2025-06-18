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

func (s *Service) FindOne(id int) (*Task, *cerror.CError) {
	return s.repo.GetOne(id)
}

func (s *Service) Create(data TaskCreateReq) *Task {
	return s.repo.Create(data.Name, data.Desc)
}

func (s *Service) Cancel(id int) *cerror.CError {
	return s.repo.Cancel(id)
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
