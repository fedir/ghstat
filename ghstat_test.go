// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/fedir/ghstat/github"
)

var jsonResponse = `
{
  "id": 3577919,
  "name": "beego",
  "full_name": "astaxie/beego",
  "owner": {
	"login": "astaxie",
	"id": 233907,
	"avatar_url": "https://avatars3.githubusercontent.com/u/233907?v=4",
	"gravatar_id": "",
	"url": "https://api.github.com/users/astaxie",
	"html_url": "https://github.com/astaxie",
	"followers_url": "https://api.github.com/users/astaxie/followers",
	"following_url": "https://api.github.com/users/astaxie/following{/other_user}",
	"gists_url": "https://api.github.com/users/astaxie/gists{/gist_id}",
	"starred_url": "https://api.github.com/users/astaxie/starred{/owner}{/repo}",
	"subscriptions_url": "https://api.github.com/users/astaxie/subscriptions",
	"organizations_url": "https://api.github.com/users/astaxie/orgs",
	"repos_url": "https://api.github.com/users/astaxie/repos",
	"events_url": "https://api.github.com/users/astaxie/events{/privacy}",
	"received_events_url": "https://api.github.com/users/astaxie/received_events",
	"type": "User",
	"site_admin": false
  },
  "private": false,
  "html_url": "https://github.com/astaxie/beego",
  "description": "beego is an open-source, high-performance web framework for the Go programming language.",
  "fork": false,
  "url": "https://api.github.com/repos/astaxie/beego",
  "forks_url": "https://api.github.com/repos/astaxie/beego/forks",
  "keys_url": "https://api.github.com/repos/astaxie/beego/keys{/key_id}",
  "collaborators_url": "https://api.github.com/repos/astaxie/beego/collaborators{/collaborator}",
  "teams_url": "https://api.github.com/repos/astaxie/beego/teams",
  "hooks_url": "https://api.github.com/repos/astaxie/beego/hooks",
  "issue_events_url": "https://api.github.com/repos/astaxie/beego/issues/events{/number}",
  "events_url": "https://api.github.com/repos/astaxie/beego/events",
  "assignees_url": "https://api.github.com/repos/astaxie/beego/assignees{/user}",
  "branches_url": "https://api.github.com/repos/astaxie/beego/branches{/branch}",
  "tags_url": "https://api.github.com/repos/astaxie/beego/tags",
  "blobs_url": "https://api.github.com/repos/astaxie/beego/git/blobs{/sha}",
  "git_tags_url": "https://api.github.com/repos/astaxie/beego/git/tags{/sha}",
  "git_refs_url": "https://api.github.com/repos/astaxie/beego/git/refs{/sha}",
  "trees_url": "https://api.github.com/repos/astaxie/beego/git/trees{/sha}",
  "statuses_url": "https://api.github.com/repos/astaxie/beego/statuses/{sha}",
  "languages_url": "https://api.github.com/repos/astaxie/beego/languages",
  "stargazers_url": "https://api.github.com/repos/astaxie/beego/stargazers",
  "contributors_url": "https://api.github.com/repos/astaxie/beego/contributors",
  "subscribers_url": "https://api.github.com/repos/astaxie/beego/subscribers",
  "subscription_url": "https://api.github.com/repos/astaxie/beego/subscription",
  "commits_url": "https://api.github.com/repos/astaxie/beego/commits{/sha}",
  "git_commits_url": "https://api.github.com/repos/astaxie/beego/git/commits{/sha}",
  "comments_url": "https://api.github.com/repos/astaxie/beego/comments{/number}",
  "issue_comment_url": "https://api.github.com/repos/astaxie/beego/issues/comments{/number}",
  "contents_url": "https://api.github.com/repos/astaxie/beego/contents/{+path}",
  "compare_url": "https://api.github.com/repos/astaxie/beego/compare/{base}...{head}",
  "merges_url": "https://api.github.com/repos/astaxie/beego/merges",
  "archive_url": "https://api.github.com/repos/astaxie/beego/{archive_format}{/ref}",
  "downloads_url": "https://api.github.com/repos/astaxie/beego/downloads",
  "issues_url": "https://api.github.com/repos/astaxie/beego/issues{/number}",
  "pulls_url": "https://api.github.com/repos/astaxie/beego/pulls{/number}",
  "milestones_url": "https://api.github.com/repos/astaxie/beego/milestones{/number}",
  "notifications_url": "https://api.github.com/repos/astaxie/beego/notifications{?since,all,participating}",
  "labels_url": "https://api.github.com/repos/astaxie/beego/labels{/name}",
  "releases_url": "https://api.github.com/repos/astaxie/beego/releases{/id}",
  "deployments_url": "https://api.github.com/repos/astaxie/beego/deployments",
  "created_at": "2012-02-29T02:32:08Z",
  "updated_at": "2018-03-10T14:56:13Z",
  "pushed_at": "2018-03-10T09:17:53Z",
  "git_url": "git://github.com/astaxie/beego.git",
  "ssh_url": "git@github.com:astaxie/beego.git",
  "clone_url": "https://github.com/astaxie/beego.git",
  "svn_url": "https://github.com/astaxie/beego",
  "homepage": "beego.me",
  "size": 4396,
  "stargazers_count": 14391,
  "watchers_count": 14391,
  "language": "Go",
  "has_issues": true,
  "has_projects": true,
  "has_downloads": true,
  "has_wiki": true,
  "has_pages": false,
  "forks_count": 3166,
  "mirror_url": null,
  "archived": false,
  "open_issues_count": 448,
  "license": {
	"key": "other",
	"name": "Other",
	"spdx_id": null,
	"url": null
  },
  "forks": 3166,
  "open_issues": 448,
  "watchers": 14391,
  "default_branch": "master",
  "network_count": 3166,
  "subscribers_count": 1042
}
`

func TestParseRepositoryData(t *testing.T) {
	repositoryDataExtracted := github.ParseRepositoryData([]byte(jsonResponse))
	repositoryDataExpected := &github.Repository{
		Name:       "beego",
		FullName:   "astaxie/beego",
		Watchers:   14391,
		Forks:      3166,
		OpenIssues: 448,
		Language:   "Go",
		CreatedAt:  time.Date(2012, 2, 29, 2, 32, 8, 0, time.UTC),
		HasIssues:  true,
	}
	if !reflect.DeepEqual(repositoryDataExtracted, repositoryDataExpected) {
		fmt.Printf("%v", repositoryDataExtracted)
		fmt.Printf("%v", repositoryDataExpected)
		t.Fail()
	}
}

func TestParseRepositoryDataNegative(t *testing.T) {
	repositoryDataExtracted := github.ParseRepositoryData([]byte(jsonResponse))
	repositoryDataExpected := &github.Repository{
		Name:       "beego",
		FullName:   "astaxie/beego",
		Watchers:   14391,
		Forks:      3167,
		OpenIssues: 448,
		Language:   "Go",
		CreatedAt:  time.Date(2012, 2, 29, 2, 32, 8, 0, time.UTC),
		HasIssues:  true,
	}
	if reflect.DeepEqual(repositoryDataExtracted, repositoryDataExpected) {
		fmt.Printf("%v", repositoryDataExtracted)
		fmt.Printf("%v", repositoryDataExpected)
		t.Fail()
	}
}

func TestMainProgram(t *testing.T) {
	main()
}

// TODO
func TestFillCSVData(t *testing.T) {
}
