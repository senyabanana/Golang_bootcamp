package repository

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ClientRepository interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) error
	JoinCluster(address, port string) error
}

type HTTPRepository struct {
	currentNode string
}

func NewHTTPRepository() *HTTPRepository {
	return &HTTPRepository{currentNode: "127.0.0.1:8765"}
}

func (r *HTTPRepository) Get(key string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/get?key=%s", r.currentNode, key))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: %s", string(body))
	}

	return string(body), nil
}

func (r *HTTPRepository) Set(key, value string) error {
	url := fmt.Sprintf("http://%s/set?key=%s", r.currentNode, key)
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(value))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error setting value: %s", string(body))
	}

	return nil
}

func (r *HTTPRepository) Delete(key string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s/delete?key=%s", r.currentNode, key), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error deleting value: %s", string(body))
	}

	return nil
}

func (r *HTTPRepository) JoinCluster(address, port string) error {
	resp, err := http.Get(fmt.Sprintf("http://%s/join?address=%s&port=%s", r.currentNode, address, port))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error joining cluster: %s", string(body))
	}

	return nil
}
