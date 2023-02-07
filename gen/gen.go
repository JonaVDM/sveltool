package gen

import (
	"errors"
	"os"

	"github.com/jonavdm/sveltool/utils"
)

func gen(filename string, contents []byte) error {
	if _, err := os.Stat(filename); err == nil {
		return errors.New("file already exists")
	}

	return os.WriteFile(filename, contents, 0664)
}

func TailwindConfig() error {
	return gen("tailwind.config.cjs", tailwindConfigTemplate())
}

func PostCssConfig() error {
	return gen("postcss.config.cjs", tailwindPostcssConfigTemplate())
}

func TailwindStyles() error {
	return gen("src/app.css", tailwindStylesTemplate())
}

func TailwindLayout() error {
	return gen("src/routes/+layout.svelte", tailwindLayoutTemplate())
}

func PicoCssLayout() error {
	return gen("src/routes/+layout.svelte", picoCssLayoutTemplate())
}

func PocketBase() error {
	if err := utils.CreateFolder("src/lib"); err != nil {
		return err
	}
	return gen("src/lib/pocketbase.ts", pocketbaseLib())
}

func ClientHook() error {
	return gen("src/hooks.client.ts", clientHooksTemplate())
}

func ServerHook() error {
	return gen("src/hooks.server.ts", serverHooksTemplate())
}

func AuthLayout() error {
	if err := utils.CreateFolder("src/routes/(auth)"); err != nil {
		return err
	}
	return gen("src/routes/(auth)/+layout.server.ts", pbAuthLayoutTemplate())
}

func LoginPage() error {
	if err := utils.CreateFolder("src/routes/(auth)/login"); err != nil {
		return err
	}
	return gen("src/routes/(auth)/login/+page.svelte", pbLoginPageTemplate())
}

func LoginServer() error {
	if err := utils.CreateFolder("src/routes/(auth)/login"); err != nil {
		return err
	}
	return gen("src/routes/(auth)/login/+page.server.ts", pbLoginServerTemplate())
}

func SignupPage() error {
	if err := utils.CreateFolder("src/routes/(auth)/signup"); err != nil {
		return err
	}
	return gen("src/routes/(auth)/signup/+page.svelte", pbSignupPageTemplate())
}

func SignupServer() error {
	if err := utils.CreateFolder("src/routes/(auth)/signup"); err != nil {
		return err
	}
	return gen("src/routes/(auth)/signup/+page.server.ts", pbSignupServerTemplate())
}

func LogoutServer() error {
	if err := utils.CreateFolder("src/routes/(auth)/logout"); err != nil {
		return err
	}
	return gen("src/routes/(auth)/logout/+server.ts", pbLogoutServerTemplate())
}
