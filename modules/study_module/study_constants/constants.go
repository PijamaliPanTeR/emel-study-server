package study_constants

import "fmt"

const ModuleID = "study_module"

// DefaultSoundIDs returns 22 sound IDs (s1..s22).
func DefaultSoundIDs() []string {
	ids := make([]string, 22)
	for i := 0; i < 22; i++ {
		ids[i] = fmt.Sprintf("s%d", i+1)
	}
	return ids
}
