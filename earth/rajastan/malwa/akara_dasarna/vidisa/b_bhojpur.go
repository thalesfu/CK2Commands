package vidisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒲阇补罗BhojpurBarony struct {
	feud.BaseBarony
}

var BBhojpur蒲阇补罗 feud.Barony = &蒲阇补罗BhojpurBarony{}

func init() {
    f := BBhojpur蒲阇补罗.(*蒲阇补罗BhojpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhojpur",
		TitleName: "蒲阇补罗",
		TitleCode: "b_bhojpur",
	}
}
