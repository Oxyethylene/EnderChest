package db

import (
	"reflect"
	"testing"
	"time"
)

func TestQueryUserById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		{
			name: "test_1",
			args: args{id: 1},
			want: &User{
				Id:         1,
				Username:   "admin",
				Password:   "p@ssword",
				Email:      "admin@example.com",
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserById(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}
