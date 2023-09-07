package commands

import (
	"fmt"
	"os"

	"github.com/andresmeireles/resume/internal/db/sql"
	"github.com/spf13/cobra"
)

func ExecuteMigrations() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Running Up migrations",
		Long:  `Execute Up migrations`,
		Run: func(cmd *cobra.Command, args []string) {
			sql.Up()
			fmt.Println("Migration executed")
			os.Exit(1)
		},
	}
}

func RollbackMigration() *cobra.Command {
	return &cobra.Command{
		Use:   "rollback",
		Short: "Rollback migrations",
		Long:  `Rollback executed migrations`,
		Run: func(cmd *cobra.Command, args []string) {
			sql.Down()
			fmt.Println("ROllback executed")
			os.Exit(1)
		},
	}
}
