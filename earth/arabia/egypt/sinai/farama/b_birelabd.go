package farama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿卜德井BirelabdBarony struct {
	feud.BaseBarony
}

var BBirelabd阿卜德井 feud.Barony = &阿卜德井BirelabdBarony{}

func init() {
    f := BBirelabd阿卜德井.(*阿卜德井BirelabdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "birelabd",
		TitleName: "阿卜德井",
		TitleCode: "b_birelabd",
	}
}
