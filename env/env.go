// pachage env - Global env configurations
package env

import "os"

//New start function for configure all environment constants
func New() {
	os.Setenv("BASIC_USERNAME", "test")
	os.Setenv("BASIC_PASSWORD", "test")

	// MySQL configurations
	os.Setenv("DB_USER", "authenticated")
	os.Setenv("DB_PASS", "123456")
	os.Setenv("DB_BASE", "authenticated")
}
