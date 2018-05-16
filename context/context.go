package context

type Context struct {
	IsDryRun bool
}

func New(isDryRun bool) Context {
	return Context{
		IsDryRun: isDryRun,
	}
}
