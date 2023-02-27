package argyll

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯文SweenBarony struct {
	feud.BaseBarony
}

var BSween斯文 feud.Barony = &斯文SweenBarony{}

func init() {
    f := BSween斯文.(*斯文SweenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sween",
		TitleName: "斯文",
		TitleCode: "b_sween",
	}
}
