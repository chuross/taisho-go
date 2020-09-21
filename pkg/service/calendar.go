package service

import (
	"context"

	"golang.org/x/xerrors"
	"google.golang.org/api/calendar/v3"
)

func CreateCalendarEvent(ctx context.Context, calendarID string, event *calendar.Event) error {
	calendarService, err := calendar.NewService(ctx)
	if err != nil {
		return xerrors.Errorf("calendar init error: %w", err)
	}

	if _, err := calendarService.Events.Insert(calendarID, event).Do(); err != nil {
		return xerrors.Errorf("calendar event insert error: %w", err)
	}

	return nil
}
