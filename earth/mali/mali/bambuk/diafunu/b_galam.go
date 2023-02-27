package diafunu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加拉姆GalamBarony struct {
	feud.BaseBarony
}

var BGalam加拉姆 feud.Barony = &加拉姆GalamBarony{}

func init() {
    f := BGalam加拉姆.(*加拉姆GalamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "galam",
		TitleName: "加拉姆",
		TitleCode: "b_galam",
	}
}
