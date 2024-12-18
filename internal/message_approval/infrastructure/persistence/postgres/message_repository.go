package postgres

import (
	"database/sql"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/miladvatankhah/maker-checker/internal/message_approval/domain/entities"
)

type MessageRepositoryImpl struct {
	db *sql.DB
}

func NewMessageRepositoryImpl(db *sql.DB) *MessageRepositoryImpl {
	return &MessageRepositoryImpl{db: db}
}

func (repo *MessageRepositoryImpl) Save(message *entities.Message) error {
	_, err := repo.db.Exec(`
        INSERT INTO messages (id, content, status, sender_id, receiver_id)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (id) DO UPDATE
        SET content = EXCLUDED.content, status = EXCLUDED.status, sender_id = EXCLUDED.sender_id, receiver_id = EXCLUDED.receiver_id
    `, message.ID, message.Content.Text, message.Status, message.SenderID, message.ReceiverID)
	return err
}

func (repo *MessageRepositoryImpl) FindByID(id string) (*entities.Message, error) {
	var message entities.Message
	err := repo.db.QueryRow(`
        SELECT id, content, status, sender_id, receiver_id
        FROM messages
        WHERE id = $1
    `, id).Scan(&message.ID, &message.Content.Text, &message.Status, &message.SenderID, &message.ReceiverID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &message, nil
}
