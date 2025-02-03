package utils

import (
	"regexp"
)

func CheckRefName(refName string) bool {
	// based on https://stackoverflow.com/questions/3651860/which-characters-are-illegal-within-a-branch-name
	invaildPattern := `^[\./]|\.\.|@{|[\/\.]$|^@$|[~^:\x00-\x20\x7F\s?*[\\]|\.lock$|\.lock/|/\.|//|/\.|^[^/]*$`

	if matched, _ := regexp.MatchString(invaildPattern, refName); matched {
		return false
	}

	return true
}

// branchName is the short name of a branch, e.g. "master" instead of "refs/heads/master"
func CheckBranchName(branchName string) bool {
	if branchName == "" {
		return false
	}

	if branchName == "-" {
		return false
	}

	if branchName == "HEAD" {
		return false
	}

	return CheckRefName("refs/heads/" + branchName)
}
