package tagant

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提希特TichittBarony struct {
	feud.BaseBarony
}

var BTichitt提希特 feud.Barony = &提希特TichittBarony{}

func init() {
    f := BTichitt提希特.(*提希特TichittBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tichitt",
		TitleName: "提希特",
		TitleCode: "b_tichitt",
	}
}
