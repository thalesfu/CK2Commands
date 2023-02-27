package vetluga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 夏瓦SyavaBarony struct {
	feud.BaseBarony
}

var BSyava夏瓦 feud.Barony = &夏瓦SyavaBarony{}

func init() {
    f := BSyava夏瓦.(*夏瓦SyavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "syava",
		TitleName: "夏瓦",
		TitleCode: "b_syava",
	}
}
