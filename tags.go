package drone_aliyun_oss

// copy from https://github.com/drone-plugins/drone-docker/blob/72ef7b1f3fa6c47c58f33ca312bdc8c484b6c5a4/tags.go

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
	"strings"
)

// UseDefaultTag
// DRONE_COMMIT_REF=refs/heads/master
// DRONE_COMMIT_REF=refs/tags/v1.0.0
// DRONE_REPO_BRANCH=master
func UseDefaultTag(ref, defaultBranch string) bool {
	if strings.HasPrefix(ref, "refs/tags/") {
		return true
	}

	if stripHeadPrefix(ref) == defaultBranch {
		return true
	}
	return false
}

func DefaultTag(ref string) (string, error) {
	if !strings.HasPrefix(ref, "refs/tags/") {
		return "latest", nil
	}
	v := stripTagPrefix(ref)
	version, err := semver.NewVersion(v)
	if err != nil {
		return "latest", err
	}
	if version.PreRelease != "" || version.Metadata != "" {
		return version.String(), nil
	}

	v = stripTagPrefix(ref)
	v = splitOff(splitOff(v, "+"), "-")
	dotParts := strings.SplitN(v, ".", 3)

	//if version.Major == 0 {
	//	return []string{
	//		fmt.Sprintf("%0*d.%0*d", len(dotParts[0]), version.Major, len(dotParts[1]), version.Minor),
	//		fmt.Sprintf("%0*d.%0*d.%0*d", len(dotParts[0]), version.Major, len(dotParts[1]), version.Minor, len(dotParts[2]), version.Patch),
	//	}, nil
	//}
	//return []string{
	//	fmt.Sprintf("%0*d", len(dotParts[0]), version.Major),
	//	fmt.Sprintf("%0*d.%0*d", len(dotParts[0]), version.Major, len(dotParts[1]), version.Minor),
	//	fmt.Sprintf("%0*d.%0*d.%0*d", len(dotParts[0]), version.Major, len(dotParts[1]), version.Minor, len(dotParts[2]), version.Patch),
	//}, nil
	return fmt.Sprintf("%0*d.%0*d.%0*d", len(dotParts[0]), version.Major, len(dotParts[1]), version.Minor, len(dotParts[2]), version.Patch), nil
}

func stripHeadPrefix(ref string) string {
	return strings.TrimPrefix(ref, "refs/heads/")
}

func stripTagPrefix(ref string) string {
	ref = strings.TrimPrefix(ref, "refs/tags/")
	ref = strings.TrimPrefix(ref, "v")
	return ref
}

func splitOff(input string, delim string) string {
	parts := strings.SplitN(input, delim, 2)

	if len(parts) == 2 {
		return parts[0]
	}

	return input
}
