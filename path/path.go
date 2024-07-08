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
	backCount := 0
	tailIsFile := false

	for eIdx > sIdx {
		if p[eIdx-1] == '.' {
			tailIsFile = true
		}
		if tailIsFile && p[eIdx-1] == '/' {
			return p[sIdx : eIdx-1]
		}
		backCount++
		eIdx--
	}

	if tailIsFile && eIdx == sIdx {
		return ""
	}

	return p[sIdx : eIdx+backCount]
}

func Base(p string) string {
	// Return empty string if p is empty.
	// FIX: if empty string is passed return empty
	if len(p) == 0 {
		return ""
	}

	sIdx, eIdx := CleanPath(p)

	tailIsFile := false
	backCount := 0
	for eIdx > sIdx {
		if p[eIdx-1] == '.' {
			tailIsFile = true
		}
		if tailIsFile && p[eIdx-1] == '/' {
			return p[eIdx : eIdx+backCount]
		}
		eIdx--
		backCount++
	}
	if tailIsFile && eIdx == sIdx {
		return p[sIdx : eIdx+backCount]
	}
	return ""
}
