package gen

import "github.com/jonavdm/sveltool/utils"

func PocketBase() error {
	if err := utils.CreateFolder("src/lib"); err != nil {
		return err
	}
	return simpleTemplate("pocketbase/pocketbase.ts", "src/lib/pocketbase.ts")
}

func ClientHook() error {
	return simpleTemplate("pocketbase/hooks.client.ts", "src/hooks.client.ts")
}

func ServerHook() error {
	return simpleTemplate("pocketbase/hooks.server.ts", "src/hooks.server.ts")
}

func AuthLayout() error {
	if err := utils.CreateFolder("src/routes/(auth)"); err != nil {
		return err
	}
	return simpleTemplate("pocketbase/+layout.server.ts", "src/routes/(auth)/+layout.server.ts")
}

func LoginPage() error {
	if err := utils.CreateFolder("src/routes/(auth)/login"); err != nil {
		return err
	}
	return simpleTemplate("pocketbase/login-page.svelte", "src/routes/(auth)/login/+page.svelte")
}

func LoginServer() error {
	if err := utils.CreateFolder("src/routes/(auth)/login"); err != nil {
		return err
	}
	return simpleTemplate("pocketbase/login-page.server.ts", "src/routes/(auth)/login/+page.server.ts")
}

func SignupPage() error {
	if err := utils.CreateFolder("src/routes/(auth)/signup"); err != nil {
		return err
	}
	return simpleTemplate("pocketbase/signup-page.svelte", "src/routes/(auth)/signup/+page.svelte")
}

func SignupServer() error {
	if err := utils.CreateFolder("src/routes/(auth)/signup"); err != nil {
		return err
	}
	return simpleTemplate("pocketbase/signup-page.server.ts", "src/routes/(auth)/signup/+page.server.ts")
}

func LogoutServer() error {
	if err := utils.CreateFolder("src/routes/(auth)/logout"); err != nil {
		return err
	}
	return simpleTemplate("pocketbase/logout-server.ts", "src/routes/(auth)/logout/+server.ts")
}
