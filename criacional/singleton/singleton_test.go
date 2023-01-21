package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNewSingleton(t *testing.T) {
	tests := []struct {
		name string
		want *Singleton
	}{
		{
			name: "test one",
			want: &Singleton{
				Count: 1,
			},
		},
		{
			name: "test AddOne",
			want: &Singleton{
				Count: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSingleton()
			assert.Equal(t, got.AddOne(), tt.want.Count)
		})
	}
}
