package bihar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比豪尔凯赖斯泰什BiharkeresztesBarony struct {
	feud.BaseBarony
}

var BBiharkeresztes比豪尔凯赖斯泰什 feud.Barony = &比豪尔凯赖斯泰什BiharkeresztesBarony{}

func init() {
	f := BBiharkeresztes比豪尔凯赖斯泰什.(*比豪尔凯赖斯泰什BiharkeresztesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "biharkeresztes",
		TitleName: "比豪尔凯赖斯泰什",
		TitleCode: "b_biharkeresztes",
	}
}
