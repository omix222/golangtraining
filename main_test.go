package main

import "testing"

func add(a, b int) int {
	return a + b
}

// TestXxx は Testxxx ではダメです。Test_xxx という関数名であれば問題ありません。
// コマンドライン上からは、go test -V で実行
// 参考　https://future-architect.github.io/articles/20200601/
func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal",
			args: args{a: 1, b: 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
