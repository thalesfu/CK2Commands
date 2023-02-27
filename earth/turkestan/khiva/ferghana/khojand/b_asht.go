package khojand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿什特AshtBarony struct {
	feud.BaseBarony
}

var BAsht阿什特 feud.Barony = &阿什特AshtBarony{}

func init() {
    f := BAsht阿什特.(*阿什特AshtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asht",
		TitleName: "阿什特",
		TitleCode: "b_asht",
	}
}
