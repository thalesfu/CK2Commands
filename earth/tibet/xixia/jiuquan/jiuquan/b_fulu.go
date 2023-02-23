package jiuquan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福禄FuluBarony struct {
	feud.BaseBarony
}

var BFulu福禄 feud.Barony = &福禄FuluBarony{}

func init() {
	f := BFulu福禄.(*福禄FuluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fulu",
		TitleName: "福禄",
		TitleCode: "b_fulu",
	}
}
