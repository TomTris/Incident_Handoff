package main

import (
	"context"
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type InstrumentedIncidentStore struct {
	s IncidentStore
}

func (s *InstrumentedIncidentStore) MetricInit() {
	incidents, err := s.s.ListIncidents(context.Background(), IncidentFilter{})
	if err != nil {
		log.Fatal(err)
	}
	for _, incident := range incidents {
		incidentTotal.WithLabelValues(incident.Status).Inc()
		totalEntries.Add(float64(len(incident.Entries)))
	}
}

func (s *InstrumentedIncidentStore) CreateIncident(ctx context.Context, req CreateIncidentRequest) (Incident, error) {
	timer := prometheus.NewTimer(dbQueryDurationSeconds.WithLabelValues("create_incident"))
	inc, err := s.s.CreateIncident(ctx, req)
	timer.ObserveDuration()
	if err == nil {
		incidentTotal.WithLabelValues(inc.Status).Inc()
	}
	return inc, err
}

func (s *InstrumentedIncidentStore) GetIncident(ctx context.Context, id string) (Incident, error) {
	timer := prometheus.NewTimer(dbQueryDurationSeconds.WithLabelValues("get_incident"))
	defer timer.ObserveDuration()
	return s.s.GetIncident(ctx, id)
}

func (s *InstrumentedIncidentStore) AddEntry(ctx context.Context, incidentID string, entry TimelineEntry) (TimelineEntry, error) {
	timer := prometheus.NewTimer(dbQueryDurationSeconds.WithLabelValues("add_entry"))
	entry, err := s.s.AddEntry(ctx, incidentID, entry)
	timer.ObserveDuration()

	if err == nil {
		totalEntries.Inc()
	}
	return entry, err
}

func (s *InstrumentedIncidentStore) ListIncidents(ctx context.Context, filter IncidentFilter) ([]Incident, error) {
	timer := prometheus.NewTimer(dbQueryDurationSeconds.WithLabelValues("list_incident"))
	defer timer.ObserveDuration()
	return s.s.ListIncidents(ctx, filter)
}

func (s *InstrumentedIncidentStore) UpdateIncident(ctx context.Context, id string, update IncidentUpdate) (Incident, error) {
	timer := prometheus.NewTimer(dbQueryDurationSeconds.WithLabelValues("update_incident"))
	incBefore, err := s.s.UpdateIncident(ctx, id, update)
	timer.ObserveDuration()
	if err == nil && update.Status != nil {
		incidentTotal.WithLabelValues(*update.Status).Inc()
		incidentTotal.WithLabelValues(incBefore.Status).Dec()
	}
	return incBefore, err
}
