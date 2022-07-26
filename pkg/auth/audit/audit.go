package audit

// Auditor main interface
type Auditor interface {
	WriteAudit(entry Entry) error
}

// Type audit type enum
type Type string

const (
	TypeSuccess Type = "success"
	TypeFail    Type = "fail"
)

// Entry single audit entry
type Entry struct {
	LoginMethod string `json:"loginMethod"`
	Detail      string `json:"detail"`
	Username    string `json:"username"`
	ClientId    string `json:"clientId"`
}
