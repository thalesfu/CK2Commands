package karelen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯姆KemBarony struct {
	feud.BaseBarony
}

var BKem凯姆 feud.Barony = &凯姆KemBarony{}

func init() {
    f := BKem凯姆.(*凯姆KemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kem",
		TitleName: "凯姆",
		TitleCode: "b_kem",
	}
}
