package dalarna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 许斯比HusbyBarony struct {
	feud.BaseBarony
}

var BHusby许斯比 feud.Barony = &许斯比HusbyBarony{}

func init() {
	f := BHusby许斯比.(*许斯比HusbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "husby",
		TitleName: "许斯比",
		TitleCode: "b_husby",
	}
}
