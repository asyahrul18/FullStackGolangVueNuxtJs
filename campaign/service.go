package campaign

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		Campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return Campaigns, err
		}

		return Campaigns, nil
	}

	Campaigns, err := s.repository.FindAll()
	if err != nil {
		return Campaigns, err
	}

	return Campaigns, nil
}
