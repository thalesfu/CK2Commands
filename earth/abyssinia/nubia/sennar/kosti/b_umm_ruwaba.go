package kosti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌姆鲁瓦巴Umm_ruwabaBarony struct {
	feud.BaseBarony
}

var BUmm_ruwaba乌姆鲁瓦巴 feud.Barony = &乌姆鲁瓦巴Umm_ruwabaBarony{}

func init() {
    f := BUmm_ruwaba乌姆鲁瓦巴.(*乌姆鲁瓦巴Umm_ruwabaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "umm_ruwaba",
		TitleName: "乌姆鲁瓦巴",
		TitleCode: "b_umm_ruwaba",
	}
}
