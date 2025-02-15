# NPM Package Version Checker

A Go-based tool to check and compare current vs latest versions of npm packages in your project.

## Features

- Reads package.json file
- Checks both dependencies and devDependencies
- Fetches latest versions from npm registry
- Displays current and latest versions side by side

## Installation

```bash
git clone https://github.com/yourusername/npmcheck.git
cd npmcheck
go mod init npmcheck
```

## Usage

Place your package.json file in the root directory
Run the checker:

```bash
go run cmd/npmcheck/main.go
```

## Sample Output


```bash
Package: @react-google-maps/api
Current Version: 2.2.0
Latest Version: 2.20.6
-------------------

Package: google-maps-react
Current Version: 2.0.6
Latest Version: 2.0.6
-------------------

Package: google-map-react
Current Version: 2.1.9
Latest Version: 2.2.1
-------------------

Package: react-google-maps
Current Version: 9.4.5
Latest Version: 9.4.5
-------------------

Package: surge
Current Version: 0.23.0
Latest Version: 0.24.6
-------------------

Package: react-map-gl
Current Version: 6.1.15
Latest Version: 8.0.1
-------------------

Package: @fortawesome/fontawesome-free
Current Version: 5.15.3
Latest Version: 6.7.2
-------------------

Package: @chainlink/contracts
Current Version: 0.1.7
Latest Version: 1.3.0
-------------------

Package: bootstrap-icons
Current Version: 1.5.0
Latest Version: 1.11.3
-------------------
```
