package messageService

import (
	"gorm.io/gorm"
)

type MessageRepository interface {
	CreateMessage(message Message) (Message, error)
	GetAllMessage() ([]Message, error)
	UpdateMessageByID(id int, message Message) (Message, error)
	DeleteMessageByID(id int) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db}
}

func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

func (r *messageRepository) GetAllMessage() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	if err != nil {
		return []Message{}, err
	}
	return messages, nil
}

func (r *messageRepository) UpdateMessageByID(id int, message Message) (Message, error) {
	result := r.db.Model(&Message{}).Where("id = ?", id).Update("text", message.Text)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

func (r *messageRepository) DeleteMessageByID(id int) error {
	result := r.db.Delete(&Message{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
