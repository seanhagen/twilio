package twiml

import (
	"encoding/xml"
	"testing"
)

var testVerbs = []struct {
	Value     interface{}
	ExpectXML string
}{
	{Value: Dial{Number: "5"}, ExpectXML: "<Dial>5</Dial>"},
	{
		Value: Dial{
			Number:                  "5",
			RecordingStatusCallback: "testing",
		},
		ExpectXML: "<Dial recordingStatusCallback=\"testing\">5</Dial>",
	},
	{
		Value: Dial{
			Number:                        "5",
			RecordingStatusCallback:       "testing",
			RecordingStatusCallbackMethod: "POST",
		},
		ExpectXML: "<Dial recordingStatusCallback=\"testing\" recordingStatusCallbackMethod=\"POST\">5</Dial>",
	},
}

func TestDialRecordingStatus(t *testing.T) {
	for idx, test := range testVerbs {
		out, err := xml.Marshal(test.Value)
		if err != nil {
			t.Errorf("Test %v failed: %v", idx, err)
		}

		got := string(out)
		if got != test.ExpectXML {
			t.Errorf("Test %v failed; expected %#v, got %#v", idx, test.ExpectXML, got)
		}
	}
}
