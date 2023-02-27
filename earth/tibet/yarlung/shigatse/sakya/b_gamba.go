package sakya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 岗巴GambaBarony struct {
	feud.BaseBarony
}

var BGamba岗巴 feud.Barony = &岗巴GambaBarony{}

func init() {
    f := BGamba岗巴.(*岗巴GambaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gamba",
		TitleName: "岗巴",
		TitleCode: "b_gamba",
	}
}
