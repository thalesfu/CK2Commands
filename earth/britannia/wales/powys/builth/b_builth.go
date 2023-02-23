package builth

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比尔斯BuilthBarony struct {
	feud.BaseBarony
}

var BBuilth比尔斯 feud.Barony = &比尔斯BuilthBarony{}

func init() {
	f := BBuilth比尔斯.(*比尔斯BuilthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "builth",
		TitleName: "比尔斯",
		TitleCode: "b_builth",
	}
}
