package models

type Entrypoints struct {
	Entrypoints_target string
	// The target of the entrypoint.

	Entrypoints_host string
	// The host of the entrypoint.

	Entrypoints_tags []string
	// The list of sharding tags associated with this entrypoint.
}
