// dup2のテスト
package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDup2(t *testing.T) {
	var tests = []struct {
		args []string
		want string
	}{
		{
			[]string{"test1.txt"},
			"ファイル: test1.txt\t重複数: 2\t内容: fuga\n" +
				"ファイル: test1.txt\t重複数: 3\t内容: hoge\n",
		},
		{
			[]string{"test1.txt", "test1.txt"},
			"ファイル: test1.txt\t重複数: 4\t内容: fuga\n" +
				"ファイル: test1.txt\t重複数: 6\t内容: hoge\n" +
				"ファイル: test1.txt\t重複数: 2\t内容: piyo\n",
		},
		{
			[]string{"test1.txt", "test2.txt"},
			"ファイル: test1.txt\t重複数: 2\t内容: fuga\n" +
				"ファイル: test1.txt, test2.txt\t重複数: 4\t内容: hoge\n" +
				"ファイル: test2.txt\t重複数: 2\t内容: hogehoge\n",
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("dup2(%q)", test.args)

		out = new(bytes.Buffer)
		if err := dup2(test.args); err != nil {
			t.Errorf("%s failed: %v", descr, err)
			continue
		}

		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
