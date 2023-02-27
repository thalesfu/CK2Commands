package teviotdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰德堡JedburghBarony struct {
	feud.BaseBarony
}

var BJedburgh杰德堡 feud.Barony = &杰德堡JedburghBarony{}

func init() {
    f := BJedburgh杰德堡.(*杰德堡JedburghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jedburgh",
		TitleName: "杰德堡",
		TitleCode: "b_jedburgh",
	}
}
