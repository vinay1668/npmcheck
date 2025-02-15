package npmcheck

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func CheckPackageVersions(filePath string) (map[string]VersionInfo, error) {
	// Read package.json
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading package.json: %v", err)
	}

	var pkg PackageJSON
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, fmt.Errorf("error parsing package.json: %v", err)
	}

	results := make(map[string]VersionInfo)

	// Check dependencies
	for name, version := range pkg.Dependencies {
		latest, err := getLatestVersion(name)
		if err != nil {
			fmt.Printf("Warning: Could not fetch latest version for %s: %v\n", name, err)
			continue
		}
		results[name] = VersionInfo{
			Current: cleanVersion(version),
			Latest:  latest,
		}
	}

	// Check devDependencies
	for name, version := range pkg.DevDependencies {
		latest, err := getLatestVersion(name)
		if err != nil {
			fmt.Printf("Warning: Could not fetch latest version for %s: %v\n", name, err)
			continue
		}
		results[name] = VersionInfo{
			Current: cleanVersion(version),
			Latest:  latest,
		}
	}

	return results, nil
}

func getLatestVersion(packageName string) (string, error) {
	url := fmt.Sprintf("https://registry.npmjs.org/%s/latest", packageName)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result struct {
		Version string `json:"version"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	return result.Version, nil
}

func cleanVersion(version string) string {
	return strings.TrimPrefix(strings.TrimPrefix(version, "^"), "~")
}
