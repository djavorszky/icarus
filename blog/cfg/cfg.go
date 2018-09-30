package cfg

// ServiceOpts contains all the options that are needed (or optional) to set up the service.
type ServiceOpts struct {
	Database DatabaseOpts
}

// DatabaseOpts contains options needed to connect to the database.
// If SkipAuth is set to true, then authentication will be skipped.
type DatabaseOpts struct {
	SkipAuth     bool
	DatabaseName string
	Address      string
	User         string
	Pass         string
}
