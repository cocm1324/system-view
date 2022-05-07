// Package passiveredundancy is representing passive redundancy in system design
// When single point of failure happens, availability of system is affected
// To avoid affecting availability, we can make copy of system component to replace when one of system component fails
// Passive redundancy is technique which prepares set of copies of system component and process the task parerelly.
// When one of component fails, other copies of system persists to process.
package passiveredundancy

import "github.com/cocm1324/system-view/pkg/service"

type PRModel struct {
	MaxServiceCount int
	Services        map[int]service.Service
}

func (p *PRModel) init() {

}
