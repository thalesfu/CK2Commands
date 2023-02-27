package sakya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨迦SakyaBarony struct {
	feud.BaseBarony
}

var BSakya萨迦 feud.Barony = &萨迦SakyaBarony{}

func init() {
    f := BSakya萨迦.(*萨迦SakyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sakya",
		TitleName: "萨迦",
		TitleCode: "b_sakya",
	}
}
