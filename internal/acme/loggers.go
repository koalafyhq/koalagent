package acme

import "log"

// Error is
func Error(ns string, err error) {
	log.Fatalf("[ERROR] [%s] %s", ns, err.Error())
}

// Debug is
func Debug(ns string, msg string) {
	log.Printf("[DEBUG] [%s] %s \n", ns, msg)
}

// Info is
func Info(ns string, msg string) {
	log.Printf("[INFO] [%s] %s \n", ns, msg)
}
