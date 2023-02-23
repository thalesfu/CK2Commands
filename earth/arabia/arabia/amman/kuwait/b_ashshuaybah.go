package kuwait

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舒艾拜AshshuaybahBarony struct {
	feud.BaseBarony
}

var BAshshuaybah舒艾拜 feud.Barony = &舒艾拜AshshuaybahBarony{}

func init() {
	f := BAshshuaybah舒艾拜.(*舒艾拜AshshuaybahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ashshuaybah",
		TitleName: "舒艾拜",
		TitleCode: "b_ashshuaybah",
	}
}
