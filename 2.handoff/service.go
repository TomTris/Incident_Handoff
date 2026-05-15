package main

import (
	"time"
)

func buildHandoffBrief(inc Incident) HandoffBrief {
	actions := 0
	openQuestions := 0
	author := ""
	handoffCount := 0
	for _, entry := range inc.Entries {
		if author != entry.Author {
			author = entry.Author
			handoffCount++
		}
		switch entry.Type {
		case ACTION:
			actions++
		case OPEN_QUESTION:
			openQuestions++
		}
	}

	if handoffCount != 0 {
		handoffCount--
	}

	return HandoffBrief{
		Severity:      inc.Severity,
		Status:        inc.Status,
		Service:       inc.Service,
		ElapsedMinute: int(time.Since(inc.CreatedAt).Minutes()),
		TotalEntry:    len(inc.Entries),
		TakenActions:  actions,
		OpenQuestion:  openQuestions,
		HandoffCount:  handoffCount,
		CreatedAt:     inc.CreatedAt,
	}
}
