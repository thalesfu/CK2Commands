package melitene

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉基KigiBarony struct {
	feud.BaseBarony
}

var BKigi吉基 feud.Barony = &吉基KigiBarony{}

func init() {
    f := BKigi吉基.(*吉基KigiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kigi",
		TitleName: "吉基",
		TitleCode: "b_kigi",
	}
}
