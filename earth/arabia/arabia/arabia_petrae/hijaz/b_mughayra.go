package hijaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆盖拉MughayraBarony struct {
	feud.BaseBarony
}

var BMughayra穆盖拉 feud.Barony = &穆盖拉MughayraBarony{}

func init() {
    f := BMughayra穆盖拉.(*穆盖拉MughayraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mughayra",
		TitleName: "穆盖拉",
		TitleCode: "b_mughayra",
	}
}
