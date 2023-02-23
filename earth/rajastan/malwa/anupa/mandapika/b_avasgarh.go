package mandapika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伐迦罗AvasgarhBarony struct {
	feud.BaseBarony
}

var BAvasgarh阿伐迦罗 feud.Barony = &阿伐迦罗AvasgarhBarony{}

func init() {
	f := BAvasgarh阿伐迦罗.(*阿伐迦罗AvasgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avasgarh",
		TitleName: "阿伐迦罗",
		TitleCode: "b_avasgarh",
	}
}
