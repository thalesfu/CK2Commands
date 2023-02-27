package tell_bashir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽乌格马ZeugmaBarony struct {
	feud.BaseBarony
}

var BZeugma泽乌格马 feud.Barony = &泽乌格马ZeugmaBarony{}

func init() {
    f := BZeugma泽乌格马.(*泽乌格马ZeugmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zeugma",
		TitleName: "泽乌格马",
		TitleCode: "b_zeugma",
	}
}
