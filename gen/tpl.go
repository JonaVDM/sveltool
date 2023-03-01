package gen

func picoCssLayoutTemplate() []byte {
	return []byte(`<script>
	import '@picocss/pico';
</script>

<slot />`)
}

func pbDockerComposeTemplate() []byte {
	return []byte(`version: '3.8'

services:
  mailhog:
    image: mailhog/mailhog:latest
    ports:
      - 1025:1025
      - 8025:8025

  pocketbase:
    image: adrianmusante/pocketbase
    ports:
      - 8090:8090
    volumes:
      - ./pb_data:/pocketbase/data
      - ./pb_migrations:/pocketbase/migrations`)
}
