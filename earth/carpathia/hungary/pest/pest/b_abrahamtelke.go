package pest

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布劳哈姆泰尔凯AbrahamtelkeBarony struct {
	feud.BaseBarony
}

var BAbrahamtelke阿布劳哈姆泰尔凯 feud.Barony = &阿布劳哈姆泰尔凯AbrahamtelkeBarony{}

func init() {
    f := BAbrahamtelke阿布劳哈姆泰尔凯.(*阿布劳哈姆泰尔凯AbrahamtelkeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abrahamtelke",
		TitleName: "阿布劳哈姆泰尔凯",
		TitleCode: "b_abrahamtelke",
	}
}
