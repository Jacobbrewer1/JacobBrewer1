package api

import (
	"testing"
)

func TestIssue_IsPullRequest(t *testing.T) {
	tests := []struct {
		name     string
		input    Issue
		expected bool
	}{
		{"is pull request", testerSetupIssue(true, false, false), true},
		{"is not pull request", testerSetupIssue(false, false, false), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool := tt.input.IsPullRequest()
			if gotBool != tt.expected {
				t.Errorf("IsPullRequest() = %v, expected %v", gotBool, tt.expected)
			}
		})
	}
}

func TestIssue_IsAssigned(t *testing.T) {
	tests := []struct {
		name     string
		input    Issue
		expected bool
	}{
		{"unassigned fully", testerSetupIssue(false, false, false), false},
		{"is single assigned", testerSetupIssue(false, true, false), true},
		{"is multi assigned", testerSetupIssue(false, false, true), true},
		{"blank user array", testerSetupBlankIssueAssigneesArray(), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool := tt.input.IsAssigned()
			if gotBool != tt.expected {
				t.Errorf("IsAssigned() = %v, expected %v", gotBool, tt.expected)
			}
		})
	}
}
