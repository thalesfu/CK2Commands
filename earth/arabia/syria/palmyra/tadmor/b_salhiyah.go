package tadmor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赛希亚SalhiyahBarony struct {
	feud.BaseBarony
}

var BSalhiyah赛希亚 feud.Barony = &赛希亚SalhiyahBarony{}

func init() {
    f := BSalhiyah赛希亚.(*赛希亚SalhiyahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salhiyah",
		TitleName: "赛希亚",
		TitleCode: "b_salhiyah",
	}
}
