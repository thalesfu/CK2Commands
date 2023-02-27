package dyrrachion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克鲁耶KrujeBarony struct {
	feud.BaseBarony
}

var BKruje克鲁耶 feud.Barony = &克鲁耶KrujeBarony{}

func init() {
    f := BKruje克鲁耶.(*克鲁耶KrujeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kruje",
		TitleName: "克鲁耶",
		TitleCode: "b_kruje",
	}
}
