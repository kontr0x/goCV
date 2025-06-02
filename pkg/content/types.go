package content

type TemplateData struct {
	Style   StyleTemplateData
	Content ContentTemplateData
}

type StyleTemplateData struct {
	Version string // Language code for babel
}

// Root structure
type Resume struct {
	Static   StaticInfo                  `yaml:"static"`
	Versions map[string][]VersionSection `yaml:"versions"`
}

// Static personal information
type StaticInfo struct {
	Name      string `yaml:"name"`
	Address   string `yaml:"address"`
	Phone     string `yaml:"phone"`
	Email     string `yaml:"email"`
	Birthdate string `yaml:"birthdate"`
}

// Each section in a language version (e.g., Education, Bildung)
type VersionSection struct {
	Heading string  `yaml:"heading"`
	Entries []Entry `yaml:"entries"`
}

// Each entry under a section
type Entry struct {
	Name        string `yaml:"name"`
	StartDate   string `yaml:"startdate"`
	EndDate     string `yaml:"enddate"`
	Description string `yaml:"desc"`
}

type ContentTemplateData struct {
	Version       string // Language code for babel
	StaticContent StaticInfo
	Content       []VersionSection
}
