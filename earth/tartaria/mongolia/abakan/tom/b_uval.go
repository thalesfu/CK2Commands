package tom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌瓦尔UvalBarony struct {
	feud.BaseBarony
}

var BUval乌瓦尔 feud.Barony = &乌瓦尔UvalBarony{}

func init() {
    f := BUval乌瓦尔.(*乌瓦尔UvalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uval",
		TitleName: "乌瓦尔",
		TitleCode: "b_uval",
	}
}
