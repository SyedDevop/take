package path

func CleanPath(p string) (sIdx, eIdx int) {
	eIdx = len(p)
	sIdx = 0
	// Remove Suffix of "/"
	for p[eIdx-1] == '/' {
		eIdx--
	}
	// Remove Prefix of "/" or "."
	for p[sIdx] == '/' || p[sIdx] == '.' {
		sIdx++
	}
	return sIdx, eIdx
}

func Dir(p string) string {
	// Return empty string if p is empty.
	// FIX: if empty string is passed return empty
	if len(p) == 0 {
		return ""
	}

	sIdx, eIdx := CleanPath(p)
	tailIsFile := false

	for i := eIdx - 1; i >= sIdx; i-- {
		if p[i] == '.' {
			tailIsFile = true
		}
		if tailIsFile && p[i] == '/' {
			return p[sIdx:i]
		}
	}

	if tailIsFile {
		return ""
	}

	return p[sIdx:eIdx]
}

func Base(p string) string {
	// Return empty string if p is empty.
	// FIX: if empty string is passed return empty
	if len(p) == 0 {
		return ""
	}

	sIdx, eIdx := CleanPath(p)

	tailIsFile := false
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
	if tailIsFile {
		return p[sIdx:eIdx]
	}
	return ""
}
