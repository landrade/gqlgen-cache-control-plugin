package cache

import (
	"context"
	"reflect"
	"testing"
)

func TestContextWithForceCacheControl(t *testing.T) {
	type args struct {
		ctx               context.Context
		forceCacheControl bool
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			"forceCacheControl is false",
			args{context.Background(), false},
			context.WithValue(context.Background(), forceCacheControlKey, false),
		},
		{
			"valid key and value forceCacheControl",
			args{context.Background(), true},
			context.WithValue(context.Background(), forceCacheControlKey, true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContextWithForceCacheControl(tt.args.ctx, tt.args.forceCacheControl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContextWithForceCacheControl() = %v, want %v", got, tt.want)
			}
		})
	}
}
