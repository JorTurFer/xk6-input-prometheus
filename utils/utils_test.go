package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	Case   string
	Start  string
	End    string
	Period string
	Error  string
}

var testCases = []TestData{
	{
		Case:   "Valid",
		Start:  "2006-01-02T15:04:05Z",
		End:    "2006-01-02T15:05:05Z",
		Period: "minute",
		Error:  "",
	},
	{
		Case:   "Empty Start",
		Start:  "",
		End:    "2006-01-02T15:04:05Z",
		Period: "minute",
		Error:  "start can't be empty",
	},
	{
		Case:   "Invalid Start",
		Start:  "2006-01-02 15:04:05",
		End:    "2006-01-02T15:05:05Z",
		Period: "minute",
		Error:  "parsing time \"2006-01-02 15:04:05\" as \"2006-01-02T15:04:05Z\": cannot parse \" 15:04:05\" as \"T\"",
	},
	{
		Case:   "Empty End",
		Start:  "2006-01-02T15:04:05Z",
		End:    "",
		Period: "minute",
		Error:  "end can't be empty",
	},
	{
		Case:   "Invalid End",
		Start:  "2006-01-02T15:04:05Z",
		End:    "2006-01-02 15:05:05",
		Period: "minute",
		Error:  "parsing time \"2006-01-02 15:05:05\" as \"2006-01-02T15:04:05Z\": cannot parse \" 15:05:05\" as \"T\"",
	},
	{
		Case:   "End before Start",
		Start:  "2006-01-02T15:07:05Z",
		End:    "2006-01-02T15:05:05Z",
		Period: "minute",
		Error:  "start can't be after the end",
	},
	{
		Case:   "Period second",
		Start:  "2006-01-02T15:03:05Z",
		End:    "2006-01-02T15:05:05Z",
		Period: "second",
		Error:  "",
	},
	{
		Case:   "Period minute",
		Start:  "2006-01-02T15:03:05Z",
		End:    "2006-01-02T15:05:05Z",
		Period: "minute",
		Error:  "",
	},
	{
		Case:   "Period hour",
		Start:  "2006-01-02T15:03:05Z",
		End:    "2006-01-02T15:05:05Z",
		Period: "hour",
		Error:  "",
	},
	{
		Case:   "Invalid Period",
		Start:  "2006-01-02T15:03:05Z",
		End:    "2006-01-02T15:05:05Z",
		Period: "random",
		Error:  "invalid period",
	},
}

func TestParseRange(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.Case, func(t *testing.T) {
			tc := tc
			t.Parallel()
			_, err := ParseRange(tc.Start, tc.End, tc.Period)

			if tc.Error == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.Error)
			}
		})
	}
}
