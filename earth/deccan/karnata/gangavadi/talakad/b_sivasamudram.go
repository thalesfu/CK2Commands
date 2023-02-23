package talakad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 湿婆三慕达罗SivasamudramBarony struct {
	feud.BaseBarony
}

var BSivasamudram湿婆三慕达罗 feud.Barony = &湿婆三慕达罗SivasamudramBarony{}

func init() {
	f := BSivasamudram湿婆三慕达罗.(*湿婆三慕达罗SivasamudramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sivasamudram",
		TitleName: "湿婆三慕达罗",
		TitleCode: "b_sivasamudram",
	}
}
