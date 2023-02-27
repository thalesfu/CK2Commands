package toledo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丰萨利达FuensalidaBarony struct {
	feud.BaseBarony
}

var BFuensalida丰萨利达 feud.Barony = &丰萨利达FuensalidaBarony{}

func init() {
    f := BFuensalida丰萨利达.(*丰萨利达FuensalidaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fuensalida",
		TitleName: "丰萨利达",
		TitleCode: "b_fuensalida",
	}
}
