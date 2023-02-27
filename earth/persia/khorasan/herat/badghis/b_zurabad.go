package badghis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祖拉巴德ZurabadBarony struct {
	feud.BaseBarony
}

var BZurabad祖拉巴德 feud.Barony = &祖拉巴德ZurabadBarony{}

func init() {
    f := BZurabad祖拉巴德.(*祖拉巴德ZurabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zurabad",
		TitleName: "祖拉巴德",
		TitleCode: "b_zurabad",
	}
}
