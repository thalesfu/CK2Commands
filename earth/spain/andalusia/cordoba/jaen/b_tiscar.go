package jaen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂斯卡尔TiscarBarony struct {
	feud.BaseBarony
}

var BTiscar蒂斯卡尔 feud.Barony = &蒂斯卡尔TiscarBarony{}

func init() {
    f := BTiscar蒂斯卡尔.(*蒂斯卡尔TiscarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiscar",
		TitleName: "蒂斯卡尔",
		TitleCode: "b_tiscar",
	}
}
