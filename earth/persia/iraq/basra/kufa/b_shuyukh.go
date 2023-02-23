package kufa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舒尤赫ShuyukhBarony struct {
	feud.BaseBarony
}

var BShuyukh舒尤赫 feud.Barony = &舒尤赫ShuyukhBarony{}

func init() {
	f := BShuyukh舒尤赫.(*舒尤赫ShuyukhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shuyukh",
		TitleName: "舒尤赫",
		TitleCode: "b_shuyukh",
	}
}
