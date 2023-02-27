package westfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯哈亨SchagenBarony struct {
	feud.BaseBarony
}

var BSchagen斯哈亨 feud.Barony = &斯哈亨SchagenBarony{}

func init() {
    f := BSchagen斯哈亨.(*斯哈亨SchagenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schagen",
		TitleName: "斯哈亨",
		TitleCode: "b_schagen",
	}
}
