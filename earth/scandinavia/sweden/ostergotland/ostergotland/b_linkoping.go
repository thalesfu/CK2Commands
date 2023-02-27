package ostergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 林雪平LinkopingBarony struct {
	feud.BaseBarony
}

var BLinkoping林雪平 feud.Barony = &林雪平LinkopingBarony{}

func init() {
    f := BLinkoping林雪平.(*林雪平LinkopingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "linkoping",
		TitleName: "林雪平",
		TitleCode: "b_linkoping",
	}
}
