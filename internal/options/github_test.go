package options

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetGitHubOptions(t *testing.T) {
	testBranch := func(branchName string) {
		_ = os.Setenv("GITHUB_ACTIONS", "true")
		_ = os.Setenv("GITHUB_WORKFLOW", "workflow")
		_ = os.Setenv("GITHUB_ACTION", "action")
		_ = os.Setenv("GITHUB_REPOSITORY", "repository")
		_ = os.Setenv("GITHUB_EVENT_NAME", "event-name")
		_ = os.Setenv("GITHUB_SHA", "sha")
		_ = os.Setenv("GITHUB_REF", "refs/heads/"+branchName)

		github, err := GetGitHubOptions()
		assert.NilError(t, err)
		assert.DeepEqual(t, GitHub{
			RunInActions: true,
			Workflow:     "workflow",
			Action:       "action",
			Repository:   "repository",
			EventName:    "event-name",
			Sha:          "sha",
			Reference: GitReference{
				Type: GitRefHead,
				Name: branchName,
			},
		}, github)
	}

	for _, b := range []string{"master", "main"} {
		testBranch(b)
	}
}

func TestParseGitRef(t *testing.T) {
	testCases := []struct {
		name         string
		ref          string
		expectedType GitReferenceType
		expectedName string
	}{
		{
			name:         "master-branch",
			ref:          "refs/heads/master",
			expectedType: GitRefHead,
			expectedName: "master",
		},
		{
			name:         "main-branch",
			ref:          "refs/heads/main",
			expectedType: GitRefHead,
			expectedName: "main",
		},
		{
			name:         "different-branch",
			ref:          "refs/heads/different",
			expectedType: GitRefHead,
			expectedName: "different",
		},
		{
			name:         "pull-request",
			ref:          "refs/pull/pr1",
			expectedType: GitRefPullRequest,
			expectedName: "pr1",
		},
		{
			name:         "tag",
			ref:          "refs/tags/tag1",
			expectedType: GitRefTag,
			expectedName: "tag1",
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ref := parseGitRef(tc.ref)
			assert.DeepEqual(t, GitReference{Type: tc.expectedType, Name: tc.expectedName}, ref)
		})
	}
}

func TestGetGitHubOptionsNotInActions(t *testing.T) {
	_ = os.Unsetenv("GITHUB_ACTIONS")
	github, err := GetGitHubOptions()
	assert.NilError(t, err)
	assert.Equal(t, false, github.RunInActions)
}
