package kanin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基亚Kiya_kaninBarony struct {
	feud.BaseBarony
}

var BKiya_kanin基亚 feud.Barony = &基亚Kiya_kaninBarony{}

func init() {
    f := BKiya_kanin基亚.(*基亚Kiya_kaninBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiya_kanin",
		TitleName: "基亚",
		TitleCode: "b_kiya_kanin",
	}
}
