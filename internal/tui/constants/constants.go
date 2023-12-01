package constants

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/gitu/jirio/internal/tui/style"
	"strings"
	"time"
)

var LogoString = strings.Join([]string{
	"░▀▀█░▀█▀░█▀▄░▀█▀░█▀█",
	"░░░█░░█░░█▀▄░░█░░█░█",
	"░▀▀░░▀▀▀░▀░▀░▀▀▀░▀▀▀",
}, "\n")

const ToastDuration = time.Second * 5

const TableSeparator = "|【=◈︿◈=】|"

const TablePadding = "   "

var JobsTableStatusStyles = map[string]lipgloss.Style{
	TablePadding + "pending" + TablePadding: style.JobRowPending,
	TablePadding + "dead" + TablePadding:    style.JobRowDead,
}

var TasksTableStatusStyles = JobsTableStatusStyles

const SaveDialogPlaceholder = "Output file name (path optional)"
