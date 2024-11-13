package structures


const (
	StateStart State = iota
	StateDefault
	StateAskFilial
	StateAskFaculty
	StateAskCathedra
	StateAskCourse
	StateAskGroup
	StateConfirm
	StateSearch
	StateGroupNotFound
)