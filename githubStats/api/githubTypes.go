package api

import (
	"time"
)

type Timestamp struct {
	time.Time
}

type NewIssue struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	Assignee string `json:"assignee"`
}

// User represents a GitHub user.
type User struct {
	Login                   *string    `json:"login,omitempty"`
	ID                      *int64     `json:"id,omitempty"`
	NodeID                  *string    `json:"node_id,omitempty"`
	AvatarURL               *string    `json:"avatar_url,omitempty"`
	HTMLURL                 *string    `json:"html_url,omitempty"`
	GravatarID              *string    `json:"gravatar_id,omitempty"`
	Name                    *string    `json:"name,omitempty"`
	Company                 *string    `json:"company,omitempty"`
	Blog                    *string    `json:"blog,omitempty"`
	Location                *string    `json:"location,omitempty"`
	Email                   *string    `json:"email,omitempty"`
	Hireable                *bool      `json:"hireable,omitempty"`
	Bio                     *string    `json:"bio,omitempty"`
	TwitterUsername         *string    `json:"twitter_username,omitempty"`
	PublicRepos             *int       `json:"public_repos,omitempty"`
	PublicGists             *int       `json:"public_gists,omitempty"`
	Followers               *int       `json:"followers,omitempty"`
	Following               *int       `json:"following,omitempty"`
	CreatedAt               *Timestamp `json:"created_at,omitempty"`
	UpdatedAt               *Timestamp `json:"updated_at,omitempty"`
	SuspendedAt             *Timestamp `json:"suspended_at,omitempty"`
	Type                    *string    `json:"type,omitempty"`
	SiteAdmin               *bool      `json:"site_admin,omitempty"`
	TotalPrivateRepos       *int       `json:"total_private_repos,omitempty"`
	OwnedPrivateRepos       *int       `json:"owned_private_repos,omitempty"`
	PrivateGists            *int       `json:"private_gists,omitempty"`
	DiskUsage               *int       `json:"disk_usage,omitempty"`
	Collaborators           *int       `json:"collaborators,omitempty"`
	TwoFactorAuthentication *bool      `json:"two_factor_authentication,omitempty"`
	Plan                    *Plan      `json:"plan,omitempty"`
	LdapDn                  *string    `json:"ldap_dn,omitempty"`

	// API URLs
	URL               *string `json:"url,omitempty"`
	EventsURL         *string `json:"events_url,omitempty"`
	FollowingURL      *string `json:"following_url,omitempty"`
	FollowersURL      *string `json:"followers_url,omitempty"`
	GistsURL          *string `json:"gists_url,omitempty"`
	OrganizationsURL  *string `json:"organizations_url,omitempty"`
	ReceivedEventsURL *string `json:"received_events_url,omitempty"`
	ReposURL          *string `json:"repos_url,omitempty"`
	StarredURL        *string `json:"starred_url,omitempty"`
	SubscriptionsURL  *string `json:"subscriptions_url,omitempty"`

	// TextMatches is only populated from search results that request text matches
	// See: search.go and https://docs.github.com/en/free-pro-team@latest/rest/reference/search/#text-match-metadata
	TextMatches []*TextMatch `json:"text_matches,omitempty"`

	// Permissions identifies the permissions that a user has on a given
	// repository. This is only populated when calling Repositories.ListCollaborators.
	Permissions map[string]bool `json:"permissions,omitempty"`
}

