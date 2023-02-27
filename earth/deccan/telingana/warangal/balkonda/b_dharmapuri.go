package balkonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达摩补利DharmapuriBarony struct {
	feud.BaseBarony
}

var BDharmapuri达摩补利 feud.Barony = &达摩补利DharmapuriBarony{}

func init() {
    f := BDharmapuri达摩补利.(*达摩补利DharmapuriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dharmapuri",
		TitleName: "达摩补利",
		TitleCode: "b_dharmapuri",
	}
}
