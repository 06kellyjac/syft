package linux

import (
	"strings"

	"github.com/anchore/packageurl-go"
)

// Release represents Linux Distribution release information as specified from https://www.freedesktop.org/software/systemd/man/os-release.html
type Release struct {
	PrettyName       string   `cyclonedx:"prettyName"` // A pretty operating system name in a format suitable for presentation to the user.
	Name             string   // identifies the operating system, without a version component, and suitable for presentation to the user.
	ID               string   `cyclonedx:"id"`     // identifies the operating system, excluding any version information and suitable for processing by scripts or usage in generated filenames.
	IDLike           []string `cyclonedx:"idLike"` // list of operating system identifiers in the same syntax as the ID= setting. It should list identifiers of operating systems that are closely related to the local operating system in regards to packaging and programming interfaces.
	Version          string   // identifies the operating system version, excluding any OS name information, possibly including a release code name, and suitable for presentation to the user.
	VersionID        string   `cyclonedx:"versionID"` // identifies the operating system version, excluding any OS name information or release code name, and suitable for processing by scripts or usage in generated filenames.
	Variant          string   `cyclonedx:"variant"`   // identifies a specific variant or edition of the operating system suitable for presentation to the user.
	VariantID        string   `cyclonedx:"variantID"` // identifies a specific variant or edition of the operating system. This may be interpreted by other packages in order to determine a divergent default configuration.
	HomeURL          string
	SupportURL       string
	BugReportURL     string
	PrivacyPolicyURL string
	CPEName          string // A CPE name for the operating system, in URI binding syntax
}

func (r *Release) String() string {
	if r == nil {
		return "unknown"
	}
	if r.PrettyName != "" {
		return r.PrettyName
	}
	if r.Name != "" {
		return r.Name
	}
	if r.Version != "" {
		return r.ID + " " + r.Version
	}

	return r.ID + " " + r.VersionID
}

func NewFromPURLDistro(purl string) *Release {
	// if package has distro information, use it
	p, err := packageurl.FromString(purl)
	if err == nil {
		for _, q := range p.Qualifiers {
			if q.Key == "distro" {
				split := strings.Split(q.Value, "-")
				if len(split) > 1 {
					name := split[0]
					version := split[1]
					return &Release{
						PrettyName: name,
						Name:       name,
						ID:         name,
						IDLike:     []string{name},
						Version:    version,
						VersionID:  version,
					}
				}
			}
		}
	}
	return nil
}
