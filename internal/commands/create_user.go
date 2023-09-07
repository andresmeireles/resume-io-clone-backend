package commands

import (
	"fmt"
	"os"
	"regexp"

	"github.com/andresmeireles/resume/internal/db/sql"
	"github.com/andresmeireles/resume/internal/user"
	"github.com/spf13/cobra"
)

func CreateUser() *cobra.Command {
	return &cobra.Command{
		Use:   "createUser",
		Short: "Create a user",
		Long:  `Create a user for the app`,
		Run: func(cmd *cobra.Command, args []string) {
			name := promptCommand(promptContent{
				label:    "Name",
				errorMsg: "Name is required",
			})

			email := promptCommand(promptContent{
				label:    "Email",
				errorMsg: "Email is required",
			})

			if !emailValidation(email) {
				fmt.Println("Email is not valid")
				os.Exit(1)
			}

			password := promptCommand(promptContent{
				label:    "Password",
				errorMsg: "Password is required",
			}, true)

			confirmPassword := promptCommand(promptContent{
				label:    "Confirm Password",
				errorMsg: "Confirm Password is required",
			}, true)

			if password != confirmPassword {
				fmt.Println("Passwords do not match")
				os.Exit(1)
			}

			fmt.Println("Creating user...")
			err := user.Create(sql.GetDB(), name, password, email)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Successfully created user")
			os.Exit(1)
		},
	}
}

func emailValidation(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regexpPattern := regexp.MustCompile(pattern)
	match := regexpPattern.MatchString(email)

	return match
}
