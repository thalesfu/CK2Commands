package finnmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔格赫亚VarghoeyaBarony struct {
	feud.BaseBarony
}

var BVarghoeya瓦尔格赫亚 feud.Barony = &瓦尔格赫亚VarghoeyaBarony{}

func init() {
	f := BVarghoeya瓦尔格赫亚.(*瓦尔格赫亚VarghoeyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varghoeya",
		TitleName: "瓦尔格赫亚",
		TitleCode: "b_varghoeya",
	}
}
