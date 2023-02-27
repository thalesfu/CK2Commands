package aegean_islands

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/aegean_islands/euboia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/aegean_islands/naxos"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Aegean_islandsDuke interface {
    feud.Duke
    CEuboia优卑亚() 	euboia.EuboiaCounty
    CNaxos纳克索斯() 	naxos.NaxosCounty
}

type 爱琴群岛Aegean_islandsDuke struct {
	feud.BaseDuke
	优卑亚Euboia 	euboia.EuboiaCounty
	纳克索斯Naxos 	naxos.NaxosCounty
}

func (d *爱琴群岛Aegean_islandsDuke) CEuboia优卑亚() euboia.EuboiaCounty {
	return d.优卑亚Euboia
}
    
func (d *爱琴群岛Aegean_islandsDuke) CNaxos纳克索斯() naxos.NaxosCounty {
	return d.纳克索斯Naxos
}
    
var DAegean_islands爱琴群岛 Aegean_islandsDuke = &爱琴群岛Aegean_islandsDuke{}

func init() {
	f := DAegean_islands爱琴群岛.(*爱琴群岛Aegean_islandsDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "aegean_islands",
		TitleName: "爱琴群岛",
		TitleCode: "d_aegean_islands",
		Counties:  map[string]feud.County{},
	}

	f.优卑亚Euboia = euboia.CEuboia优卑亚
	f.优卑亚Euboia.SetParent(f)
	
	f.纳克索斯Naxos = naxos.CNaxos纳克索斯
	f.纳克索斯Naxos.SetParent(f)
	
}
