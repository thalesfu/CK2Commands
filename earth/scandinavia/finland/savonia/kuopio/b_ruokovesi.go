package kuopio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁奥韦西RuokovesiBarony struct {
	feud.BaseBarony
}

var BRuokovesi鲁奥韦西 feud.Barony = &鲁奥韦西RuokovesiBarony{}

func init() {
    f := BRuokovesi鲁奥韦西.(*鲁奥韦西RuokovesiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ruokovesi",
		TitleName: "鲁奥韦西",
		TitleCode: "b_ruokovesi",
	}
}
