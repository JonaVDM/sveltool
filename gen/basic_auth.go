package gen

import "github.com/jonavdm/sveltool/utils"

func BasicAuthLayout() error {
	if err := utils.CreateFolder("src/routes/(auth)"); err != nil {
		return err
	}

	return simpleTemplate("basic_auth/+layout.server.ts", "src/routes/(auth)/+layout.server.ts")
}

func BasicAuthLoginServer() error {
	if err := utils.CreateFolder("src/routes/(auth)/login"); err != nil {
		return err
	}

	return simpleTemplate("basic_auth/login+page.ts", "src/routes/(auth)/login/+page.server.ts")
}

func BasicAuthLogin() error {
	if err := utils.CreateFolder("src/routes/(auth)/login"); err != nil {
		return err
	}

	return simpleTemplate("basic_auth/login+page.svelte", "src/routes/(auth)/login/+page.svelte")
}

func BasicAuthLogout() error {
	if err := utils.CreateFolder("src/routes/(auth)/logout"); err != nil {
		return err
	}

	return simpleTemplate("basic_auth/logout+server.ts", "src/routes/(auth)/logout/+server.ts")
}

func BasicAuthClientHook() error {
	return simpleTemplate("basic_auth/hooks.client.ts", "src/hooks.client.ts")
}

func BasicAuthServerHook() error {
	return simpleTemplate("basic_auth/hooks.server.ts", "src/hooks.server.ts")
}
