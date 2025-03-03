package local_test

import (
	"testing"

	sdktesting "gosdk/internal/testing"
	"gosdk/pkg/storage/local"
)

func TestNewConfig(t *testing.T) {
	t.Setenv("LOCAL_PATH", "")

	tests := []struct {
		name string
		want bool
	}{
		{
			name: "ok",
			want: true,
		},
		{
			name: "panic",
			want: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.want {
				t.Setenv("LOCAL_PATH", "/tmp")
				sdktesting.Equals(t, local.NewConfig().LocalPath, "/tmp")
			}

			if !testCase.want {
				expectedMsg := "missing env LOCAL_PATH"

				defer func() {
					rec := recover()

					if rec == nil {
						t.Error("expected panic but got none")
					}

					if msg, ok := rec.(string); !ok || msg != expectedMsg {
						t.Errorf("expected panic message %q but got %v", expectedMsg, rec)
					}
				}()

				local.NewConfig()
			}
		})
	}
}
