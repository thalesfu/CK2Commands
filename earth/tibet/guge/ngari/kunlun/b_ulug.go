package kunlun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌鲁克UlugBarony struct {
	feud.BaseBarony
}

var BUlug乌鲁克 feud.Barony = &乌鲁克UlugBarony{}

func init() {
    f := BUlug乌鲁克.(*乌鲁克UlugBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulug",
		TitleName: "乌鲁克",
		TitleCode: "b_ulug",
	}
}
