package bacs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕尔达尼PardanyBarony struct {
	feud.BaseBarony
}

var BPardany帕尔达尼 feud.Barony = &帕尔达尼PardanyBarony{}

func init() {
    f := BPardany帕尔达尼.(*帕尔达尼PardanyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pardany",
		TitleName: "帕尔达尼",
		TitleCode: "b_pardany",
	}
}
