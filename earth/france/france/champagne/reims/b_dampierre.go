package reims

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 当皮耶尔DampierreBarony struct {
	feud.BaseBarony
}

var BDampierre当皮耶尔 feud.Barony = &当皮耶尔DampierreBarony{}

func init() {
    f := BDampierre当皮耶尔.(*当皮耶尔DampierreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dampierre",
		TitleName: "当皮耶尔",
		TitleCode: "b_dampierre",
	}
}
