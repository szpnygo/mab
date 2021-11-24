package httpx

import (
	"reflect"
	"testing"
)

func Test_httpx_Params(t *testing.T) {
	tests := []struct {
		name string
		url  string
		args map[string]string
		want string
	}{
		{
			name: "httpx_Params_1",
			url:  "https://www.neobaran.com/api",
			args: map[string]string{},
			want: "https://www.neobaran.com/api",
		},
		{
			name: "httpx_Params_2",
			url:  "https://www.neobaran.com/api",
			args: map[string]string{
				"a": "b",
			},
			want: "https://www.neobaran.com/api?a=b",
		},
		{
			name: "httpx_Params_3",
			url:  "https://www.neobaran.com/api?a=c",
			args: map[string]string{
				"a": "b",
			},
			want: "https://www.neobaran.com/api?a=b",
		},
		{
			name: "httpx_Params_4",
			url:  "https://www.neobaran.com/api?a=c",
			args: map[string]string{
				"b": "d",
			},
			want: "https://www.neobaran.com/api?a=c&b=d",
		},
		{
			name: "httpx_Params_5",
			url:  "https://www.neobaran.com/api",
			args: map[string]string{
				"a": "b",
				"c": "d",
			},
			want: "https://www.neobaran.com/api?a=b&c=d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.url).Params(tt.args).URL(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpx.Params() = %v, want %v", got, tt.want)
			}
		})
	}
}
