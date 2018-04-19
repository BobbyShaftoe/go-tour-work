package types

type Resources struct {
	Path     string
	Resource string
	Consul   string
	Unused   string
}

func (resource *Resources) FromCSV(ss []string) {
	if nil == resource {
		return
	}

	if nil == ss {
		return
	}

	resource.Path = ss[0]
	resource.Resource = ss[1]
	resource.Consul = ss[2]
	resource.Unused = ss[3]
}

