package nagormo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 察尔汗QarhamBarony struct {
	feud.BaseBarony
}

var BQarham察尔汗 feud.Barony = &察尔汗QarhamBarony{}

func init() {
	f := BQarham察尔汗.(*察尔汗QarhamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qarham",
		TitleName: "察尔汗",
		TitleCode: "b_qarham",
	}
}
