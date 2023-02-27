package brno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔诺BrnoBarony struct {
	feud.BaseBarony
}

var BBrno布尔诺 feud.Barony = &布尔诺BrnoBarony{}

func init() {
    f := BBrno布尔诺.(*布尔诺BrnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brno",
		TitleName: "布尔诺",
		TitleCode: "b_brno",
	}
}
