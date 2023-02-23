package murzuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰萨韦TassawaBarony struct {
	feud.BaseBarony
}

var BTassawa泰萨韦 feud.Barony = &泰萨韦TassawaBarony{}

func init() {
	f := BTassawa泰萨韦.(*泰萨韦TassawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tassawa",
		TitleName: "泰萨韦",
		TitleCode: "b_tassawa",
	}
}
