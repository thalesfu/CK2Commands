package bobo_dyulasso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔巴TabaBarony struct {
	feud.BaseBarony
}

var BTaba塔巴 feud.Barony = &塔巴TabaBarony{}

func init() {
    f := BTaba塔巴.(*塔巴TabaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taba",
		TitleName: "塔巴",
		TitleCode: "b_taba",
	}
}
