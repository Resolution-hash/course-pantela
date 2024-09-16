package messageService

type MessageService struct {
	repo MessageRepository
}

func NewMessageService(repo MessageRepository) *MessageService {
	return &MessageService{repo}
}

func (s *MessageService) CreateMessage(message Message) (Message, error) {
	return s.repo.CreateMessage(message)
}

func (s *MessageService) GetAllMessage() ([]Message, error) {
	return s.repo.GetAllMessage()
}

func (s *MessageService) UpdateMessageByID(id int, message Message) (Message, error) {
	return s.repo.UpdateMessageByID(id, message)
}

func (s *MessageService) DeleteMessageByID(id int) error {
	return s.repo.DeleteMessageByID(id)
}
