package ir

import "github.com/ogen-go/ogen/openapi"

// SecurityKind defines security kind.
type SecurityKind string

const (
	// QuerySecurity is URL query security parameter. Matches "apiKey" type with "in" = "query".
	QuerySecurity SecurityKind = "query"
	// HeaderSecurity is HTTP header security parameter. Matches some "http" schemes and "apiKey" with "in" = "header".
	HeaderSecurity SecurityKind = "header"
	// OAuth2Security is special type for OAuth2-based authentication. Matches "oauth2" and "openIdConnect".
	OAuth2Security SecurityKind = "oauth2"
)

// IsQuery whether s is QuerySecurity.
func (s SecurityKind) IsQuery() bool {
	return s == QuerySecurity
}

// IsHeader whether s is HeaderSecurity.
func (s SecurityKind) IsHeader() bool {
	return s == HeaderSecurity
}

// IsOAuth2 whether s is OAuth2Security.
func (s SecurityKind) IsOAuth2() bool {
	return s == OAuth2Security
}

// SecurityFormat defines security parameter format.
type SecurityFormat string

const (
	// APIKeySecurityFormat is plain value format.
	APIKeySecurityFormat SecurityFormat = "apiKey"
	// BearerSecurityFormat is Bearer authentication (RFC 6750) format.
	//
	// Unsupported yet.
	BearerSecurityFormat SecurityFormat = "bearer"
	// BasicHTTPSecurityFormat is Basic HTTP authentication (RFC 7617) format.
	BasicHTTPSecurityFormat SecurityFormat = "basic"
	// DigestHTTPSecurityFormat is Digest HTTP authentication (RFC 7616) format.
	//
	// Unsupported yet.
	DigestHTTPSecurityFormat SecurityFormat = "digest"
)

// IsAPIKeySecurity whether s is APIKeySecurityFormat.
func (s SecurityFormat) IsAPIKeySecurity() bool {
	return s == APIKeySecurityFormat
}

// IsBearerSecurity whether s is BearerSecurityFormat.
func (s SecurityFormat) IsBearerSecurity() bool {
	return s == BearerSecurityFormat
}

// IsBasicHTTPSecurity whether s is BasicHTTPSecurityFormat.
func (s SecurityFormat) IsBasicHTTPSecurity() bool {
	return s == BasicHTTPSecurityFormat
}

// IsDigestHTTPSecurity whether s is DigestHTTPSecurityFormat.
func (s SecurityFormat) IsDigestHTTPSecurity() bool {
	return s == DigestHTTPSecurityFormat
}

type Security struct {
	Kind          SecurityKind
	Format        SecurityFormat
	ParameterName string
	Description   string
	Type          *Type
}

func (s *Security) GoDoc() []string {
	return prettyDoc(s.Description)
}

type SecurityRequirement struct {
	Security *Security
	Spec     openapi.SecurityRequirements
}
