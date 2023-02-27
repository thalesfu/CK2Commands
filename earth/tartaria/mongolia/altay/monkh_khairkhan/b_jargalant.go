package monkh_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎尔嘎朗特JargalantBarony struct {
	feud.BaseBarony
}

var BJargalant扎尔嘎朗特 feud.Barony = &扎尔嘎朗特JargalantBarony{}

func init() {
    f := BJargalant扎尔嘎朗特.(*扎尔嘎朗特JargalantBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jargalant",
		TitleName: "扎尔嘎朗特",
		TitleCode: "b_jargalant",
	}
}
