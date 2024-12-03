package aoc

import (
	"aoc_2024/utils"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
)

const (
	aocUrl  = "https://adventofcode.com"
	aocYear = 2024
)

var ErrFailedAPICall = errors.New("Something went wrong with the AOC API call")

type Aoc struct {
	aocURL  string
	aocYear int
}

func AocApi() *Aoc {
	return &Aoc{aocURL: aocUrl, aocYear: aocYear}
}

func NewAoc(url string, year int) *Aoc {
	return &Aoc{aocURL: url, aocYear: year}
}

func (aoc *Aoc) InputUrl(day int) string {
	return fmt.Sprintf("%s/%d/day/%d/input", aoc.aocURL, aoc.aocYear, day)
}

func (aoc *Aoc) validateParameters(day int, cookie string) error {
	if day < 1 || day > 25 {
		return fmt.Errorf("AOC is only between day 1 and 25, got %d: %w", day, ErrFailedAPICall)
	}

	if cookie == "" {
		return fmt.Errorf("Cookie is not given, no log in possible: %w", ErrFailedAPICall)
	}

	return nil
}

func (aoc *Aoc) GetInput(day int, cookie string) ([]byte, error) {
	if err := aoc.validateParameters(day, cookie); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", aoc.InputUrl(day), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v: %w", err, ErrFailedAPICall)
	}

	sessionCookie := &http.Cookie{
		Name:  "session",
		Value: cookie,
	}
	req.AddCookie(sessionCookie)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v: %w", err, ErrFailedAPICall)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected status code %d, got %d: %w", http.StatusOK, res.StatusCode, ErrFailedAPICall)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading request: %v: %w", err, ErrFailedAPICall)
	}

	return body, err
}

func (aoc *Aoc) SaveInput(day int, cookie string) (bool, error) {
	starFolder := utils.StarFolder(day)
	inputBody, err := aoc.GetInput(day, cookie)
	if err != nil {
		return false, err
	}

	writeErr := utils.WriteInput(filepath.Join(starFolder, "input.txt"), inputBody)
	if writeErr != nil {
		return false, fmt.Errorf("failed writing input.txt: %v: %w", writeErr, ErrFailedAPICall)
	}
	writeErr = utils.WriteInput(filepath.Join(starFolder, "sample-input.txt"), []byte(""))
	if writeErr != nil {
		return false, fmt.Errorf("failed writing input-sample.txt: %v: %w", writeErr, ErrFailedAPICall)
	}

	return true, nil
}

func (aoc *Aoc) SetUpDay(day int, cookie string) (bool, error) {
	_, err := utils.CopyTemplate(day)
	if err != nil {
		return false, fmt.Errorf("failed copying template: %v: %w", err, ErrFailedAPICall)
	}

	_, err = aoc.SaveInput(day, cookie)
	if err != nil {
		return false, fmt.Errorf("failed writing input.txt: %v: %w", err, ErrFailedAPICall)
	}

	return true, nil
}
