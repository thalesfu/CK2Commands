package idel_ural

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Idel_uralEmpire interface {
    feud.Empire
    KNenets涅涅茨() 	nenets.NenetsKingdom
    KPerm彼尔姆() 	perm.PermKingdom
    KVolga_bulgaria伏尔加保加利亚() 	volga_bulgaria.Volga_bulgariaKingdom
}

type 伏尔加_乌拉尔Idel_uralEmpire struct {
	feud.BaseEmpire
	涅涅茨Nenets 	nenets.NenetsKingdom
	彼尔姆Perm 	perm.PermKingdom
	伏尔加保加利亚Volga_bulgaria 	volga_bulgaria.Volga_bulgariaKingdom
}

func (e *伏尔加_乌拉尔Idel_uralEmpire) KNenets涅涅茨() nenets.NenetsKingdom {
	return e.涅涅茨Nenets
}
    
func (e *伏尔加_乌拉尔Idel_uralEmpire) KPerm彼尔姆() perm.PermKingdom {
	return e.彼尔姆Perm
}
    
func (e *伏尔加_乌拉尔Idel_uralEmpire) KVolga_bulgaria伏尔加保加利亚() volga_bulgaria.Volga_bulgariaKingdom {
	return e.伏尔加保加利亚Volga_bulgaria
}
    
var EIdel_ural伏尔加_乌拉尔 Idel_uralEmpire = &伏尔加_乌拉尔Idel_uralEmpire{}

func init() {
	f := EIdel_ural伏尔加_乌拉尔.(*伏尔加_乌拉尔Idel_uralEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "idel_ural",
		TitleName: "伏尔加_乌拉尔",
		TitleCode: "e_idel_ural",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.涅涅茨Nenets = nenets.KNenets涅涅茨
	f.涅涅茨Nenets.SetParent(f)
	f.彼尔姆Perm = perm.KPerm彼尔姆
	f.彼尔姆Perm.SetParent(f)
	f.伏尔加保加利亚Volga_bulgaria = volga_bulgaria.KVolga_bulgaria伏尔加保加利亚
	f.伏尔加保加利亚Volga_bulgaria.SetParent(f)
}
