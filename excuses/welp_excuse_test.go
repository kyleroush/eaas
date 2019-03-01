package excuses

import (
	"reflect"
	"testing"

	"github.com/kyleroush/eaas/models"
)

func Test_getWelp(t *testing.T) {
	tests := []struct {
		name string
		want models.SimpleExcuse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWelp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getWelp() = %v, want %v", got, tt.want)
			}
		})
	}
}
