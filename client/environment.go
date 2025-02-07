package client

// Environment represents the Lalamove API environment
type Environment string

const (
	// Sandbox environment for testing
	Sandbox Environment = "sandbox"
	// Production environment for live requests
	Production Environment = "production"
)

// GetBaseURL returns the base URL for the environment
func (e Environment) GetBaseURL() string {
	switch e {
	case Sandbox:
		return "https://rest.sandbox.lalamove.com/v3"
	case Production:
		return "https://rest.lalamove.com/v3"
	default:
		return "https://rest.sandbox.lalamove.com/v3"
	}
} 