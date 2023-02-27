package al_alamayn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿莱曼ElalameinBarony struct {
	feud.BaseBarony
}

var BElalamein阿莱曼 feud.Barony = &阿莱曼ElalameinBarony{}

func init() {
    f := BElalamein阿莱曼.(*阿莱曼ElalameinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elalamein",
		TitleName: "阿莱曼",
		TitleCode: "b_elalamein",
	}
}
