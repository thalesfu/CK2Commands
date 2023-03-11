package bryansk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波切普PochepBarony struct {
	feud.BaseBarony
}

var BPochep波切普 feud.Barony = &波切普PochepBarony{}

func init() {
    f := BPochep波切普.(*波切普PochepBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pochep",
		TitleName: "波切普",
		TitleCode: "b_pochep",
	}
}
