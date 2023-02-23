package gurjaratra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提毗TibbiBarony struct {
	feud.BaseBarony
}

var BTibbi提毗 feud.Barony = &提毗TibbiBarony{}

func init() {
	f := BTibbi提毗.(*提毗TibbiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tibbi",
		TitleName: "提毗",
		TitleCode: "b_tibbi",
	}
}
