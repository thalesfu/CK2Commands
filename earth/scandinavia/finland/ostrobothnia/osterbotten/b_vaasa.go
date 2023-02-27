package osterbotten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦萨VaasaBarony struct {
	feud.BaseBarony
}

var BVaasa瓦萨 feud.Barony = &瓦萨VaasaBarony{}

func init() {
    f := BVaasa瓦萨.(*瓦萨VaasaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vaasa",
		TitleName: "瓦萨",
		TitleCode: "b_vaasa",
	}
}
