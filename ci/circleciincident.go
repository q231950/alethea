package ci

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

type CircleCIIncident struct {
	Payload struct {
		Compare                 string `json:"compare"`
		PreviousSuccessfulBuild struct {
			BuildNum        int    `json:"build_num"`
			Status          string `json:"status"`
			BuildTimeMillis int    `json:"build_time_millis"`
		} `json:"previous_successful_build"`
		BuildParameters interface{} `json:"build_parameters"`
		Oss             bool        `json:"oss"`
		CommitterDate   time.Time   `json:"committer_date"`
		Steps           []struct {
			Name    string `json:"name"`
			Actions []struct {
				Truncated          bool        `json:"truncated"`
				Index              int         `json:"index"`
				Parallel           bool        `json:"parallel"`
				Failed             interface{} `json:"failed"`
				InfrastructureFail interface{} `json:"infrastructure_fail"`
				Name               string      `json:"name"`
				BashCommand        interface{} `json:"bash_command"`
				Status             string      `json:"status"`
				Timedout           interface{} `json:"timedout"`
				Continue           interface{} `json:"continue"`
				EndTime            time.Time   `json:"end_time"`
				Type               string      `json:"type"`
				AllocationID       string      `json:"allocation_id"`
				OutputURL          string      `json:"output_url"`
				StartTime          time.Time   `json:"start_time"`
				Background         bool        `json:"background"`
				ExitCode           interface{} `json:"exit_code"`
				Insignificant      bool        `json:"insignificant"`
				Canceled           interface{} `json:"canceled"`
				Step               int         `json:"step"`
				RunTimeMillis      int         `json:"run_time_millis"`
				HasOutput          bool        `json:"has_output"`
			} `json:"actions"`
		} `json:"steps"`
		Body          string        `json:"body"`
		UsageQueuedAt time.Time     `json:"usage_queued_at"`
		FailReason    interface{}   `json:"fail_reason"`
		RetryOf       int           `json:"retry_of"`
		Reponame      string        `json:"reponame"`
		SSHUsers      []interface{} `json:"ssh_users"`
		BuildURL      string        `json:"build_url"`
		Parallel      int           `json:"parallel"`
		Failed        bool          `json:"failed"`
		Branch        string        `json:"branch"`
		Username      string        `json:"username"`
		AuthorDate    time.Time     `json:"author_date"`
		Why           string        `json:"why"`
		User          struct {
			IsUser    bool   `json:"is_user"`
			Login     string `json:"login"`
			AvatarURL string `json:"avatar_url"`
			Name      string `json:"name"`
			VcsType   string `json:"vcs_type"`
			ID        int    `json:"id"`
		} `json:"user"`
		VcsRevision        string        `json:"vcs_revision"`
		Owners             []string      `json:"owners"`
		VcsTag             interface{}   `json:"vcs_tag"`
		PullRequests       []interface{} `json:"pull_requests"`
		BuildNum           int           `json:"build_num"`
		InfrastructureFail bool          `json:"infrastructure_fail"`
		CommitterEmail     string        `json:"committer_email"`
		HasArtifacts       bool          `json:"has_artifacts"`
		Previous           struct {
			BuildNum        int    `json:"build_num"`
			Status          string `json:"status"`
			BuildTimeMillis int    `json:"build_time_millis"`
		} `json:"previous"`
		Status            string      `json:"status"`
		CommitterName     string      `json:"committer_name"`
		Retries           interface{} `json:"retries"`
		Subject           string      `json:"subject"`
		VcsType           string      `json:"vcs_type"`
		Timedout          bool        `json:"timedout"`
		DontBuild         interface{} `json:"dont_build"`
		Lifecycle         string      `json:"lifecycle"`
		NoDependencyCache bool        `json:"no_dependency_cache"`
		StopTime          time.Time   `json:"stop_time"`
		SSHDisabled       bool        `json:"ssh_disabled"`
		BuildTimeMillis   int         `json:"build_time_millis"`
		Picard            struct {
			BuildAgent struct {
				Image      interface{} `json:"image"`
				Properties struct {
					BuildAgent string `json:"build_agent"`
					Executor   string `json:"executor"`
				} `json:"properties"`
			} `json:"build_agent"`
			ResourceClass struct {
				CPU   float64 `json:"cpu"`
				RAM   int     `json:"ram"`
				Class string  `json:"class"`
			} `json:"resource_class"`
			Executor string `json:"executor"`
		} `json:"picard"`
		CircleYml struct {
			String string `json:"string"`
		} `json:"circle_yml"`
		Messages          []interface{} `json:"messages"`
		IsFirstGreenBuild bool          `json:"is_first_green_build"`
		JobName           interface{}   `json:"job_name"`
		StartTime         time.Time     `json:"start_time"`
		Canceler          interface{}   `json:"canceler"`
		AllCommitDetails  []struct {
			CommitterDate  time.Time `json:"committer_date"`
			Body           string    `json:"body"`
			Branch         string    `json:"branch"`
			AuthorDate     time.Time `json:"author_date"`
			CommitterEmail string    `json:"committer_email"`
			Commit         string    `json:"commit"`
			CommitterLogin string    `json:"committer_login"`
			CommitterName  string    `json:"committer_name"`
			Subject        string    `json:"subject"`
			CommitURL      string    `json:"commit_url"`
			AuthorLogin    string    `json:"author_login"`
			AuthorName     string    `json:"author_name"`
			AuthorEmail    string    `json:"author_email"`
		} `json:"all_commit_details"`
		Platform    string      `json:"platform"`
		Outcome     string      `json:"outcome"`
		VcsURL      string      `json:"vcs_url"`
		AuthorName  string      `json:"author_name"`
		Node        interface{} `json:"node"`
		QueuedAt    time.Time   `json:"queued_at"`
		Canceled    bool        `json:"canceled"`
		AuthorEmail string      `json:"author_email"`
	} `json:"payload"`
}

func (ci *CircleCIIncident) CI() string {
	return Circle.String()
}

func (ci *CircleCIIncident) Identifier() string {
	return uuid.New().String()
}

func (ci *CircleCIIncident) Failed() bool {
	return ci.Payload.Failed
}

func (ci *CircleCIIncident) Committer() string {
	return ci.Payload.CommitterName
}

func (ci *CircleCIIncident) Project() string {
	return ci.Payload.Reponame
}

func (ci *CircleCIIncident) BuildUrl() string {
	return ci.Payload.BuildURL
}

func (ci *CircleCIIncident) BuildNumber() string {
	return strconv.Itoa(ci.Payload.BuildNum)
}
