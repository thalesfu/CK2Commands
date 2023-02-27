package amiens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩罗讷PeronneBarony struct {
	feud.BaseBarony
}

var BPeronne佩罗讷 feud.Barony = &佩罗讷PeronneBarony{}

func init() {
    f := BPeronne佩罗讷.(*佩罗讷PeronneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "peronne",
		TitleName: "佩罗讷",
		TitleCode: "b_peronne",
	}
}
