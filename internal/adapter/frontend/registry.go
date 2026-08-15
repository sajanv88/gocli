package frontend

import "github.com/sajanv88/gocli/internal/domain"

func All() map[domain.FrontendOption]domain.FrontendGenerator {
	return map[domain.FrontendOption]domain.FrontendGenerator{
		domain.FrontendVite:   ViteGenerator{},
		domain.FrontendNextJS: NextGenerator{},
	}
}
