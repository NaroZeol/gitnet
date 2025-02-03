package utils

import (
	"math/rand"
	"os/exec"
	"testing"
)

func checkBranchNameWithGit(branchName string) bool {
	// use git check-ref-format --branch to check branch name
	cmd := exec.Command("git", "check-ref-format", "--branch", branchName)
	err := cmd.Run()
	return err == nil
}

func checkRefNameWithGit(refName string) bool {
	// use git check-ref-format to check ref name
	cmd := exec.Command("git", "check-ref-format", refName)
	err := cmd.Run()
	return err == nil
}

func randomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789\\/.@{[~^:?*\"")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestCheckRefName(t *testing.T) {
	tests := []string{
		"-",
		"refs/heads/master",
		"refs/heads/main",
		"refs/heads/HEAD",
		"refs/heads/ HEAD",
		"refs/heads/ invalid",
		"refs/heads/invalid.lock",
		"refs/heads/invalid.lock/",
		"refs/heads/invalid.lock/invalid",
		"refs/heads/invalid.lock/invalid.lock",
		"refs/heads/invalid.lock/invalid.lock/",
		"refs/heads/invalid.lock/invalid.lock/invalid",
		"refs/heads/invalid.lock/invalid.lock/invalid.lock",
	}

	// random ref names
	for i := 0; i < 10000; i++ {
		tests = append(tests, randomString(10))
	}

	for _, test := range tests {
		expected := checkRefNameWithGit(test)
		result := CheckRefName(test)

		if expected != result {
			t.Errorf("CheckRefName(%s) = %v, want %v", test, result, expected)
		}
	}
}

func TestCheckBranchName(t *testing.T) {
	tests := []string{
		"-",
		"master",
		"main",
		"refs/heads/master",
		"HEAD",
		" HEAD",
		" invalid",
		"invalid.lock",
		"new-branch",
	}

	// random branch names
	for i := 0; i < 10000; i++ {
		tests = append(tests, randomString(10))
	}

	for _, test := range tests {
		expected := checkBranchNameWithGit(test)
		result := CheckBranchName(test)

		if expected != result {
			t.Errorf("CheckBranchName(%s) = %v, want %v", test, result, expected)
		}
	}
}
