package miass

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米阿斯MiassBarony struct {
	feud.BaseBarony
}

var BMiass米阿斯 feud.Barony = &米阿斯MiassBarony{}

func init() {
    f := BMiass米阿斯.(*米阿斯MiassBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "miass",
		TitleName: "米阿斯",
		TitleCode: "b_miass",
	}
}
