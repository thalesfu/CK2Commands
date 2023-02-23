package alcantara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉马塔LamataBarony struct {
	feud.BaseBarony
}

var BLamata拉马塔 feud.Barony = &拉马塔LamataBarony{}

func init() {
	f := BLamata拉马塔.(*拉马塔LamataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lamata",
		TitleName: "拉马塔",
		TitleCode: "b_lamata",
	}
}
