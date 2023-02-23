package apulia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎尼CannaeBarony struct {
	feud.BaseBarony
}

var BCannae坎尼 feud.Barony = &坎尼CannaeBarony{}

func init() {
	f := BCannae坎尼.(*坎尼CannaeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cannae",
		TitleName: "坎尼",
		TitleCode: "b_cannae",
	}
}
