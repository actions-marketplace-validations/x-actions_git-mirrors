/*
 * 码云 Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.3.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gitee

// 编辑评论
type PullRequestComments struct {
	Url              string     `json:"url,omitempty"`
	Id               int32      `json:"id,omitempty"`
	Path             string     `json:"path,omitempty"`
	Position         string     `json:"position,omitempty"`
	OriginalPosition string     `json:"original_position,omitempty"`
	CommitId         string     `json:"commit_id,omitempty"`
	OriginalCommitId string     `json:"original_commit_id,omitempty"`
	User             *UserBasic `json:"user,omitempty"`
	Body             string     `json:"body,omitempty"`
	CreatedAt        string     `json:"created_at,omitempty"`
	UpdatedAt        string     `json:"updated_at,omitempty"`
	HtmlUrl          string     `json:"html_url,omitempty"`
	PullRequestUrl   string     `json:"pull_request_url,omitempty"`
	Links            string     `json:"_links,omitempty"`
}