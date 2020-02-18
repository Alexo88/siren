package lib

// StatusKind represents a status of a model
type StatusKind int

// Model statuses
const (
	StatusUnknown StatusKind = iota
	StatusOffline
	StatusOnline
	StatusNotFound
	StatusDenied
	StatusExists
)

func (s StatusKind) String() string {
	switch s {
	case StatusOffline:
		return "offline"
	case StatusOnline:
		return "online"
	case StatusNotFound:
		return "not found"
	case StatusDenied:
		return "denied"
	case StatusExists:
		return "exists"
	}
	return "unknown"
}
