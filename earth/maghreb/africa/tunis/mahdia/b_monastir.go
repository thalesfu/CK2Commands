package mahdia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫纳斯提尔MonastirBarony struct {
	feud.BaseBarony
}

var BMonastir莫纳斯提尔 feud.Barony = &莫纳斯提尔MonastirBarony{}

func init() {
    f := BMonastir莫纳斯提尔.(*莫纳斯提尔MonastirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monastir",
		TitleName: "莫纳斯提尔",
		TitleCode: "b_monastir",
	}
}
