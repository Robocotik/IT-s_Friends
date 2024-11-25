package structures

const (
	StateStart State = iota
	StateDefault

	StateAskForMe
	StateStartMenu
	StateAskForFriend

	StateSearch
	StateGroupFound
	StateAskNickname
	StateShowTimetable
	StateRedirectToStartSearch
)

const (
	StateAskFilial State = iota
	StateAskFaculty
	StateAskCathedra
	StateAskCourse
	StateAskGroup
	StateConfirm
)
