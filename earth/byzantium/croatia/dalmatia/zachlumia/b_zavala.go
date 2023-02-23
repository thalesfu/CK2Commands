package zachlumia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎瓦拉ZavalaBarony struct {
	feud.BaseBarony
}

var BZavala扎瓦拉 feud.Barony = &扎瓦拉ZavalaBarony{}

func init() {
	f := BZavala扎瓦拉.(*扎瓦拉ZavalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zavala",
		TitleName: "扎瓦拉",
		TitleCode: "b_zavala",
	}
}
