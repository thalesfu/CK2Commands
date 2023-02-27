package gurvan_saikhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汗洪戈尔KhankhongorBarony struct {
	feud.BaseBarony
}

var BKhankhongor汗洪戈尔 feud.Barony = &汗洪戈尔KhankhongorBarony{}

func init() {
    f := BKhankhongor汗洪戈尔.(*汗洪戈尔KhankhongorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khankhongor",
		TitleName: "汗洪戈尔",
		TitleCode: "b_khankhongor",
	}
}
