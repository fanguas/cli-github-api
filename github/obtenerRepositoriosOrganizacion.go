package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Repositorio struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	DBranch  string `json:"default_branch"`
	Language string `json:"language"`
	HtmlURL  string `json:"html_url"`
}

func ObtenerRepositoriosOrganizacion(token, org string) ([]Repositorio, error) {
	if token == "" || org == "" {
		return nil, fmt.Errorf("el token y el nombre de la organización son obligatorios")
	}

	var todosLosRepos []Repositorio
	cliente := &http.Client{}
	pagina := 1
	elementosPorPag := 30

	for {
		urlPaginada := fmt.Sprintf(GitHubAPIBaseURL+"/orgs/%s/repos?page=%d&per_page=%d", org, pagina, elementosPorPag)

		solicitud, err := http.NewRequest(http.MethodGet, urlPaginada, nil)
		if err != nil {
			return nil, fmt.Errorf("error al crear la solicitud: %v", err)
		}

		solicitud.Header.Set("Accept", AcceptHeader)
		solicitud.Header.Set(AuthHeader, "Bearer "+token)

		respuesta, err := cliente.Do(solicitud)
		if err != nil {
			return nil, fmt.Errorf("error al realizar la solicitud: %v", err)
		}
		defer respuesta.Body.Close()

		if respuesta.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error al obtener los repositorios: código %d", respuesta.StatusCode)
		}

		cuerpo, err := io.ReadAll(respuesta.Body)
		if err != nil {
			return nil, fmt.Errorf("error al leer la respuesta: %v", err)
		}

		var repos []Repositorio
		err = json.Unmarshal(cuerpo, &repos)
		if err != nil {
			return nil, fmt.Errorf("error al deserializar la respuesta: %v", err)
		}

		todosLosRepos = append(todosLosRepos, repos...)

		if len(repos) == 0 || len(repos) < elementosPorPag {
			break
		}
		pagina++
	}

	return todosLosRepos, nil
}
