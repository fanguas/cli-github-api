package layout

import (
	"fmt"
	"os"
	"strings"

	"github.com/the-palace-company/cli-github-api-go/github"
	"github.com/the-palace-company/cli-github-api-go/module"
)

func Menu() {
	println("Bienvenido al CLI de GitHub API en Go")
	println("1. Listar miembros de la organización")
	println("2. Listar repositorios de la organización")
	println("3. Otorgar permiso a colaborador")
	println("4. Salir\n")
}

func SeleccionaOpcion(opcion int) {
	org := "the-palace-company"
	token := module.ValidateTokenGithub()

	switch opcion {
	case 1:
		println("Listando miembros de la organización...")
		miembros, err := github.ObtenerMiembrosOrganizacion(token, org)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("\nMiembros de la organización %s:\n\n", org)
		for i, miembro := range miembros {
			fmt.Printf("👨🏻‍💻 %d %s id: %d perfil: %s\n", i+1, miembro.Alias, miembro.ID, miembro.URLPerfil)
		}

	case 2:
		println("Listando repositorios de la organización...")
		repos, err := github.ObtenerRepositoriosOrganizacion(token, org)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("\nRepositorios de la organización %s:\n\n", org)
		for i, repo := range repos {
			fmt.Printf("📦 %d %s\n   id: %d\n   url: %s\n   Rama por defecto: %s\n\n", i+1, repo.Name, repo.ID, repo.HtmlURL, repo.DBranch)
		}
	case 3:
		println("Otorgando permiso a un miembro...")
		var miembro, permiso string
		var repositorios []string

		fmt.Print("Introduce el nombre del colaborador: ")
		fmt.Scanln(&miembro)

		fmt.Print("Introduce el tipo de permiso (pull, triage, push, maintain, admin): ")
		fmt.Scanln(&permiso)

		fmt.Print("Introduce los repositorios (separados por comas): ")
		var repos string
		fmt.Scanln(&repos)
		repositorios = strings.Split(repos, ",")

		err := github.OtorgaAccesoAMiembro(token, org, repositorios, miembro, permiso)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Permiso otorgado con éxito.")
	case 4:
		println("Saliendo del programa...")
		os.Exit(0)
	default:
		println("Opción no válida. Por favor, selecciona una opción del menú.")
	}
}
