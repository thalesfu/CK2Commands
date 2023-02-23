package betpaqdala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奎甘KuyganBarony struct {
	feud.BaseBarony
}

var BKuygan奎甘 feud.Barony = &奎甘KuyganBarony{}

func init() {
	f := BKuygan奎甘.(*奎甘KuyganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuygan",
		TitleName: "奎甘",
		TitleCode: "b_kuygan",
	}
}
