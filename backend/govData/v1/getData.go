package v1

import (
	"encoding/json"
	"net/http"

	"github.com/ethirajmudhaliar/backend/react-go-gov-search-data/common"
	"github.com/ethirajmudhaliar/backend/react-go-gov-search-data/logger"
)

// GetGovernmentData handles the API request to the Census Bureau API.
func GetGovernmentData(w http.ResponseWriter, r *http.Request) {
	// URL to fetch all states data
	apiURL := "https://api.census.gov/data/2020/acs/acs5?get=NAME,B01001_001E&for=state:*"

	// Make a GET request to the Census Bureau API
	resp, err := http.Get(apiURL)
	if err != nil {
		logger.Error("Failed to fetch government data: " + err.Error())
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch government data")
		return
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var rawData [][]string
	if err := json.NewDecoder(resp.Body).Decode(&rawData); err != nil {
		logger.Error("Failed to decode JSON: " + err.Error())
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to process government data")
		return
	}

	// Ensure the response contains data
	if len(rawData) < 2 {
		logger.Error("No data found in the API response")
		common.RespondWithError(w, http.StatusNoContent, "No data available")
		return
	}

	// Transform raw data into a structured format
	var aggregatedData []map[string]string
	for _, row := range rawData[1:] { // Skip header row
		if len(row) < 3 {
			continue
		}
		transformed := map[string]string{
			"state":      row[0],
			"population": row[1],
			"state_fips": row[2],
		}
		aggregatedData = append(aggregatedData, transformed)
	}

	// Respond with the aggregated data
	common.RespondWithSuccess(w, http.StatusOK, aggregatedData, "Government data retrieved successfully")
}
