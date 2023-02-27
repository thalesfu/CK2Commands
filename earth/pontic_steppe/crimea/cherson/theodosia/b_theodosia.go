package theodosia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 狄奥多西亚TheodosiaBarony struct {
	feud.BaseBarony
}

var BTheodosia狄奥多西亚 feud.Barony = &狄奥多西亚TheodosiaBarony{}

func init() {
    f := BTheodosia狄奥多西亚.(*狄奥多西亚TheodosiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "theodosia",
		TitleName: "狄奥多西亚",
		TitleCode: "b_theodosia",
	}
}
