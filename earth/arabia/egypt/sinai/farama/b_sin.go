package farama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 训SinBarony struct {
	feud.BaseBarony
}

var BSin训 feud.Barony = &训SinBarony{}

func init() {
    f := BSin训.(*训SinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sin",
		TitleName: "训",
		TitleCode: "b_sin",
	}
}
