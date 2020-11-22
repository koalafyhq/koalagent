package main

import (
	worker "github.com/contribsys/faktory_worker_go"

	"github.com/koalafyhq/koalagent/internal/acme"
)

func main() {
	mgr := worker.NewManager()

	mgr.Concurrency = 5

	mgr.Register("issue_new_certificate", acme.IssueNewCertificate)

	mgr.Run()
}
