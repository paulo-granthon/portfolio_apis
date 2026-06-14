package models

// Portfolio is the composed, ordered document for a single user: their profile
// plus every project they took part in (ordered by semester), each carrying the
// user's contributions to it. It is the single source of truth rendered both as
// JSON (for the interactive web view) and as markdown (for the document export).
type Portfolio struct {
	User     User               `json:"user"`
	Projects []PortfolioProject `json:"projects"`
}

// PortfolioProject embeds Project so its fields are inlined in the JSON output,
// and adds the user's narrative participation summary plus their contributions.
type PortfolioProject struct {
	Project
	Participation string                  `json:"participation"`
	Contributions []PortfolioContribution `json:"contributions"`
}

// PortfolioContribution is a single contribution flattened for the document:
// just what the view and the markdown need, with skill names resolved.
type PortfolioContribution struct {
	Id      uint64   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Skills  []string `json:"skills"`
}
