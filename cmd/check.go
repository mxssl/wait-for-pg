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
	sslmode  string
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

	checkCmd.Flags().StringVar(&c.sslmode,
		"sslmode",
		"require",
		"sslmode")
}

func buildConnString(c config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s connect_timeout=1",
		c.host, c.port, c.user, c.password, c.dbname, c.sslmode)
}

func tryConnect(connString string) error {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer func() {
		if closeErr := db.Close(); closeErr != nil {
			log.Printf("Error closing database connection: %s", closeErr.Error())
		}
	}()
	if err := db.Ping(); err != nil {
		return fmt.Errorf("ping: %w", err)
	}
	return nil
}

func checkWithRetry(c config) error {
	connString := buildConnString(c)
	var lastErr error
	for i := 0; i < c.retry; i++ {
		time.Sleep(time.Duration(c.sleep) * time.Second)
		if err := tryConnect(connString); err != nil {
			log.Printf("Error: %s", err.Error())
			lastErr = err
			continue
		}
		log.Printf("DB is ready!")
		return nil
	}
	return fmt.Errorf("DB isn't ready! Retry counter exceeded: %w", lastErr)
}

func check(c config) {
	if err := checkWithRetry(c); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
