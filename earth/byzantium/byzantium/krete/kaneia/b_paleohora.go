package kaneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕莱奥霍拉PaleohoraBarony struct {
	feud.BaseBarony
}

var BPaleohora帕莱奥霍拉 feud.Barony = &帕莱奥霍拉PaleohoraBarony{}

func init() {
    f := BPaleohora帕莱奥霍拉.(*帕莱奥霍拉PaleohoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paleohora",
		TitleName: "帕莱奥霍拉",
		TitleCode: "b_paleohora",
	}
}
