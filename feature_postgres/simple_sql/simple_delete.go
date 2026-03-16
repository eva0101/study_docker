package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DelteRow(ctx context.Context, conn *pgx.Conn, tasksID []int) error {
	sqlQuery := `
	DELETE FROM tasks
	WHERE id = ANY($1);
	`

	_, err := conn.Exec(ctx, sqlQuery, tasksID)

	return err
}
