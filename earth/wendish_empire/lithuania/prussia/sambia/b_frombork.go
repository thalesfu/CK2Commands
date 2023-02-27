package sambia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗龙堡FromborkBarony struct {
	feud.BaseBarony
}

var BFrombork弗龙堡 feud.Barony = &弗龙堡FromborkBarony{}

func init() {
    f := BFrombork弗龙堡.(*弗龙堡FromborkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "frombork",
		TitleName: "弗龙堡",
		TitleCode: "b_frombork",
	}
}
