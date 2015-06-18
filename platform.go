package omni

type Version struct {
	Label     string
	Installed bool
}

type Platform struct {
	Label    string
	Versions []Version
}

func GetPlatforms() []Platform {
	return []Platform{
		{
			Label:    "puppet",
			Versions: []Version{},
		},
		{
			Label:    "ansible",
			Versions: []Version{},
		},
	}
}
