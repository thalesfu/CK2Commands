package kandalax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔祖加VarzugaBarony struct {
	feud.BaseBarony
}

var BVarzuga瓦尔祖加 feud.Barony = &瓦尔祖加VarzugaBarony{}

func init() {
	f := BVarzuga瓦尔祖加.(*瓦尔祖加VarzugaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varzuga",
		TitleName: "瓦尔祖加",
		TitleCode: "b_varzuga",
	}
}
