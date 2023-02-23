package ephesos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 以弗所EphesosBarony struct {
	feud.BaseBarony
}

var BEphesos以弗所 feud.Barony = &以弗所EphesosBarony{}

func init() {
	f := BEphesos以弗所.(*以弗所EphesosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ephesos",
		TitleName: "以弗所",
		TitleCode: "b_ephesos",
	}
}
