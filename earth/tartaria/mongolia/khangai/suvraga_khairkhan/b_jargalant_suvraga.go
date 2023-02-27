package suvraga_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎尔嘎朗特Jargalant_suvragaBarony struct {
	feud.BaseBarony
}

var BJargalant_suvraga扎尔嘎朗特 feud.Barony = &扎尔嘎朗特Jargalant_suvragaBarony{}

func init() {
    f := BJargalant_suvraga扎尔嘎朗特.(*扎尔嘎朗特Jargalant_suvragaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jargalant_suvraga",
		TitleName: "扎尔嘎朗特",
		TitleCode: "b_jargalant_suvraga",
	}
}
