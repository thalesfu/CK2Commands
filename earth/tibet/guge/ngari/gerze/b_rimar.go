package gerze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日玛RimarBarony struct {
	feud.BaseBarony
}

var BRimar日玛 feud.Barony = &日玛RimarBarony{}

func init() {
	f := BRimar日玛.(*日玛RimarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rimar",
		TitleName: "日玛",
		TitleCode: "b_rimar",
	}
}
