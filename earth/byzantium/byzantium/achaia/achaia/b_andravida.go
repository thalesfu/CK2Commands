package achaia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安兹拉维扎AndravidaBarony struct {
	feud.BaseBarony
}

var BAndravida安兹拉维扎 feud.Barony = &安兹拉维扎AndravidaBarony{}

func init() {
	f := BAndravida安兹拉维扎.(*安兹拉维扎AndravidaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "andravida",
		TitleName: "安兹拉维扎",
		TitleCode: "b_andravida",
	}
}
