package vestisland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫瓦姆HvammBarony struct {
	feud.BaseBarony
}

var BHvamm赫瓦姆 feud.Barony = &赫瓦姆HvammBarony{}

func init() {
	f := BHvamm赫瓦姆.(*赫瓦姆HvammBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hvamm",
		TitleName: "赫瓦姆",
		TitleCode: "b_hvamm",
	}
}
