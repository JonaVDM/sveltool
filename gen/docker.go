package gen

func PocketbaseDocker() error {
	return simpleTemplate("docker/pocketbase-docker-compose.yml.stub", "docker-compose.yml")
}

func DockerIgnore() error {
	return simpleTemplate("docker/dockerignore", ".dockerignore")
}

func BasicDockerFile() error {
	return simpleTemplate("docker/Dockerfile", "Dockerfile")
}
