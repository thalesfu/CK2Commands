package algeciras

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 龙达RondaBarony struct {
	feud.BaseBarony
}

var BRonda龙达 feud.Barony = &龙达RondaBarony{}

func init() {
    f := BRonda龙达.(*龙达RondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ronda",
		TitleName: "龙达",
		TitleCode: "b_ronda",
	}
}
