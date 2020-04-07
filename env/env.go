package env

import "os"

//New start function for configure all environment constants
func New() {
	os.Setenv("BASIC_USERNAME", "test")
	os.Setenv("BASIC_PASSWORD", "test")
	os.Setenv("SQLITE_BASE", "test.db")
}
