package kham

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/kham/dege"
	"github.com/thalesfu/CK2Commands/earth/tibet/kham/nyingchi"
	"github.com/thalesfu/CK2Commands/earth/tibet/kham/qamdo"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhamKingdom interface {
    feud.Kingdom
    DDege德格() 	dege.DegeDuke
    DNyingchi尼池() 	nyingchi.NyingchiDuke
    DQamdo察木多() 	qamdo.QamdoDuke
}

type 康区KhamKingdom struct {
	feud.BaseKingdom
	德格Dege 	dege.DegeDuke
	尼池Nyingchi 	nyingchi.NyingchiDuke
	察木多Qamdo 	qamdo.QamdoDuke
}

func (k *康区KhamKingdom) DDege德格() dege.DegeDuke {
	return k.德格Dege
}
    
func (k *康区KhamKingdom) DNyingchi尼池() nyingchi.NyingchiDuke {
	return k.尼池Nyingchi
}
    
func (k *康区KhamKingdom) DQamdo察木多() qamdo.QamdoDuke {
	return k.察木多Qamdo
}
    
var KKham康区 KhamKingdom = &康区KhamKingdom{}

func init() {
	f := KKham康区.(*康区KhamKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "kham",
		TitleName: "康区",
		TitleCode: "k_kham",
		Dukes:  map[string]feud.Duke{},
	}

	f.德格Dege = dege.DDege德格
	f.德格Dege.SetParent(f)
	
	f.尼池Nyingchi = nyingchi.DNyingchi尼池
	f.尼池Nyingchi.SetParent(f)
	
	f.察木多Qamdo = qamdo.DQamdo察木多
	f.察木多Qamdo.SetParent(f)
	
}
