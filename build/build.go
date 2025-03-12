package build

import "time"

const version = "v1.0.1"
const buildDate = "Wed, 12 Mar 2025 02:59:29 +0000"

type build struct {
	Version string
	Date    time.Time
	Debug   bool
}

func instance() build {
	debugValue := false
	versionValue := version
	buildDateValue := time.Now()

	if version == "${version}" {
		versionValue = "snapshot"
		debugValue = true
	}
	if buildDate != "${buildDate}" {
		v, err := time.Parse(time.RFC1123Z, buildDate)
		if err != nil {
			buildDateValue = v
		}
	}
	return build{
		Date:    buildDateValue,
		Version: versionValue,
		Debug:   debugValue,
	}
}

var buildInstnace = instance()

func Build() build {
	return buildInstnace
}
