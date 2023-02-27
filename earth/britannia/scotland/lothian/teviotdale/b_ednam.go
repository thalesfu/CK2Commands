package teviotdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃德纳姆EdnamBarony struct {
	feud.BaseBarony
}

var BEdnam埃德纳姆 feud.Barony = &埃德纳姆EdnamBarony{}

func init() {
    f := BEdnam埃德纳姆.(*埃德纳姆EdnamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ednam",
		TitleName: "埃德纳姆",
		TitleCode: "b_ednam",
	}
}
