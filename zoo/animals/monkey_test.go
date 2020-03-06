// package animal用テストプログラム
package animals

import "testing"

func TestMonkeyFeed(t *testing.T){
	expect := "banana"
	actual := MonkeyFeed()

	if expect != actual{
		t.Errorf("%s != %s", expect, actual)
	}
}

