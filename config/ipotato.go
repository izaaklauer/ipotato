package config

type Ipotato struct {
HelloWorldMessage string `hcl:"hello_world_message,attr"`

// ... your config here
}

// DefaultIpotatoConfig returns default config values
func DefaultIpotatoConfig() Ipotato {
	return Ipotato{
	HelloWorldMessage:
		"hello from the default config",
	}
}
