package asturias_de_oviedo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂内奥TineoBarony struct {
	feud.BaseBarony
}

var BTineo蒂内奥 feud.Barony = &蒂内奥TineoBarony{}

func init() {
    f := BTineo蒂内奥.(*蒂内奥TineoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tineo",
		TitleName: "蒂内奥",
		TitleCode: "b_tineo",
	}
}