// Issue represents a GitHub Issue on a repository.
//
// Note: As far as the GitHub API is concerned, every pull request is an Issue,
// but not every Issue is a pull request. Some endpoints, events, and webhooks
// may also return pull requests via this struct. If PullRequestLinks is nil,
// this is an Issue, and if PullRequestLinks is not nil, this is a pull request.
// The IsPullRequest helper method can be used to check that.
type Issue struct {
	ID                *int64            `json:"id,omitempty"`
	Number            *int              `json:"number,omitempty"`
	State             *string           `json:"state,omitempty"`
	Locked            *bool             `json:"locked,omitempty"`
	Title             *string           `json:"title,omitempty"`
	Body              *string           `json:"body,omitempty"`
	AuthorAssociation *string           `json:"author_association,omitempty"`
	User              *User             `json:"user,omitempty"`
	Labels            []*Label          `json:"labels,omitempty"`
	Assignee          *User             `json:"assignee,omitempty"`
	Comments          *int              `json:"comments,omitempty"`
	ClosedAt          *time.Time        `json:"closed_at,omitempty"`
	CreatedAt         *time.Time        `json:"created_at,omitempty"`
	UpdatedAt         *time.Time        `json:"updated_at,omitempty"`
	ClosedBy          *User             `json:"closed_by,omitempty"`
	URL               *string           `json:"url,omitempty"`
	HTMLURL           *string           `json:"html_url,omitempty"`
	CommentsURL       *string           `json:"comments_url,omitempty"`
	EventsURL         *string           `json:"events_url,omitempty"`
	LabelsURL         *string           `json:"labels_url,omitempty"`
	RepositoryURL     *string           `json:"repository_url,omitempty"`
	Milestone         *Milestone        `json:"milestone,omitempty"`
	PullRequestLinks  *PullRequestLinks `json:"pull_request,omitempty"`
	Repository        *Repository       `json:"repository,omitempty"`
	Reactions         *Reactions        `json:"reactions,omitempty"`
	Assignees         []*User           `json:"assignees,omitempty"`
	NodeID            *string           `json:"node_id,omitempty"`

	// TextMatches is only populated from search results that request text matches
	// See: search.go and https://docs.github.com/en/free-pro-team@latest/rest/reference/search/#text-match-metadata
	TextMatches []*TextMatch `json:"text_matches,omitempty"`

	// ActiveLockReason is populated only when LockReason is provided while locking the Issue.
	// Possible values are: "off-topic", "too heated", "resolved", and "spam".
	ActiveLockReason *string `json:"active_lock_reason,omitempty"`
}

// IsPullRequest reports whether the Issue is also a pull request. It uses the
// method recommended by GitHub's API documentation, which is to check whether
// PullRequestLinks is non-nil.
func (i Issue) IsPullRequest() bool {
	return i.PullRequestLinks != nil
}

// IsAssigned reports whether the Issue is assigned to a user. It checks to see
// whether the assigned user's username/Login is non-nil
func (i Issue) IsAssigned() bool {
	return i.Assignee != nil || (i.Assignees != nil && len(i.Assignees) > 0)
}

// PullRequestLinks object is added to the Issue object when it's an Issue included
// in the IssueCommentEvent webhook payload, if the webhook is fired by a comment on a PR.
type PullRequestLinks struct {
	URL      *string `json:"url,omitempty"`
	HTMLURL  *string `json:"html_url,omitempty"`
	DiffURL  *string `json:"diff_url,omitempty"`
	PatchURL *string `json:"patch_url,omitempty"`
}

