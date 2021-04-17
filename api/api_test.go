package api

import (
	"context"
	"testing"
	"time"
)

func TestApi_Run(t *testing.T) {
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	type args struct {
		ctx  context.Context
		Port int
	}
	tests := []struct {
		name string
		p    *Api
		args args
	}{
		{"Timeout", new(Api), args{ctx, 4242}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Run(tt.args.ctx, tt.args.Port)
		})
	}
}
