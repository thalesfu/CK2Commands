package corsica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢里LuriBarony struct {
	feud.BaseBarony
}

var BLuri卢里 feud.Barony = &卢里LuriBarony{}

func init() {
	f := BLuri卢里.(*卢里LuriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luri",
		TitleName: "卢里",
		TitleCode: "b_luri",
	}
}