// Milestone represents a GitHub repository milestone.
type Milestone struct {
	URL          *string    `json:"url,omitempty"`
	HTMLURL      *string    `json:"html_url,omitempty"`
	LabelsURL    *string    `json:"labels_url,omitempty"`
	ID           *int64     `json:"id,omitempty"`
	Number       *int       `json:"number,omitempty"`
	State        *string    `json:"state,omitempty"`
	Title        *string    `json:"title,omitempty"`
	Description  *string    `json:"description,omitempty"`
	Creator      *User      `json:"creator,omitempty"`
	OpenIssues   *int       `json:"open_issues,omitempty"`
	ClosedIssues *int       `json:"closed_issues,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	ClosedAt     *time.Time `json:"closed_at,omitempty"`
	DueOn        *time.Time `json:"due_on,omitempty"`
	NodeID       *string    `json:"node_id,omitempty"`
}

// Label represents a GitHub label on an Issue
type Label struct {
	ID          *int64  `json:"id,omitempty"`
	URL         *string `json:"url,omitempty"`
	Name        *string `json:"name,omitempty"`
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Default     *bool   `json:"default,omitempty"`
	NodeID      *string `json:"node_id,omitempty"`
}

// Plan represents the payment plan for an account. See plans at https://github.com/plans.
type Plan struct {
	Name          *string `json:"name,omitempty"`
	Space         *int    `json:"space,omitempty"`
	Collaborators *int    `json:"collaborators,omitempty"`
	PrivateRepos  *int    `json:"private_repos,omitempty"`
	FilledSeats   *int    `json:"filled_seats,omitempty"`
	Seats         *int    `json:"seats,omitempty"`
}

// TextMatch represents a text match for a SearchResult
type TextMatch struct {
	ObjectURL  *string  `json:"object_url,omitempty"`
	ObjectType *string  `json:"object_type,omitempty"`
	Property   *string  `json:"property,omitempty"`
	Fragment   *string  `json:"fragment,omitempty"`
	Matches    []*Match `json:"matches,omitempty"`
}

// Match represents a single text match.
type Match struct {
	Text    *string `json:"text,omitempty"`
	Indices []int   `json:"indices,omitempty"`
}

// Reactions represents a summary of GitHub reactions.
type Reactions struct {
	TotalCount *int    `json:"total_count,omitempty"`
	PlusOne    *int    `json:"+1,omitempty"`
	MinusOne   *int    `json:"-1,omitempty"`
	Laugh      *int    `json:"laugh,omitempty"`
	Confused   *int    `json:"confused,omitempty"`
	Heart      *int    `json:"heart,omitempty"`
	Hooray     *int    `json:"hooray,omitempty"`
	Rocket     *int    `json:"rocket,omitempty"`
	Eyes       *int    `json:"eyes,omitempty"`
	URL        *string `json:"url,omitempty"`
}

// Repository represents a GitHub repository.
type Repository struct {
	ID                  *int64          `json:"id,omitempty"`
	NodeID              *string         `json:"node_id,omitempty"`
	Owner               *User           `json:"owner,omitempty"`
	Name                *string         `json:"name,omitempty"`
	FullName            *string         `json:"full_name,omitempty"`
	Description         *string         `json:"description,omitempty"`
	Homepage            *string         `json:"homepage,omitempty"`
	CodeOfConduct       *CodeOfConduct  `json:"code_of_conduct,omitempty"`
	DefaultBranch       *string         `json:"default_branch,omitempty"`
	MasterBranch        *string         `json:"master_branch,omitempty"`
	CreatedAt           *Timestamp      `json:"created_at,omitempty"`
	PushedAt            *Timestamp      `json:"pushed_at,omitempty"`
	UpdatedAt           *Timestamp      `json:"updated_at,omitempty"`
	HTMLURL             *string         `json:"html_url,omitempty"`
	CloneURL            *string         `json:"clone_url,omitempty"`
	GitURL              *string         `json:"git_url,omitempty"`
	MirrorURL           *string         `json:"mirror_url,omitempty"`
	SSHURL              *string         `json:"ssh_url,omitempty"`
	SVNURL              *string         `json:"svn_url,omitempty"`
	Language            *string         `json:"language,omitempty"`
	Fork                *bool           `json:"fork,omitempty"`
	ForksCount          *int            `json:"forks_count,omitempty"`
	NetworkCount        *int            `json:"network_count,omitempty"`
	OpenIssuesCount     *int            `json:"open_issues_count,omitempty"`
	OpenIssues          *int            `json:"open_issues,omitempty"` // Deprecated: Replaced by OpenIssuesCount. For backward compatibility OpenIssues is still populated.
	StargazersCount     *int            `json:"stargazers_count,omitempty"`
	SubscribersCount    *int            `json:"subscribers_count,omitempty"`
	WatchersCount       *int            `json:"watchers_count,omitempty"` // Deprecated: Replaced by StargazersCount. For backward compatibility WatchersCount is still populated.
	Watchers            *int            `json:"watchers,omitempty"`       // Deprecated: Replaced by StargazersCount. For backward compatibility Watchers is still populated.
	Size                *int            `json:"size,omitempty"`
	AutoInit            *bool           `json:"auto_init,omitempty"`
	Parent              *Repository     `json:"parent,omitempty"`
	Source              *Repository     `json:"source,omitempty"`
	TemplateRepository  *Repository     `json:"template_repository,omitempty"`
	Organization        *Organization   `json:"organization,omitempty"`
	Permissions         map[string]bool `json:"permissions,omitempty"`
	AllowRebaseMerge    *bool           `json:"allow_rebase_merge,omitempty"`
	AllowSquashMerge    *bool           `json:"allow_squash_merge,omitempty"`
	AllowMergeCommit    *bool           `json:"allow_merge_commit,omitempty"`
	AllowAutoMerge      *bool           `json:"allow_auto_merge,omitempty"`
	AllowForking        *bool           `json:"allow_forking,omitempty"`
	DeleteBranchOnMerge *bool           `json:"delete_branch_on_merge,omitempty"`
	Topics              []string        `json:"topics,omitempty"`
	Archived            *bool           `json:"archived,omitempty"`
	Disabled            *bool           `json:"disabled,omitempty"`

	// Only provided when using RepositoriesService.Get while in preview
	License *License `json:"license,omitempty"`

	// Additional mutable fields when creating and editing a repository
	Private           *bool   `json:"private,omitempty"`
	HasIssues         *bool   `json:"has_issues,omitempty"`
	HasWiki           *bool   `json:"has_wiki,omitempty"`
	HasPages          *bool   `json:"has_pages,omitempty"`
	HasProjects       *bool   `json:"has_projects,omitempty"`
	HasDownloads      *bool   `json:"has_downloads,omitempty"`
	IsTemplate        *bool   `json:"is_template,omitempty"`
	LicenseTemplate   *string `json:"license_template,omitempty"`
	GitignoreTemplate *string `json:"gitignore_template,omitempty"`

	// Options for configuring Advanced Security and Secret Scanning
	SecurityAndAnalysis *SecurityAndAnalysis `json:"security_and_analysis,omitempty"`

	// Creating an organization repository. Required for non-owners.
	TeamID *int64 `json:"team_id,omitempty"`

	// API URLs
	URL              *string `json:"url,omitempty"`
	ArchiveURL       *string `json:"archive_url,omitempty"`
	AssigneesURL     *string `json:"assignees_url,omitempty"`
	BlobsURL         *string `json:"blobs_url,omitempty"`
	BranchesURL      *string `json:"branches_url,omitempty"`
	CollaboratorsURL *string `json:"collaborators_url,omitempty"`
	CommentsURL      *string `json:"comments_url,omitempty"`
	CommitsURL       *string `json:"commits_url,omitempty"`
	CompareURL       *string `json:"compare_url,omitempty"`
	ContentsURL      *string `json:"contents_url,omitempty"`
	ContributorsURL  *string `json:"contributors_url,omitempty"`
	DeploymentsURL   *string `json:"deployments_url,omitempty"`
	DownloadsURL     *string `json:"downloads_url,omitempty"`
	EventsURL        *string `json:"events_url,omitempty"`
	ForksURL         *string `json:"forks_url,omitempty"`
	GitCommitsURL    *string `json:"git_commits_url,omitempty"`
	GitRefsURL       *string `json:"git_refs_url,omitempty"`
	GitTagsURL       *string `json:"git_tags_url,omitempty"`
	HooksURL         *string `json:"hooks_url,omitempty"`
	IssueCommentURL  *string `json:"issue_comment_url,omitempty"`
	IssueEventsURL   *string `json:"issue_events_url,omitempty"`
	IssuesURL        *string `json:"issues_url,omitempty"`
	KeysURL          *string `json:"keys_url,omitempty"`
	LabelsURL        *string `json:"labels_url,omitempty"`
	LanguagesURL     *string `json:"languages_url,omitempty"`
	MergesURL        *string `json:"merges_url,omitempty"`
	MilestonesURL    *string `json:"milestones_url,omitempty"`
	NotificationsURL *string `json:"notifications_url,omitempty"`
	PullsURL         *string `json:"pulls_url,omitempty"`
	ReleasesURL      *string `json:"releases_url,omitempty"`
	StargazersURL    *string `json:"stargazers_url,omitempty"`
	StatusesURL      *string `json:"statuses_url,omitempty"`
	SubscribersURL   *string `json:"subscribers_url,omitempty"`
	SubscriptionURL  *string `json:"subscription_url,omitempty"`
	TagsURL          *string `json:"tags_url,omitempty"`
	TreesURL         *string `json:"trees_url,omitempty"`
	TeamsURL         *string `json:"teams_url,omitempty"`

	// TextMatches is only populated from search results that request text matches
	// See: search.go and https://docs.github.com/en/free-pro-team@latest/rest/reference/search/#text-match-metadata
	TextMatches []*TextMatch `json:"text_matches,omitempty"`

	// Visibility is only used for Create and Edit endpoints. The visibility field
	// overrides the field parameter when both are used.
	// Can be one of public, private or internal.
	Visibility *string `json:"visibility,omitempty"`
}

// CodeOfConduct represents a code of conduct.
type CodeOfConduct struct {
	Name *string `json:"name,omitempty"`
	Key  *string `json:"key,omitempty"`
	URL  *string `json:"url,omitempty"`
	Body *string `json:"body,omitempty"`
}

// Organization represents a GitHub organization account.
type Organization struct {
	Login                       *string    `json:"login,omitempty"`
	ID                          *int64     `json:"id,omitempty"`
	NodeID                      *string    `json:"node_id,omitempty"`
	AvatarURL                   *string    `json:"avatar_url,omitempty"`
	HTMLURL                     *string    `json:"html_url,omitempty"`
	Name                        *string    `json:"name,omitempty"`
	Company                     *string    `json:"company,omitempty"`
	Blog                        *string    `json:"blog,omitempty"`
	Location                    *string    `json:"location,omitempty"`
	Email                       *string    `json:"email,omitempty"`
	TwitterUsername             *string    `json:"twitter_username,omitempty"`
	Description                 *string    `json:"description,omitempty"`
	PublicRepos                 *int       `json:"public_repos,omitempty"`
	PublicGists                 *int       `json:"public_gists,omitempty"`
	Followers                   *int       `json:"followers,omitempty"`
	Following                   *int       `json:"following,omitempty"`
	CreatedAt                   *time.Time `json:"created_at,omitempty"`
	UpdatedAt                   *time.Time `json:"updated_at,omitempty"`
	TotalPrivateRepos           *int       `json:"total_private_repos,omitempty"`
	OwnedPrivateRepos           *int       `json:"owned_private_repos,omitempty"`
	PrivateGists                *int       `json:"private_gists,omitempty"`
	DiskUsage                   *int       `json:"disk_usage,omitempty"`
	Collaborators               *int       `json:"collaborators,omitempty"`
	BillingEmail                *string    `json:"billing_email,omitempty"`
	Type                        *string    `json:"type,omitempty"`
	Plan                        *Plan      `json:"plan,omitempty"`
	TwoFactorRequirementEnabled *bool      `json:"two_factor_requirement_enabled,omitempty"`
	IsVerified                  *bool      `json:"is_verified,omitempty"`
	HasOrganizationProjects     *bool      `json:"has_organization_projects,omitempty"`
	HasRepositoryProjects       *bool      `json:"has_repository_projects,omitempty"`

	// DefaultRepoPermission can be one of: "read", "write", "admin", or "none". (Default: "read").
	// It is only used in OrganizationsService.Edit.
	DefaultRepoPermission *string `json:"default_repository_permission,omitempty"`
	// DefaultRepoSettings can be one of: "read", "write", "admin", or "none". (Default: "read").
	// It is only used in OrganizationsService.Get.
	DefaultRepoSettings *string `json:"default_repository_settings,omitempty"`

	// MembersCanCreateRepos default value is true and is only used in Organizations.Edit.
	MembersCanCreateRepos *bool `json:"members_can_create_repositories,omitempty"`

	// https://developer.github.com/changes/2019-12-03-internal-visibility-changes/#rest-v3-api
	MembersCanCreatePublicRepos   *bool `json:"members_can_create_public_repositories,omitempty"`
	MembersCanCreatePrivateRepos  *bool `json:"members_can_create_private_repositories,omitempty"`
	MembersCanCreateInternalRepos *bool `json:"members_can_create_internal_repositories,omitempty"`

	// MembersAllowedRepositoryCreationType denotes if organization members can create repositories
	// and the type of repositories they can create. Possible values are: "all", "private", or "none".
	//
	// Deprecated: Use MembersCanCreatePublicRepos, MembersCanCreatePrivateRepos, MembersCanCreateInternalRepos
	// instead. The new fields overrides the existing MembersAllowedRepositoryCreationType during 'edit'
	// operation and does not consider 'internal' repositories during 'get' operation
	MembersAllowedRepositoryCreationType *string `json:"members_allowed_repository_creation_type,omitempty"`

	// MembersCanCreatePages toggles whether organization members can create GitHub Pages sites.
	MembersCanCreatePages *bool `json:"members_can_create_pages,omitempty"`
	// MembersCanCreatePublicPages toggles whether organization members can create public GitHub Pages sites.
	MembersCanCreatePublicPages *bool `json:"members_can_create_public_pages,omitempty"`
	// MembersCanCreatePrivatePages toggles whether organization members can create private GitHub Pages sites.
	MembersCanCreatePrivatePages *bool `json:"members_can_create_private_pages,omitempty"`

	// API URLs
	URL              *string `json:"url,omitempty"`
	EventsURL        *string `json:"events_url,omitempty"`
	HooksURL         *string `json:"hooks_url,omitempty"`
	IssuesURL        *string `json:"issues_url,omitempty"`
	MembersURL       *string `json:"members_url,omitempty"`
	PublicMembersURL *string `json:"public_members_url,omitempty"`
	ReposURL         *string `json:"repos_url,omitempty"`
}

// License represents an open source license.
type License struct {
	Key  *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`

	SPDXID         *string   `json:"spdx_id,omitempty"`
	HTMLURL        *string   `json:"html_url,omitempty"`
	Featured       *bool     `json:"featured,omitempty"`
	Description    *string   `json:"description,omitempty"`
	Implementation *string   `json:"implementation,omitempty"`
	Permissions    *[]string `json:"permissions,omitempty"`
	Conditions     *[]string `json:"conditions,omitempty"`
	Limitations    *[]string `json:"limitations,omitempty"`
	Body           *string   `json:"body,omitempty"`
}

