package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleExcuse_BuildText(t *testing.T) {

	tests := []struct {
		name string
		args Input
		want string
	}{
		{
			name: "text has just the message",
			want: "message",
			args: Input{},
		},
		{
			name: "text has 'to' when to is given",
			want: "Dear :to, message",
			args: Input{
				Inputs: map[string]string{
					"to": ":to",
				},
			},
		},
		{
			name: "text has 'from' when from is given",
			want: "message Sincerly :from",
			args: Input{
				Inputs: map[string]string{
					"from": ":from",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			excuse := SimpleExcuse{
				Key:     "key",
				Message: "message",
			}
			assert.Equal(t, excuse.BuildText(tt.args), tt.want)
		})
	}
}
