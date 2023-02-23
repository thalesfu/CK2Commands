package pecs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫哈奇MohacsBarony struct {
	feud.BaseBarony
}

var BMohacs莫哈奇 feud.Barony = &莫哈奇MohacsBarony{}

func init() {
	f := BMohacs莫哈奇.(*莫哈奇MohacsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mohacs",
		TitleName: "莫哈奇",
		TitleCode: "b_mohacs",
	}
}
