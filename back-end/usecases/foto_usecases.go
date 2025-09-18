package usecases

import (
	"back-end/models"
	"back-end/repository"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type FotoUseCases struct {
	repository repository.FotoRepository
}

func NewFotoUseCases(repo repository.FotoRepository) FotoUseCases {
	return FotoUseCases{
		repository: repo,
	}
}

func (usecase FotoUseCases) GetAllFotos() ([]models.Foto, error) {
	return usecase.repository.GetAllFotos()
}

func (usecase FotoUseCases) GetFotoById(idFoto int) (models.Foto, error) {
	return usecase.repository.GetFotoById(idFoto)
}

func (usecase FotoUseCases) CreateFoto(foto models.Foto) (int, error) {
	return usecase.repository.CreateFoto(foto)
}

func (usecase FotoUseCases) UpdateFoto(foto models.Foto) (models.Foto, error) {
	return usecase.repository.UpdateFoto(foto)
}

func (usecase FotoUseCases) DeleteFoto(idFoto int) error {
	return usecase.repository.DeleteFoto(idFoto)
}

func (usecase FotoUseCases) PostFoto(foto models.Foto) (int, error) {
	return usecase.repository.PostFoto(foto)
}

func (usecase FotoUseCases) GetFotosByTelefone(telefone string) ([]models.Foto, error) {
	return usecase.repository.GetFotosByTelefone(telefone)
}

func (usecase FotoUseCases) EnviarFotosClienteParaVendedor(telefoneCliente string) error {
	fotos, err := usecase.GetFotosByTelefone(telefoneCliente)
	if err != nil {
		return err
	}
	for _, foto := range fotos {
		err = EnviarImagem(foto)
		if err != nil {
			return err
		}
	}
	return nil
}

func EnviarImagem(foto models.Foto) error {
	baseURL := "http://147.93.10.167:5678/webhook/808f0e3b-8449-4d7b-8e33-bed687e3a64cFigueiredo"
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	data := url.Values{}
	data.Set("number", "5515998223027")
	data.Set("media", foto.LinkFoto)
	data.Set("nomeInstancia", "Figueiredo_Imoveis")

	req, err := http.NewRequest("POST", baseURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erro ao enviar requisição: %s", resp.Status)
	}
	return nil
}
