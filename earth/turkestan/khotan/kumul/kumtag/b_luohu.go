package kumtag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗护LuohuBarony struct {
	feud.BaseBarony
}

var BLuohu罗护 feud.Barony = &罗护LuohuBarony{}

func init() {
	f := BLuohu罗护.(*罗护LuohuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luohu",
		TitleName: "罗护",
		TitleCode: "b_luohu",
	}
}
