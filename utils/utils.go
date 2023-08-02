package utils

import (
	"errors"
	"time"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func ParseRange(start, end, period string) (v1.Range, error) {
	template := "2006-01-02T15:04:05Z"

	if start == "" {
		return v1.Range{}, errors.New("start can't be empty")
	}
	startTime, err := time.Parse(template, start)
	if err != nil {
		return v1.Range{}, err
	}
	if end == "" {
		return v1.Range{}, errors.New("end can't be empty")
	}
	endTime, err := time.Parse(template, end)
	if err != nil {
		return v1.Range{}, err
	}

	if startTime.Sub(endTime) >= 0 {
		return v1.Range{}, errors.New("start can't be after the end")
	}

	var periodTime time.Duration
	switch period {
	case "hour":
		periodTime = time.Hour
	case "minute":
		periodTime = time.Minute
	case "second":
		periodTime = time.Second
	default:
		return v1.Range{}, errors.New("invalid period")
	}

	r := v1.Range{
		Start: startTime,
		End:   endTime,
		Step:  periodTime,
	}
	return r, nil
}
