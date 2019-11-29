package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

type config struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	retry    int
	sleep    int
}

var c config

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if pg is ready",
	Long:  `Check if pg is ready`,
	Run: func(cmd *cobra.Command, args []string) {
		check(c)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().StringVar(&c.host,
		"host",
		c.host,
		"host")

	checkCmd.Flags().IntVar(&c.port,
		"port",
		5432,
		"port")

	checkCmd.Flags().StringVar(&c.user,
		"user",
		c.user,
		"user")

	checkCmd.Flags().StringVar(&c.password,
		"password",
		c.password,
		"password")

	checkCmd.Flags().StringVar(&c.dbname,
		"dbname",
		c.dbname,
		"dbname")

	checkCmd.Flags().IntVar(&c.retry,
		"retry",
		1,
		"retry")

	checkCmd.Flags().IntVar(&c.sleep,
		"sleep",
		1,
		"sleep")
}

func check(c config) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.host, c.port, c.user, c.password, c.dbname)
	for i := 0; i < c.retry; i++ {
		time.Sleep(time.Duration(c.sleep) * time.Second)
		db, err := sql.Open("postgres", connString)
		if db != nil {
			err := db.Ping()
			defer db.Close()
			if err != nil {
				log.Printf("Error: %s", err.Error())
				continue
			}
		}
		if db == nil {
			log.Printf("Error: %s", err.Error())
			continue
		}
		log.Printf("DB is ready!")
		os.Exit(0)
	}
	log.Printf("DB isn't ready! Retry counter exceeded")
	os.Exit(1)
}
