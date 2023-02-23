package saargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅尔克林根MerkingenBarony struct {
	feud.BaseBarony
}

var BMerkingen梅尔克林根 feud.Barony = &梅尔克林根MerkingenBarony{}

func init() {
	f := BMerkingen梅尔克林根.(*梅尔克林根MerkingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "merkingen",
		TitleName: "梅尔克林根",
		TitleCode: "b_merkingen",
	}
}
