package pocketbase

import (
	"fmt"

	"github.com/jonavdm/sveltool/utils"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var crudCmd = &cobra.Command{
	Use:   "crud [name]",
	Short: "Creates a basic crud for a pocketbase collection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		upper := cases.Title(language.Georgian).String(name)
		path := "src/routes/" + name

		if p, _ := cmd.Flags().GetString("path"); p != "" {
			path = "src/routes/" + p
		}

		if err := utils.CreateFolder(path); err != nil {
			fmt.Println(err)
			return
		}
		if err := utils.CreateFolder(path + "/new"); err != nil {
			fmt.Println(err)
			return
		}
		if err := utils.CreateFolder(path + "/[id]/edit"); err != nil {
			fmt.Println(err)
			return
		}

		utils.SimpleTemplate("pocketbase_crud/[id]/edit/+page.server.ts", path+"/[id]/edit/+page.server.ts")
		utils.SimpleTemplate("pocketbase_crud/[id]/edit/+page.svelte", path+"/[id]/edit/+page.svelte")

		utils.SimpleTemplate("pocketbase_crud/[id]/+page.server.ts", path+"/[id]/+page.server.ts")
		utils.SimpleTemplate("pocketbase_crud/[id]/+page.svelte", path+"/[id]/+page.svelte")
		utils.SimpleTemplate("pocketbase_crud/[id]/+page.ts", path+"/[id]/+page.ts")

		utils.SimpleTemplate("pocketbase_crud/new/+page.server.ts", path+"/new/+page.server.ts")
		utils.SimpleTemplate("pocketbase_crud/new/+page.svelte", path+"/new/+page.svelte")

		utils.SimpleTemplate("pocketbase_crud/+layout.svelte", path+"/+layout.svelte")
		utils.SimpleTemplate("pocketbase_crud/+page.svelte", path+"/+page.svelte")
		utils.SimpleTemplate("pocketbase_crud/+page.ts", path+"/+page.ts")
		utils.SimpleTemplate("pocketbase_crud/Form.svelte", path+"/"+upper+"Form.svelte")
		utils.SimpleTemplate("pocketbase_crud/types.ts", path+"/"+name+"Types.ts")
	},
}

func init() {
	PocketbaseCmd.AddCommand(crudCmd)
	crudCmd.Flags().String("path", "", "The root path for where to place the crud, defaults to [name]")
}
