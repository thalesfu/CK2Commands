package ulm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅明根MemmingenBarony struct {
	feud.BaseBarony
}

var BMemmingen梅明根 feud.Barony = &梅明根MemmingenBarony{}

func init() {
	f := BMemmingen梅明根.(*梅明根MemmingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "memmingen",
		TitleName: "梅明根",
		TitleCode: "b_memmingen",
	}
}
