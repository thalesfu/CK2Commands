package kamatapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿婆地耶罗寺AvadiyarkoilBarony struct {
	feud.BaseBarony
}

var BAvadiyarkoil阿婆地耶罗寺 feud.Barony = &阿婆地耶罗寺AvadiyarkoilBarony{}

func init() {
    f := BAvadiyarkoil阿婆地耶罗寺.(*阿婆地耶罗寺AvadiyarkoilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avadiyarkoil",
		TitleName: "阿婆地耶罗寺",
		TitleCode: "b_avadiyarkoil",
	}
}
