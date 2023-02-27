package zhmud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤尔巴尔卡斯JurbarkasBarony struct {
	feud.BaseBarony
}

var BJurbarkas尤尔巴尔卡斯 feud.Barony = &尤尔巴尔卡斯JurbarkasBarony{}

func init() {
    f := BJurbarkas尤尔巴尔卡斯.(*尤尔巴尔卡斯JurbarkasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jurbarkas",
		TitleName: "尤尔巴尔卡斯",
		TitleCode: "b_jurbarkas",
	}
}
