package bam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索尔哈桑Sol_hasanBarony struct {
	feud.BaseBarony
}

var BSol_hasan索尔哈桑 feud.Barony = &索尔哈桑Sol_hasanBarony{}

func init() {
    f := BSol_hasan索尔哈桑.(*索尔哈桑Sol_hasanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sol_hasan",
		TitleName: "索尔哈桑",
		TitleCode: "b_sol_hasan",
	}
}
