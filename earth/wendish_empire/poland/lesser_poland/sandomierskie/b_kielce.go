package sandomierskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯尔采KielceBarony struct {
	feud.BaseBarony
}

var BKielce凯尔采 feud.Barony = &凯尔采KielceBarony{}

func init() {
    f := BKielce凯尔采.(*凯尔采KielceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kielce",
		TitleName: "凯尔采",
		TitleCode: "b_kielce",
	}
}
