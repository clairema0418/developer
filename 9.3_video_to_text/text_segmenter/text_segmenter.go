package text_segmenter

import (
	"strings"
	"unicode"
)

type TextSegment struct {
	Text      string
	StartTime float64
	EndTime   float64
}

func SegmentText(text string) []TextSegment {
	segments := []TextSegment{}
	words := strings.FieldsFunc(text, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSpace(r)
	})

	startTime := 0.0
	for _, word := range words {
		segment := TextSegment{
			Text:      word,
			StartTime: startTime,
			EndTime:   startTime + float64(len(word)),
		}
		segments = append(segments, segment)
		startTime = segment.EndTime
	}

	return segments
}