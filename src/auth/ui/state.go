package ui

type LoginHintState struct {
	Password string
	ShowHint bool
}

func NewLoginHintState() LoginHintState {
	return LoginHintState{ShowHint: true}
}
