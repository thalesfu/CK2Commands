package khokh_serkh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呼赫色尔赫Khokh_serkhBarony struct {
	feud.BaseBarony
}

var BKhokh_serkh呼赫色尔赫 feud.Barony = &呼赫色尔赫Khokh_serkhBarony{}

func init() {
    f := BKhokh_serkh呼赫色尔赫.(*呼赫色尔赫Khokh_serkhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khokh_serkh",
		TitleName: "呼赫色尔赫",
		TitleCode: "b_khokh_serkh",
	}
}
