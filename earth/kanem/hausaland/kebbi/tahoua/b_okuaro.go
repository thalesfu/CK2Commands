package tahoua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥夸罗OkuaroBarony struct {
	feud.BaseBarony
}

var BOkuaro奥夸罗 feud.Barony = &奥夸罗OkuaroBarony{}

func init() {
	f := BOkuaro奥夸罗.(*奥夸罗OkuaroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "okuaro",
		TitleName: "奥夸罗",
		TitleCode: "b_okuaro",
	}
}
