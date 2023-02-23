package makuria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯拉特KeratBarony struct {
	feud.BaseBarony
}

var BKerat凯拉特 feud.Barony = &凯拉特KeratBarony{}

func init() {
	f := BKerat凯拉特.(*凯拉特KeratBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kerat",
		TitleName: "凯拉特",
		TitleCode: "b_kerat",
	}
}
