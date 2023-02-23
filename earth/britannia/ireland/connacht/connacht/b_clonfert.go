package connacht

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 康菲特ClonfertBarony struct {
	feud.BaseBarony
}

var BClonfert康菲特 feud.Barony = &康菲特ClonfertBarony{}

func init() {
	f := BClonfert康菲特.(*康菲特ClonfertBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clonfert",
		TitleName: "康菲特",
		TitleCode: "b_clonfert",
	}
}
