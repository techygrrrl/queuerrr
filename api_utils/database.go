package api_utils

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type QueueRepository struct {
	db *sqlx.DB
}

func NewDatabaseClient() (*sqlx.DB, error) {
	connectionString := os.Getenv("POSTGRES_CONNECTION_URL")
	if connectionString == "" {
		return nil, fmt.Errorf("missing environment variable: POSTGRES_CONNECTION_URL")
	}
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewQueueRepository(dbClient *sqlx.DB) QueueRepository {
	return QueueRepository{dbClient}
}

type UserPosition struct {
	UserId   string `db:"twitch_user_id"`
	Position int64
}

func (q QueueRepository) GetPosition(userId string) int {
	var positions []UserPosition
	err := q.db.Select(&positions, "SELECT twitch_user_id, row_number() OVER(ORDER BY created_at ASC) AS position FROM queue GROUP BY twitch_user_id")
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	userPosition := -1
	for _, v := range positions {
		if v.UserId == userId {
			userPosition = int(v.Position)
		}
	}

	return int(userPosition)
}

func (q QueueRepository) FindUser(userId string) (*Entry, error) {
	var entries []Entry
	err := q.db.Select(&entries, "SELECT * FROM queue WHERE twitch_user_id = $1 LIMIT 1", userId)
	if err != nil {
		return nil, err
	}

	if len(entries) == 0 {
		return nil, fmt.Errorf("not found")
	}

	return &entries[0], nil
}

func (q QueueRepository) CloseDatabaseConnection() error {
	return q.db.Close()
}

func (q QueueRepository) JoinQueue(userId string, username string, notes string) error {
	entry := Entry{
		Username: username,
		UserId:   userId,
		Notes:    notes,
	}
	_, err := q.db.NamedExec(
		`
			INSERT INTO queue (twitch_user_id, twitch_username, notes)
			VALUES (:twitch_user_id, :twitch_username, :notes);
		`,
		&entry,
	)
	if err != nil {
		return err
	}

	return nil
}
