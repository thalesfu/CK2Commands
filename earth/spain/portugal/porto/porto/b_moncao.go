package porto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙桑MoncaoBarony struct {
	feud.BaseBarony
}

var BMoncao蒙桑 feud.Barony = &蒙桑MoncaoBarony{}

func init() {
    f := BMoncao蒙桑.(*蒙桑MoncaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moncao",
		TitleName: "蒙桑",
		TitleCode: "b_moncao",
	}
}
