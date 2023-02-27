package znojmo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊赫拉瓦JihlavaBarony struct {
	feud.BaseBarony
}

var BJihlava伊赫拉瓦 feud.Barony = &伊赫拉瓦JihlavaBarony{}

func init() {
    f := BJihlava伊赫拉瓦.(*伊赫拉瓦JihlavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jihlava",
		TitleName: "伊赫拉瓦",
		TitleCode: "b_jihlava",
	}
}
