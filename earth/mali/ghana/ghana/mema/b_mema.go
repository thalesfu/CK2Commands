package mema

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅马MemaBarony struct {
	feud.BaseBarony
}

var BMema梅马 feud.Barony = &梅马MemaBarony{}

func init() {
	f := BMema梅马.(*梅马MemaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mema",
		TitleName: "梅马",
		TitleCode: "b_mema",
	}
}
