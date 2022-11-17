package blob

type ExistenceState int64

const (
	Existing         ExistenceState = 0
	NotExisting                     = 1
	ExistenceUnknown                = 2
)
