package bereg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙卡奇MunkacsBarony struct {
	feud.BaseBarony
}

var BMunkacs蒙卡奇 feud.Barony = &蒙卡奇MunkacsBarony{}

func init() {
	f := BMunkacs蒙卡奇.(*蒙卡奇MunkacsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "munkacs",
		TitleName: "蒙卡奇",
		TitleCode: "b_munkacs",
	}
}
