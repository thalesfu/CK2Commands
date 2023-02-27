package gallura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿朱斯AggiusBarony struct {
	feud.BaseBarony
}

var BAggius阿朱斯 feud.Barony = &阿朱斯AggiusBarony{}

func init() {
    f := BAggius阿朱斯.(*阿朱斯AggiusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aggius",
		TitleName: "阿朱斯",
		TitleCode: "b_aggius",
	}
}
