// Package path provides utilities to clean and manage file paths.
package path

// PathPrefix is a global variable used to store the prefix of the path, such as "~/" if the path starts with "~".
var PathPrefix = ""

// CleanPath removes trailing slashes from the end of a path string,
// as well as any leading slashes, dots, or tilde (~) characters from the beginning of the path.
// It sets the PathPrefix variable to "~" or "~/" if the path starts with "~" or "~/".
// Returns the start and end indices (sIdx, eIdx) of the cleaned path within the input string.
func CleanPath(p string) (sIdx, eIdx int) {
	eIdx = len(p)
	sIdx = 0

	// Remove trailing slashes from the end of the path.
	for p[eIdx-1] == '/' {
		eIdx--
	}

	// Reset PathPrefix.
	PathPrefix = ""

	// Remove leading slashes, dots, or tilde characters.
	for p[sIdx] == '/' || p[sIdx] == '.' || p[sIdx] == '~' {
		if p[sIdx] == '~' && len(PathPrefix) == 0 {
			PathPrefix += "~"
		}
		if p[sIdx] == '/' && len(PathPrefix) == 1 {
			PathPrefix += "/"
		}
		sIdx++
		if sIdx == len(p)-1 {
			break
		}
	}

	return sIdx, eIdx
}

// Dir extracts and returns the directory part of the path `p`,
// excluding any base file component.
// If the path ends with a file name, it returns the directory portion with PathPrefix applied.
// If the path doesn't include a file name, it returns an empty string.
func Dir(p string) string {
	if len(p) == 0 {
		return ""
	}
	sIdx, eIdx := CleanPath(p)
	tailIsFile := false

	// Traverse the cleaned path in reverse to locate the last directory separator.
	for i := eIdx - 1; i >= sIdx; i-- {
		if p[i] == '.' {
			tailIsFile = true
		}
		if tailIsFile && p[i] == '/' {
			return PathPrefix + p[sIdx:i]
		}
	}

	// If no directory separator is found in a file path, return empty string if it’s a file.
	if tailIsFile {
		return ""
	}

	return PathPrefix + p[sIdx:eIdx]
}

// Base extracts and returns the base file component of the path `p`.
// If the path does not end with a valid file name, it returns an empty string.
// PathPrefix is prepended to the result if it exists.
func Base(p string) string {
	if len(p) == 0 {
		return ""
	}
	sIdx, eIdx := CleanPath(p)
	tailIsFile := false

	// Traverse the cleaned path in reverse to locate the start of the file name.
	for i := eIdx - 1; i >= sIdx; i-- {
		if p[i] == '.' {
			tailIsFile = true
		}
		if p[i] == '/' {
			if tailIsFile {
				return p[i+1 : eIdx]
			}
			return ""
		}
	}

	// Return the entire cleaned path if it represents a file.
	if tailIsFile {
		return PathPrefix + p[sIdx:eIdx]
	}

	return ""
}
