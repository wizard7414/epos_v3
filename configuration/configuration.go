package configuration

type DbConfig struct {
	// Path to database
	DbPath string
	// Save result to database
	SaveToDb bool
}

type ParseConfig struct {
	// Path to parse URl list
	ParseList string
	// URL to parse list of elements
	ParseListUrl string
	// Limit to parsed elements (0 - not limited)
	ParseLimit int
	// Target parsing path
	ParseTarget string
	// Result success list and error list URLs path
	ResultListTarget string
	// Create sub folders in target for separated parsed elements
	SubFolders bool
}

type SecureConfig struct {
	// Enable parsing with proxy connection
	UseProxy bool
	// Auth HTTP header
	AuthHeader string
}

type BaseConfig struct {
	Parse ParseConfig
	Db    DbConfig
	// Log path
	LogPath string
	// Parallel processing pool size
	ProcessPoolSize int
	// Save result to Zim notes
	SaveToZim bool
}

type SourceConfig struct {
	Secure SecureConfig
}
