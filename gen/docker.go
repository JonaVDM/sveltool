package gen

func PocketbaseDocker() error {
	return simpleTemplate("docker/pocketbase-docker-compose.yml", "docker-compose.yml")
}
