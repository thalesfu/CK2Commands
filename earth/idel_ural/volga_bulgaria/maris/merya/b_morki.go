package merya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫尔基MorkiBarony struct {
	feud.BaseBarony
}

var BMorki莫尔基 feud.Barony = &莫尔基MorkiBarony{}

func init() {
    f := BMorki莫尔基.(*莫尔基MorkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "morki",
		TitleName: "莫尔基",
		TitleCode: "b_morki",
	}
}
