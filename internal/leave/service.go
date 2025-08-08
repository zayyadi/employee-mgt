package leave

import (
	"employee-management/internal/models"
	"errors"

	"github.com/google/uuid"
)

// Service handles leave-related operations
type Service struct {
	repo Repository
}

// NewService creates a new leave service
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// Leave Type services
func (s *Service) CreateLeaveType(leaveTypeData *models.LeaveTypeCreate) (*models.LeaveType, error) {
	return s.repo.CreateLeaveType(leaveTypeData)
}

func (s *Service) GetLeaveTypeByID(id uuid.UUID) (*models.LeaveType, error) {
	return s.repo.GetLeaveTypeByID(id)
}

func (s *Service) ListLeaveTypes() ([]models.LeaveType, error) {
	return s.repo.ListLeaveTypes()
}

func (s *Service) UpdateLeaveType(id uuid.UUID, leaveTypeData *models.LeaveTypeUpdate) (*models.LeaveType, error) {
	return s.repo.UpdateLeaveType(id, leaveTypeData)
}

func (s *Service) DeleteLeaveType(id uuid.UUID) error {
	return s.repo.DeleteLeaveType(id)
}

// Leave Request services
func (s *Service) CreateLeaveRequest(leaveRequestData *models.LeaveRequestCreate) (*models.LeaveRequest, error) {
	// Business logic: check if leave type exists, if employee has enough balance etc.
	// For now, we'll keep it simple.
	return s.repo.CreateLeaveRequest(leaveRequestData)
}

func (s *Service) GetLeaveRequestByID(id uuid.UUID) (*models.LeaveRequest, error) {
	return s.repo.GetLeaveRequestByID(id)
}

func (s *Service) ListLeaveRequests() ([]models.LeaveRequest, error) {
	return s.repo.ListLeaveRequests()
}

func (s *Service) ApproveLeaveRequest(id uuid.UUID, approvedByUserID uuid.UUID) (*models.LeaveRequest, error) {
	request, err := s.repo.GetLeaveRequestByID(id)
	if err != nil {
		return nil, err
	}
	if request.Status != "pending" {
		return nil, errors.New("leave request is not in pending state")
	}
	return s.repo.UpdateLeaveRequestStatus(id, "approved", &approvedByUserID)
}

func (s *Service) RejectLeaveRequest(id uuid.UUID, rejectedByUserID uuid.UUID) (*models.LeaveRequest, error) {
	request, err := s.repo.GetLeaveRequestByID(id)
	if err != nil {
		return nil, err
	}
	if request.Status != "pending" {
		return nil, errors.New("leave request is not in pending state")
	}
	return s.repo.UpdateLeaveRequestStatus(id, "rejected", &rejectedByUserID)
}
