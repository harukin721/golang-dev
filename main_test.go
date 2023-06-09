package main

import "testing"

// テスト関数は `Test` で始める
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		// `t.Errorf` はテスト失敗を報告する
		t.Errorf("got %q want %q", got, want)
	}
}
