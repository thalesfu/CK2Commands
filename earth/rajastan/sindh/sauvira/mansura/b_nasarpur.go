package mansura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那萨补罗NasarpurBarony struct {
	feud.BaseBarony
}

var BNasarpur那萨补罗 feud.Barony = &那萨补罗NasarpurBarony{}

func init() {
    f := BNasarpur那萨补罗.(*那萨补罗NasarpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nasarpur",
		TitleName: "那萨补罗",
		TitleCode: "b_nasarpur",
	}
}
