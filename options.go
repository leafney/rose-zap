package rzap

type (
	Option func(opt *options)

	options struct {
		level      string
		callerSkip int
	}
)
