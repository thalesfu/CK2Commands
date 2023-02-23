package bizerte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坚杜拜JendoubaBarony struct {
	feud.BaseBarony
}

var BJendouba坚杜拜 feud.Barony = &坚杜拜JendoubaBarony{}

func init() {
	f := BJendouba坚杜拜.(*坚杜拜JendoubaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jendouba",
		TitleName: "坚杜拜",
		TitleCode: "b_jendouba",
	}
}
