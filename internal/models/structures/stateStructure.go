package structures

const (
	StateStart State = iota
	StateDefault

	StateAskForMe
	UserNotExists
	StateStartMenu
	StateAskForFriend

	StateSetNotifications
	StateSetCustomNotification

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
