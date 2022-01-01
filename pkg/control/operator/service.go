package operator

import (
	"github.com/NexClipper/sudory-prototype-r1/pkg/database"
	"github.com/NexClipper/sudory-prototype-r1/pkg/model"
	"github.com/labstack/echo/v4"
)

type Service struct {
	db *database.DBManipulator

	ID        uint64
	Name      string
	ClusterID uint64
	StepCount uint
	Steps     []*Step

	Response ResponseFn
}

type Step struct {
	ID        uint64
	Name      string
	Sequence  uint64
	Command   string
	Parameter string
}

func NewService(d *database.DBManipulator) Operator {
	return &Service{db: d}
}

func (o *Service) toModelService() *model.Service {
	m := &model.Service{
		Name:      o.Name,
		ClusterID: o.ClusterID,
		StepCount: o.StepCount,
	}

	return m
}

func (o *Service) toModelStep(serviceID uint64) []*model.Step {
	var m []*model.Step
	for _, s := range o.Steps {
		modelStep := &model.Step{
			Name:      s.Name,
			ServiceID: serviceID,
			Sequence:  s.Sequence,
			Command:   s.Command,
			Parameter: s.Parameter,
		}

		m = append(m, modelStep)
	}

	return m
}

func (o *Service) Create(ctx echo.Context) error {
	service := o.toModelService()

	_, err := o.db.CreateService(service)
	if err != nil {
		return err
	}

	steps := o.toModelStep(service.ID)
	_, err = o.db.CreateStep(steps)
	if err != nil {
		return err
	}

	if o.Response != nil {
		o.Response(ctx, nil)
	}

	return nil
}

func (o *Service) Get(ctx echo.Context) error {
	return nil
}
