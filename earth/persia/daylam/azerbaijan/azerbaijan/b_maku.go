package azerbaijan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马库MakuBarony struct {
	feud.BaseBarony
}

var BMaku马库 feud.Barony = &马库MakuBarony{}

func init() {
    f := BMaku马库.(*马库MakuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maku",
		TitleName: "马库",
		TitleCode: "b_maku",
	}
}
