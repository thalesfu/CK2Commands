package benevento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔扎ConzaBarony struct {
	feud.BaseBarony
}

var BConza孔扎 feud.Barony = &孔扎ConzaBarony{}

func init() {
    f := BConza孔扎.(*孔扎ConzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "conza",
		TitleName: "孔扎",
		TitleCode: "b_conza",
	}
}
