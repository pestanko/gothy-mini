package oauth2

type Flows struct {
	ROPC Flow
}

func (f *Flows) GetFlow(grantType string) Flow {
	switch grantType {
	case "password":
		return f.ROPC
	}
	return nil
}
