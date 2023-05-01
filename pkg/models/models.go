package models

type RepoItem struct {
	GitURL     string `json:"git_url,omitempty" yaml:"git_url,omitempty"`
	Branch     string `json:"branch,omitempty" yaml:"branch,omitempty"`
	BuildCMD   string `json:"build_cmd,omitempty" yaml:"build_cmd,omitempty"`
	BprintFile string `json:"bprint_file,omitempty" yaml:"bprint_file,omitempty"`
}

type BuildConfig struct {
	Items        map[string]RepoItem `json:"items,omitempty" yaml:"items,omitempty"`
	BuildFolder  string              `json:"build_folder,omitempty" yaml:"build_folder,omitempty"`
	OutputFolder string              `json:"output_folder,omitempty" yaml:"output_folder,omitempty"`
}

type Bprint struct {
	Name        string         `yaml:"name,omitempty"`
	Slug        string         `yaml:"slug,omitempty"`
	Type        string         `yaml:"type,omitempty"`
	Description string         `yaml:"description,omitempty"`
	Icon        string         `yaml:"icon,omitempty"`
	Version     []string       `yaml:"versions,omitempty"`
	Tags        []string       `yaml:"tags,omitempty"`
	ExtraMeta   map[string]any `yaml:"extra_meta,omitempty"`
}
