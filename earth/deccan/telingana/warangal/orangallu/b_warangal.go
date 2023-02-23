package orangallu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆浪伽罗WarangalBarony struct {
	feud.BaseBarony
}

var BWarangal婆浪伽罗 feud.Barony = &婆浪伽罗WarangalBarony{}

func init() {
	f := BWarangal婆浪伽罗.(*婆浪伽罗WarangalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "warangal",
		TitleName: "婆浪伽罗",
		TitleCode: "b_warangal",
	}
}
