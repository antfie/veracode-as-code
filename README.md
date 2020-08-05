[![Go Report Card](https://goreportcard.com/badge/github.com/antfie/veracode-as-code)](https://goreportcard.com/report/github.com/antfie/veracode-as-code) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/antfie/veracode-go-api/blob/master/LICENSE)

# Veracode As Code


## Building


## Running


## Configurtion

### Authentication

You can either use a Veracode API Credentials File (https://help.veracode.com/reader/LMv_dtSHyb7iIxAQznC~9w/zm4hbaPkrXi02YmacwH3wQ) or environment variables (`VERACODE_API_KEY_ID` and `VERACODE_API_KEY_SECRET`) to configure your API authentication.

### YAML Example

scan.yml:

```yaml
application:
    name: Lunar Lander

    # Optional.
    # This is only used if the application does not exist.
    # If no policy is defined the Veracode "Recommended Very High + SCA" policy will be applied.
    policy: STRICT

# Optional.
# If a sandbox is not specified the scan will be performed at the policy level
# If a sandbox is specified and it does not exist, a new sandbox will be created
sandbox: Release Candidate

name: 2020-06-16-1

# Optional.
# If not specified the scan will start after the prescan
autoScanAfterPrescan: true

uploads:
    - ${artifact_app.war}
    - ${artifact_front_end.js}
```

### JSON Example

scan.json:

```json
{
    "application": {
        "name": "Lunar Lander",
        "policy": "STRICT"
    },
    "sandbox": "Release Candidate",
    "name": "2020-06-16-1",
    "autoScanAfterPrescan": true,
    "uploads": [
        "${artifact_app.war}",
        "${artifact_front_end.js}"
    ]
}
```
