package sakmara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽尔克洛ZerkloBarony struct {
	feud.BaseBarony
}

var BZerklo泽尔克洛 feud.Barony = &泽尔克洛ZerkloBarony{}

func init() {
    f := BZerklo泽尔克洛.(*泽尔克洛ZerkloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zerklo",
		TitleName: "泽尔克洛",
		TitleCode: "b_zerklo",
	}
}
