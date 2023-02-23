package hisar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯迦AsikaBarony struct {
	feud.BaseBarony
}

var BAsika阿斯迦 feud.Barony = &阿斯迦AsikaBarony{}

func init() {
	f := BAsika阿斯迦.(*阿斯迦AsikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asika",
		TitleName: "阿斯迦",
		TitleCode: "b_asika",
	}
}
