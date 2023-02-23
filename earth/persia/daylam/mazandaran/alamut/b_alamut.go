package alamut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉穆特AlamutBarony struct {
	feud.BaseBarony
}

var BAlamut阿拉穆特 feud.Barony = &阿拉穆特AlamutBarony{}

func init() {
	f := BAlamut阿拉穆特.(*阿拉穆特AlamutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alamut",
		TitleName: "阿拉穆特",
		TitleCode: "b_alamut",
	}
}
