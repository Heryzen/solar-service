package models

// Config struct holding all sub details
type Config struct {
	Cache 		CacheConfig
	Service     ServiceConfig
	Database    DatabaseConfig
	Migration   MigrationConfig
}

// ServiceConfig ...
type ServiceConfig struct {
	Name  string
	Port  int
}

// MigrationConfig ...
type MigrationConfig struct {
	Run      bool
	Version  int
}

// DatabaseConfig ...
type DatabaseConfig struct {
	Host     string
	Port     int
	DbName   string
	Username string
	Password string
	SSLMode  string
}

type CacheConfig struct {
	Address  string
	Port     int    
	Driver   string
	Username string 
	Password string
}