package conf

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	fname := "test.conf"
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Setenv("env", "1"); err != nil {
		t.Fatal(err)
	}
	c, err := New(fname, string(b))
	if err != nil {
		t.Fatal(err)
	}
	if w := c.Alerts["os.high_cpu"].Warn.Text; w != `avg(q("avg:rate:os.cpu{host=ny-nexpose01}", "2m", "")) > 80` {
		t.Error("bad warn:", w)
	}
	if w := c.Alerts["m"].Crit.Text; w != `avg(q("", "", "")) > 1` {
		t.Errorf("bad crit: %v", w)
	}
	if w := c.Alerts["braceTest"].Crit.Text; w != `avg(q("o{t}", "", "")) > 1` {
		t.Errorf("bad crit: %v", w)
	}
}

func TestInvalid(t *testing.T) {
	fname := "broken.conf"
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Fatal(err)
	}
	_, err = New(fname, string(b))
	if err == nil {
		t.Error("expected error")
	}
}
