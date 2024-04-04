package autoload

/*
	You can just read the .env file on import just by doing

		import _ "github.com/harness/godotenv/v2/autoload"

	And bob's your mother's brother
*/

import "github.com/harness/godotenv/v2"

func init() {
	godotenv.Load()
}
