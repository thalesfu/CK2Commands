package zeila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿姆德AmudBarony struct {
	feud.BaseBarony
}

var BAmud阿姆德 feud.Barony = &阿姆德AmudBarony{}

func init() {
    f := BAmud阿姆德.(*阿姆德AmudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amud",
		TitleName: "阿姆德",
		TitleCode: "b_amud",
	}
}
