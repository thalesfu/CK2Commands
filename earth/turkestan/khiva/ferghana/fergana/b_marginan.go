package fergana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 麻耳亦囊MarginanBarony struct {
	feud.BaseBarony
}

var BMarginan麻耳亦囊 feud.Barony = &麻耳亦囊MarginanBarony{}

func init() {
	f := BMarginan麻耳亦囊.(*麻耳亦囊MarginanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marginan",
		TitleName: "麻耳亦囊",
		TitleCode: "b_marginan",
	}
}
