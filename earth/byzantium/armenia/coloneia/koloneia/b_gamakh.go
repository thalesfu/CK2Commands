package koloneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯马赫GamakhBarony struct {
	feud.BaseBarony
}

var BGamakh凯马赫 feud.Barony = &凯马赫GamakhBarony{}

func init() {
	f := BGamakh凯马赫.(*凯马赫GamakhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gamakh",
		TitleName: "凯马赫",
		TitleCode: "b_gamakh",
	}
}
