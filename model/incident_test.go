package model

import (
	"testing"

	"github.com/q231950/alethea/ci"
	"github.com/stretchr/testify/assert"
)

func TestNewIncident(t *testing.T) {
	incident, err := NewIncidentFromJson(ci.Circle, jsonCircleIncident())
	assert.Nil(t, err)
	assert.NotNil(t, incident.Identifier, "An incident's identifier should never be nil")
	assert.Equal(t, incident.CI, "Circle CI")
}

func jsonCircleIncident() []byte {
	str := `{
		"payload": {
		  "compare": "https://github.com/q231950/alethea/compare/ed7005e103e6...f753c9b5a0cb",
		  "previous_successful_build": {
			"build_num": 85,
			"status": "success",
			"build_time_millis": 10767
		  },
		  "build_parameters": null,
		  "oss": true,
		  "committer_date": "2018-02-10T23:17:00+01:00",
		  "steps": [
			{
			  "name": "Spin up Environment",
			  "actions": [
				{
				  "truncated": false,
				  "index": 0,
				  "parallel": true,
				  "failed": null,
				  "infrastructure_fail": null,
				  "name": "Spin up Environment",
				  "bash_command": null,
				  "status": "success",
				  "timedout": null,
				  "continue": null,
				  "end_time": "2018-02-24T18:32:31.296Z",
				  "type": "test",
				  "allocation_id": "5a91afb8c9e77c0001462a32-0-build/B63E590",
				  "output_url": "https://circle-production-action-output.s3.amazonaws.com/e3d2dd1000ddfd3edbfa19a5-q231950-alethea-0-0?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20180224T183245Z&X-Amz-SignedHeaders=host&X-Amz-Expires=431999&X-Amz-Credential=AKIAIQ65EYQDTMSJK2DQ%2F20180224%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Signature=57d7bd6866472683381fa7a11120543259039b82b2bc6c9341a2e90d208ba4dc",
				  "start_time": "2018-02-24T18:32:29.234Z",
				  "background": false,
				  "exit_code": null,
				  "insignificant": false,
				  "canceled": null,
				  "step": 0,
				  "run_time_millis": 2062,
				  "has_output": true
				}
			  ]
			},
			{
			  "name": "Container circleci/postgres:9.4",
			  "actions": [
				{
				  "truncated": false,
				  "index": 0,
				  "parallel": true,
				  "failed": null,
				  "infrastructure_fail": null,
				  "name": "Container circleci/postgres:9.4",
				  "bash_command": null,
				  "status": "running",
				  "timedout": null,
				  "continue": null,
				  "end_time": null,
				  "type": "test",
				  "allocation_id": "5a91afb8c9e77c0001462a32-0-build/B63E590",
				  "start_time": "2018-02-24T18:32:31.288Z",
				  "background": true,
				  "exit_code": null,
				  "insignificant": true,
				  "canceled": null,
				  "step": 1,
				  "run_time_millis": null,
				  "has_output": true
				}
			  ]
			},
			{
			  "name": "Checkout code",
			  "actions": [
				{
				  "truncated": false,
				  "index": 0,
				  "parallel": true,
				  "failed": null,
				  "infrastructure_fail": null,
				  "name": "Checkout code",
				  "bash_command": "#!/bin/sh\nset -e\n\n# Workaround old docker images with incorrect $HOME\n# check https://github.com/docker/docker/issues/2968 for details\nif [ \"${HOME}\" = \"/\" ]\nthen\n  export HOME=$(getent passwd $(id -un) | cut -d: -f6)\nfi\n\nmkdir -p ~/.ssh\n\necho 'github.com ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAq2A7hRGmdnm9tUDbO9IDSwBK6TbQa+PXYPCPy6rbTrTtw7PHkccKrpp0yVhp5HdEIcKr6pLlVDBfOLX9QUsyCOV0wzfjIJNlGEYsdlLJizHhbn2mUjvSAHQqZETYP81eFzLQNnPHt4EVVUh7VfDESU84KezmD5QlWpXLmvU31/yMf+Se8xhHTvKSCZIFImWwoG6mbUoWf9nzpIoaSjB+weqqUUmpaaasXVal72J+UX2B+2RPW3RcT0eOzQgqlJL3RKrTJvdsjE3JEAvGq3lGHSZXy28G3skua2SmVi/w4yCE6gbODqnTWlg7+wC604ydGXA8VJiS5ap43JXiUFFAaQ==\nbitbucket.org ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAubiN81eDcafrgMeLzaFPsw2kNvEcqTKl/VqLat/MaB33pZy0y3rJZtnqwR2qOOvbwKZYKiEO1O6VqNEBxKvJJelCq0dTXWT5pbO2gDXC6h6QDXCaHo6pOHGPUy+YBaGQRGuSusMEASYiWunYN0vCAI8QaXnWMXNMdFP3jHAJH0eDsoiGnLPBlBp4TNm6rYI74nMzgz3B9IikW4WVK+dc8KZJZWYjAuORU3jc1c/NPskD2ASinf8v3xnfXeukU0sJ5N6m5E8VLjObPEO+mN2t/FZTMZLiFqPWc/ALSqnMnnhwrNi2rbfg/rd/IpL8Le3pSBne8+seeFVBoGqzHM9yXw==\n' >> ~/.ssh/known_hosts\n\n(umask 077; touch ~/.ssh/id_rsa)\nchmod 0600 ~/.ssh/id_rsa\n(cat <<EOF > ~/.ssh/id_rsa\n$CHECKOUT_KEY\nEOF\n)\n\n# use git+ssh instead of https\ngit config --global url.\"ssh://git@github.com\".insteadOf \"https://github.com\" || true\n\nif [ -e /go/src/github.com/q231950/alethea/.git ]\nthen\n  cd /go/src/github.com/q231950/alethea\n  git remote set-url origin \"$CIRCLE_REPOSITORY_URL\" || true\nelse\n  mkdir -p /go/src/github.com/q231950/alethea\n  cd /go/src/github.com/q231950/alethea\n  git clone \"$CIRCLE_REPOSITORY_URL\" .\nfi\n\nif [ -n \"$CIRCLE_TAG\" ]\nthen\n  git fetch --force origin \"refs/tags/${CIRCLE_TAG}\"\nelse\n  git fetch --force origin \"failing_branch:remotes/origin/failing_branch\"\nfi\n\n\nif [ -n \"$CIRCLE_TAG\" ]\nthen\n  git reset --hard \"$CIRCLE_SHA1\"\n  git checkout -q \"$CIRCLE_TAG\"\nelif [ -n \"$CIRCLE_BRANCH\" ]\nthen\n  git reset --hard \"$CIRCLE_SHA1\"\n  git checkout -q -B \"$CIRCLE_BRANCH\"\nfi\n\ngit reset --hard \"$CIRCLE_SHA1\"",
				  "status": "success",
				  "timedout": null,
				  "continue": null,
				  "end_time": "2018-02-24T18:32:31.854Z",
				  "type": "test",
				  "allocation_id": "5a91afb8c9e77c0001462a32-0-build/B63E590",
				  "output_url": "https://circle-production-action-output.s3.amazonaws.com/04d2dd1000ddfd3efbfa19a5-q231950-alethea-101-0?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20180224T183245Z&X-Amz-SignedHeaders=host&X-Amz-Expires=431999&X-Amz-Credential=AKIAIQ65EYQDTMSJK2DQ%2F20180224%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Signature=12f71faefdbade59c8f1d793f3f0e939a862af4d436929633c5e34da608d6659",
				  "start_time": "2018-02-24T18:32:31.391Z",
				  "background": false,
				  "exit_code": 0,
				  "insignificant": false,
				  "canceled": null,
				  "step": 101,
				  "run_time_millis": 463,
				  "has_output": true
				}
			  ]
			},
			{
			  "name": "go get -v -t -d ./...",
			  "actions": [
				{
				  "truncated": false,
				  "index": 0,
				  "parallel": true,
				  "failed": null,
				  "infrastructure_fail": null,
				  "name": "go get -v -t -d ./...",
				  "bash_command": "#!/bin/bash -eo pipefail\ngo get -v -t -d ./...",
				  "status": "success",
				  "timedout": null,
				  "continue": null,
				  "end_time": "2018-02-24T18:32:43.357Z",
				  "type": "test",
				  "allocation_id": "5a91afb8c9e77c0001462a32-0-build/B63E590",
				  "output_url": "https://circle-production-action-output.s3.amazonaws.com/14d2dd1000ddfd3efbfa19a5-q231950-alethea-102-0?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20180224T183245Z&X-Amz-SignedHeaders=host&X-Amz-Expires=431999&X-Amz-Credential=AKIAIQ65EYQDTMSJK2DQ%2F20180224%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Signature=0e55ad295931380e94399b92a2b43431cf27d7bcf845fc951f55756fcbb07f5e",
				  "start_time": "2018-02-24T18:32:31.859Z",
				  "background": false,
				  "exit_code": 0,
				  "insignificant": false,
				  "canceled": null,
				  "step": 102,
				  "run_time_millis": 11498,
				  "has_output": true
				}
			  ]
			},
			{
			  "name": "go test -v ./...",
			  "actions": [
				{
				  "truncated": false,
				  "index": 0,
				  "parallel": true,
				  "failed": true,
				  "infrastructure_fail": null,
				  "name": "go test -v ./...",
				  "bash_command": "#!/bin/bash -eo pipefail\ngo test -v ./...",
				  "status": "failed",
				  "timedout": null,
				  "continue": null,
				  "end_time": "2018-02-24T18:32:45.608Z",
				  "type": "test",
				  "allocation_id": "5a91afb8c9e77c0001462a32-0-build/B63E590",
				  "start_time": "2018-02-24T18:32:43.364Z",
				  "background": false,
				  "exit_code": 2,
				  "insignificant": false,
				  "canceled": null,
				  "step": 103,
				  "run_time_millis": 2244,
				  "has_output": true
				}
			  ]
			}
		  ],
		  "body": "",
		  "usage_queued_at": "2018-02-24T18:32:24.599Z",
		  "fail_reason": null,
		  "retry_of": 84,
		  "reponame": "alethea",
		  "ssh_users": [],
		  "build_url": "https://circleci.com/gh/q231950/alethea/86",
		  "parallel": 1,
		  "failed": true,
		  "branch": "failing_branch",
		  "username": "q231950",
		  "author_date": "2018-02-10T23:17:00+01:00",
		  "why": "retry",
		  "user": {
			"is_user": true,
			"login": "q231950",
			"avatar_url": "https://avatars1.githubusercontent.com/u/1648215?v=4",
			"name": "Martin Kim Dung-Pham",
			"vcs_type": "github",
			"id": 1648215
		  },
		  "vcs_revision": "f753c9b5a0cb71ca4ee0257bd95da958c5dc4f2b",
		  "owners": [
			"q231950"
		  ],
		  "vcs_tag": null,
		  "pull_requests": [],
		  "build_num": 86,
		  "infrastructure_fail": false,
		  "committer_email": "kim@elbedev.com",
		  "has_artifacts": true,
		  "previous": {
			"build_num": 84,
			"status": "failed",
			"build_time_millis": 10334
		  },
		  "status": "failed",
		  "committer_name": "Martin Kim Dung-Pham",
		  "retries": null,
		  "subject": "Print on failing branch",
		  "vcs_type": "github",
		  "timedout": false,
		  "dont_build": null,
		  "lifecycle": "finished",
		  "no_dependency_cache": false,
		  "stop_time": "2018-02-24T18:32:45.614Z",
		  "ssh_disabled": true,
		  "build_time_millis": 16420,
		  "picard": {
			"build_agent": {
			  "image": null,
			  "properties": {
				"build_agent": "0.0.4666-05dad47",
				"executor": "docker"
			  }
			},
			"resource_class": {
			  "cpu": 2.0,
			  "ram": 4096,
			  "class": "medium"
			},
			"executor": "docker"
		  },
		  "circle_yml": {
			"string": "# Golang CircleCI 2.0 configuration file\n#\n# Check https://circleci.com/docs/2.0/language-go/ for more details\nversion: 2\njobs:\n  build:\n    docker:\n      - image: circleci/golang:1.9\n      - image: circleci/postgres:9.4\n      \n    working_directory: /go/src/github.com/q231950/alethea\n    steps:\n      - checkout\n      - run: go get -v -t -d ./...\n      - run: go test -v ./...\n\nnotify:\n  webhooks:\n    - url: https://alethea3000.herokuapp.com/print\n\ndeployment:\n  staging:\n    branch: master\n    heroku:\n      appname: alethea3000\n"
		  },
		  "messages": [],
		  "is_first_green_build": false,
		  "job_name": null,
		  "start_time": "2018-02-24T18:32:29.194Z",
		  "canceler": null,
		  "all_commit_details": [
			{
			  "committer_date": "2018-02-10T23:17:00+01:00",
			  "body": "",
			  "branch": "failing_branch",
			  "author_date": "2018-02-10T23:17:00+01:00",
			  "committer_email": "kim@elbedev.com",
			  "commit": "f753c9b5a0cb71ca4ee0257bd95da958c5dc4f2b",
			  "committer_login": "q231950",
			  "committer_name": "Martin Kim Dung-Pham",
			  "subject": "Print on failing branch",
			  "commit_url": "https://github.com/q231950/alethea/commit/f753c9b5a0cb71ca4ee0257bd95da958c5dc4f2b",
			  "author_login": "q231950",
			  "author_name": "Martin Kim Dung-Pham",
			  "author_email": "kim@elbedev.com"
			}
		  ],
		  "platform": "2.0",
		  "outcome": "failed",
		  "vcs_url": "https://github.com/q231950/alethea",
		  "author_name": "Martin Kim Dung-Pham",
		  "node": null,
		  "queued_at": "2018-02-24T18:32:24.626Z",
		  "canceled": false,
		  "author_email": "kim@elbedev.com"
		}
	  }`

	return []byte(str)
}
