package hormuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布穆萨Abu_musaBarony struct {
	feud.BaseBarony
}

var BAbu_musa阿布穆萨 feud.Barony = &阿布穆萨Abu_musaBarony{}

func init() {
    f := BAbu_musa阿布穆萨.(*阿布穆萨Abu_musaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abu_musa",
		TitleName: "阿布穆萨",
		TitleCode: "b_abu_musa",
	}
}
