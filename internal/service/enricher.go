package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"
)

type EnricherService struct {
	genderApi      string
	ageApi         string
	nationalizeApi string
	logger         *slog.Logger
}

func NewEnricherService(genderApi, ageApi, nationalizeApi string, logger *slog.Logger) *EnricherService {
	return &EnricherService{
		genderApi:      genderApi,
		ageApi:         ageApi,
		nationalizeApi: nationalizeApi,
		logger:         logger.With("service", "enricher"),
	}
}

func (s *EnricherService) GetGender(ctx context.Context, name string) (gender string, err error) {
	startTime := time.Now()

	logger := s.logger.With(
		"method", "GetGender",
		"name", name,
		"start_time", startTime,
	)
	queryParams := url.Values{}
	queryParams.Add("name", name)
	fullURL := fmt.Sprintf("%s?%s", s.genderApi, queryParams.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API request failed: status %d, body: %s", resp.StatusCode, string(body))
	}

	var result GenderResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("JSON decode failed: %w", err)
	}

	logger.InfoContext(ctx, "API call succeeded",
		"duration", time.Since(startTime).String(),
		"result", result.Gender,
	)

	return result.Gender, nil
}

func (s *EnricherService) GetAge(ctx context.Context, name string) (age int, err error) {

	startTime := time.Now()

	logger := s.logger.With(
		"method", "GetAge",
		"name", name,
		"start_time", startTime,
	)

	queryParams := url.Values{}
	queryParams.Add("name", name)
	fullURL := fmt.Sprintf("%s?%s", s.ageApi, queryParams.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		return 0, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("API request failed: status %d, body: %s", resp.StatusCode, string(body))
	}

	var result AgeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("JSON decode failed: %w", err)
	}

	logger.InfoContext(ctx, "API call succeeded",
		"duration", time.Since(startTime).String(),
		"result", result.Age,
	)

	return result.Age, nil
}

func (s *EnricherService) GetNationality(ctx context.Context, name string) (nationality string, err error) {

	startTime := time.Now()

	logger := s.logger.With(
		"method", "GetNationality",
		"name", name,
		"start_time", startTime,
	)
	queryParams := url.Values{}
	queryParams.Add("name", name)
	fullURL := fmt.Sprintf("%s?%s", s.nationalizeApi, queryParams.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API request failed: status %d, body: %s", resp.StatusCode, string(body))
	}

	var result NationalityResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("JSON decode failed: %w", err)
	}

	if len(result.Countries) == 0 {
		return "", fmt.Errorf("empty field: country")
	}

	logger.InfoContext(ctx, "API call succeeded",
		"duration", time.Since(startTime).String(),
		"result", result.Countries[0].CountryID,
	)

	return result.Countries[0].CountryID, nil
}
