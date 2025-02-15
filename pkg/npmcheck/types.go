package npmcheck

type PackageJSON struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

type VersionInfo struct {
	Current string
	Latest  string
}
