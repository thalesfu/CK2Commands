package yezhuga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维瓦VyvaBarony struct {
	feud.BaseBarony
}

var BVyva维瓦 feud.Barony = &维瓦VyvaBarony{}

func init() {
    f := BVyva维瓦.(*维瓦VyvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vyva",
		TitleName: "维瓦",
		TitleCode: "b_vyva",
	}
}
