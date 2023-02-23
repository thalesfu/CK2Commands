package auvergne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米罗勒MurolBarony struct {
	feud.BaseBarony
}

var BMurol米罗勒 feud.Barony = &米罗勒MurolBarony{}

func init() {
	f := BMurol米罗勒.(*米罗勒MurolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "murol",
		TitleName: "米罗勒",
		TitleCode: "b_murol",
	}
}