// SecurityAndAnalysis specifies the optional advanced security features
// that are enabled on a given repository.
type SecurityAndAnalysis struct {
	AdvancedSecurity *AdvancedSecurity `json:"advanced_security,omitempty"`
	SecretScanning   *SecretScanning   `json:"secret_scanning,omitempty"`
}

// AdvancedSecurity specifies the state of advanced security on a repository.
//
// GitHub API docs: https://docs.github.com/en/github/getting-started-with-github/learning-about-github/about-github-advanced-security
type AdvancedSecurity struct {
	Status *string `json:"status,omitempty"`
}

// SecretScanning specifies the state of secret scanning on a repository.
//
// GitHub API docs: https://docs.github.com/en/code-security/secret-security/about-secret-scanning
type SecretScanning struct {
	Status *string `json:"status,omitempty"`
}

// Branch represents a Branch in Github
type Branch struct {
	Name          *string     `json:"name,omitempty"`
	Commit        *Commit     `json:"commit,omitempty"`
	Protected     *bool       `json:"protected,omitempty"`
	Protection    *Protection `json:"protection,omitempty"`
	ProtectionURL *string     `json:"protection_url,omitempty"`
}

func (b Branch) IsProtected() bool {
	return *b.Protected
}

type Protection struct {
	RequiredStatusChecks *RequiredStatusChecks `json:"required_status_checks,omitempty"`
}

type RequiredStatusChecks struct {
	EnforcementLevel *string   `json:"enforcement_level,omitempty"`
	Contexts         []*string `json:"contexts,omitempty"`
	Checks           []*Checks `json:"checks,omitempty"`
}
type Checks struct {
	Context *string `json:"context,omitempty"`
	AppId   *int64  `json:"app_id,omitempty"`
}

type Commit struct {
	sha *string `json:"sha,omitempty"`
	URL *string `json:"url,omitempty"`
}
