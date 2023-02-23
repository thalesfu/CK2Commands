package gastrikland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔博ValboBarony struct {
	feud.BaseBarony
}

var BValbo瓦尔博 feud.Barony = &瓦尔博ValboBarony{}

func init() {
	f := BValbo瓦尔博.(*瓦尔博ValboBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valbo",
		TitleName: "瓦尔博",
		TitleCode: "b_valbo",
	}
}
