package time_recorder

import (
	"fmt"
	"strings"
	"time"
)

type TimeSegment struct {
	Text      string
	StartTime time.Duration
	EndTime   time.Duration
}

func RecordTimeSegments(transcript string, timestamps []time.Duration) ([]TimeSegment, error) {
	segments := strings.Split(transcript, ".")
	if len(segments)-1 != len(timestamps) {
		return nil, fmt.Errorf("timestamps and segments count mismatch")
	}

	timeSegments := make([]TimeSegment, len(timestamps))
	for i := 0; i < len(timestamps); i++ {
		timeSegments[i] = TimeSegment{
			Text:      strings.TrimSpace(segments[i]),
			StartTime: timestamps[i],
			EndTime:   timestamps[i+1],
		}
	}

	return timeSegments, nil
}

func FindSegmentByTime(timeSegments []TimeSegment, targetTime time.Duration) (*TimeSegment, error) {
	for _, segment := range timeSegments {
		if targetTime >= segment.StartTime && targetTime <= segment.EndTime {
			return &segment, nil
		}
	}
	return nil, fmt.Errorf("no segment found for the given time")
}

func FindSegmentByText(timeSegments []TimeSegment, targetText string) (*TimeSegment, error) {
	for _, segment := range timeSegments {
		if strings.Contains(segment.Text, targetText) {
			return &segment, nil
		}
	}
	return nil, fmt.Errorf("no segment found containing the given text")
}