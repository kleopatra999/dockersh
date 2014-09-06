package main

import (
	"testing"
)

func Test_dockerPid_1(t *testing.T) {
	pid, err := dockerpid("testcontainer")
	if err != nil {
		t.Errorf("Error from dockerpid: %v", err)
	}
	if pid != 666 {
		t.Errorf("PID was %i expected 666", pid)
	}
}

func Test_dockerSha_1(t *testing.T) {
	sha, err := dockersha("testcontainer")
	if err != nil {
		t.Errorf("Error from dockersha: %v", err)
	}
	if sha != "666" {
		t.Errorf("SHA was %s expected 666", sha)
	}
}

func Test_dockerStart(t *testing.T) {
	c := Configuration{MountHome:true,MountHomeFrom:"/home/fred",MountHomeTo:"/home/fred"}
	pid, err := dockerstart(c, "name", "container", "dockersock", true, true, "internal", []string{"foo"}, []string{"bar"})
	if err != nil {
		t.Errorf("Error from dockerstart: %v", err)
	}
	if pid != 666 {
		t.Errorf("PID was %i expected 666", pid)
	}
}
