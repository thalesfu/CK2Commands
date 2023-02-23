package toropets

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥赫瓦特OkhvatBarony struct {
	feud.BaseBarony
}

var BOkhvat奥赫瓦特 feud.Barony = &奥赫瓦特OkhvatBarony{}

func init() {
	f := BOkhvat奥赫瓦特.(*奥赫瓦特OkhvatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "okhvat",
		TitleName: "奥赫瓦特",
		TitleCode: "b_okhvat",
	}
}
