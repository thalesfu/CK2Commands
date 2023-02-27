package macon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达瓦耶DavayeBarony struct {
	feud.BaseBarony
}

var BDavaye达瓦耶 feud.Barony = &达瓦耶DavayeBarony{}

func init() {
    f := BDavaye达瓦耶.(*达瓦耶DavayeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "davaye",
		TitleName: "达瓦耶",
		TitleCode: "b_davaye",
	}
}
