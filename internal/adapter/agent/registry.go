package agent

import (
	"github.com/sajanv88/gocli/internal/adapter/agent/adk"
	"github.com/sajanv88/gocli/internal/adapter/agent/enio"
	"github.com/sajanv88/gocli/internal/domain"
)

func All() map[domain.AgentOption]domain.AgentGenerator {
	return map[domain.AgentOption]domain.AgentGenerator{
		domain.AgentADK:  adk.Generator{},
		domain.AgentEino: enio.Generator{},
	}
}
