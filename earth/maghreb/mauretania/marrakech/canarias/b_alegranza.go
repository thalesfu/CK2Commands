package canarias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿莱格兰萨AlegranzaBarony struct {
	feud.BaseBarony
}

var BAlegranza阿莱格兰萨 feud.Barony = &阿莱格兰萨AlegranzaBarony{}

func init() {
    f := BAlegranza阿莱格兰萨.(*阿莱格兰萨AlegranzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alegranza",
		TitleName: "阿莱格兰萨",
		TitleCode: "b_alegranza",
	}
}
