package chagai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努什基NushkiBarony struct {
	feud.BaseBarony
}

var BNushki努什基 feud.Barony = &努什基NushkiBarony{}

func init() {
    f := BNushki努什基.(*努什基NushkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nushki",
		TitleName: "努什基",
		TitleCode: "b_nushki",
	}
}
