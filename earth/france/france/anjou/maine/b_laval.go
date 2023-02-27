package maine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉瓦勒LavalBarony struct {
	feud.BaseBarony
}

var BLaval拉瓦勒 feud.Barony = &拉瓦勒LavalBarony{}

func init() {
    f := BLaval拉瓦勒.(*拉瓦勒LavalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laval",
		TitleName: "拉瓦勒",
		TitleCode: "b_laval",
	}
}
