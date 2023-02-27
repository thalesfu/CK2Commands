package torres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥塔纳OttanaBarony struct {
	feud.BaseBarony
}

var BOttana奥塔纳 feud.Barony = &奥塔纳OttanaBarony{}

func init() {
    f := BOttana奥塔纳.(*奥塔纳OttanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ottana",
		TitleName: "奥塔纳",
		TitleCode: "b_ottana",
	}
}
