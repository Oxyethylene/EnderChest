package util

import (
	"github.com/golang-jwt/jwt/v5"
	"reflect"
	"testing"
)

func TestParseToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    *Claims
		wantErr bool
	}{
		{
			name: "test_1",
			args: args{token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiJwYXNzd29yZCIsImV4cCI6MTcwNDgyNDEyNCwiaWF0IjoxNzA0ODEzMzI0fQ.dFg8P--JQyiLzeb3Bnvmz0JDOeRe2uMTVGSo1Y7eWGo"},
			want: &Claims{
				Id:               1,
				Username:         "admin",
				RegisteredClaims: jwt.RegisteredClaims{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
