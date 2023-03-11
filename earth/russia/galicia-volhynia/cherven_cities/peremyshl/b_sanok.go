package peremyshl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨诺克SanokBarony struct {
	feud.BaseBarony
}

var BSanok萨诺克 feud.Barony = &萨诺克SanokBarony{}

func init() {
    f := BSanok萨诺克.(*萨诺克SanokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanok",
		TitleName: "萨诺克",
		TitleCode: "b_sanok",
	}
}
