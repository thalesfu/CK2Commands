package archa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿法米亚FamiaBarony struct {
	feud.BaseBarony
}

var BFamia阿法米亚 feud.Barony = &阿法米亚FamiaBarony{}

func init() {
    f := BFamia阿法米亚.(*阿法米亚FamiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "famia",
		TitleName: "阿法米亚",
		TitleCode: "b_famia",
	}
}
