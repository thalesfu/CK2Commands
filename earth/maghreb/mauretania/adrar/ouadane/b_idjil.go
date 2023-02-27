package ouadane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊吉勒IdjilBarony struct {
	feud.BaseBarony
}

var BIdjil伊吉勒 feud.Barony = &伊吉勒IdjilBarony{}

func init() {
    f := BIdjil伊吉勒.(*伊吉勒IdjilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "idjil",
		TitleName: "伊吉勒",
		TitleCode: "b_idjil",
	}
}
