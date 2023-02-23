package fuqi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伏俟FuqiBarony struct {
	feud.BaseBarony
}

var BFuqi伏俟 feud.Barony = &伏俟FuqiBarony{}

func init() {
	f := BFuqi伏俟.(*伏俟FuqiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fuqi",
		TitleName: "伏俟",
		TitleCode: "b_fuqi",
	}
}
