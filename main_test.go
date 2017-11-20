package main

import(
	"testing"

	"github.com/taylorzr/hibye/compare"
	"github.com/taylorzr/hibye/storage"
	"github.com/taylorzr/hibye/root"
)

func TestFindFired(t *testing.T) {
	old, _ := storage.Read("test_old_users.csv")
	new, _ := storage.Read("test_new_users.csv")

	fired := compare.FindFired(old, new)

	if len(fired) != 1 {
		t.Error()
	}

	if (fired[0] != root.User{ ID: 4242424, Name: "Zach Taylor" }) {
		t.Error()
	}
}

func TestNoFindFired(t *testing.T) {
	same, _ := storage.Read("test_old_users.csv")

	fired := compare.FindFired(same, same)

	if len(fired) != 0 {
		t.Error()
	}
}

func TestFindHired(t *testing.T) {
	old, _ := storage.Read("test_old_users.csv")
	new, _ := storage.Read("test_new_users.csv")

	fired := compare.FindHired(old, new)

	if len(fired) != 1 {
		t.Error()
	}

	if (fired[0] != root.User{ ID: 6666666, Name: "Taco Bell" }) {
		t.Error()
	}
}

func TestNoFindHired(t *testing.T) {
	same, _ := storage.Read("test_old_users.csv")

	fired := compare.FindHired(same, same)

	if len(fired) != 0 {
		t.Error()
	}
}
