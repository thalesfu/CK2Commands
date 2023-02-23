package chandax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉西锡LassithiBarony struct {
	feud.BaseBarony
}

var BLassithi拉西锡 feud.Barony = &拉西锡LassithiBarony{}

func init() {
	f := BLassithi拉西锡.(*拉西锡LassithiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lassithi",
		TitleName: "拉西锡",
		TitleCode: "b_lassithi",
	}
}
