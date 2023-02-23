package methone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安德鲁萨AndrousaBarony struct {
	feud.BaseBarony
}

var BAndrousa安德鲁萨 feud.Barony = &安德鲁萨AndrousaBarony{}

func init() {
	f := BAndrousa安德鲁萨.(*安德鲁萨AndrousaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "androusa",
		TitleName: "安德鲁萨",
		TitleCode: "b_androusa",
	}
}
