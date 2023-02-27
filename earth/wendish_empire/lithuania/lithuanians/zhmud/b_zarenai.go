package zhmud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎雷奈ZarenaiBarony struct {
	feud.BaseBarony
}

var BZarenai扎雷奈 feud.Barony = &扎雷奈ZarenaiBarony{}

func init() {
    f := BZarenai扎雷奈.(*扎雷奈ZarenaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zarenai",
		TitleName: "扎雷奈",
		TitleCode: "b_zarenai",
	}
}
