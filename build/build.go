package build

import "time"

const version = "${version}"
const buildDate = "${buildDate}"

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
