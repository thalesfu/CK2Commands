package glamorgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼思NeathBarony struct {
	feud.BaseBarony
}

var BNeath尼思 feud.Barony = &尼思NeathBarony{}

func init() {
    f := BNeath尼思.(*尼思NeathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neath",
		TitleName: "尼思",
		TitleCode: "b_neath",
	}
}
