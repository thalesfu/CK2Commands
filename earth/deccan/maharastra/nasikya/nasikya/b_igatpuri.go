package nasikya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊格德布里IgatpuriBarony struct {
	feud.BaseBarony
}

var BIgatpuri伊格德布里 feud.Barony = &伊格德布里IgatpuriBarony{}

func init() {
	f := BIgatpuri伊格德布里.(*伊格德布里IgatpuriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "igatpuri",
		TitleName: "伊格德布里",
		TitleCode: "b_igatpuri",
	}
}
