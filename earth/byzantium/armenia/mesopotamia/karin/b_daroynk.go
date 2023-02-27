package karin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达洛克DaroynkBarony struct {
	feud.BaseBarony
}

var BDaroynk达洛克 feud.Barony = &达洛克DaroynkBarony{}

func init() {
    f := BDaroynk达洛克.(*达洛克DaroynkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daroynk",
		TitleName: "达洛克",
		TitleCode: "b_daroynk",
	}
}
