package forcalquier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福卡尔基耶ForcalquierBarony struct {
	feud.BaseBarony
}

var BForcalquier福卡尔基耶 feud.Barony = &福卡尔基耶ForcalquierBarony{}

func init() {
    f := BForcalquier福卡尔基耶.(*福卡尔基耶ForcalquierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "forcalquier",
		TitleName: "福卡尔基耶",
		TitleCode: "b_forcalquier",
	}
}
