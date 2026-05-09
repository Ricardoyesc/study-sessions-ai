package handlers

type Handlers struct {
	A2UI      *A2UIHandler
	ColdStart *ColdStartHandler
	Session   *SessionHandler
	Capsule   *CapsuleHandler
	Quiz      *QuizHandler
	Socratic  *SocraticHandler
	User      *UserHandler
}
