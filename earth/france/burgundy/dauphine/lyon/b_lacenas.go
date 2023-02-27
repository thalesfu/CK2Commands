package lyon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉塞纳LacenasBarony struct {
	feud.BaseBarony
}

var BLacenas拉塞纳 feud.Barony = &拉塞纳LacenasBarony{}

func init() {
    f := BLacenas拉塞纳.(*拉塞纳LacenasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lacenas",
		TitleName: "拉塞纳",
		TitleCode: "b_lacenas",
	}
}
