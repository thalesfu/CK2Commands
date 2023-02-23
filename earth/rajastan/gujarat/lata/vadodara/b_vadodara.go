package vadodara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆度陀罗VadodaraBarony struct {
	feud.BaseBarony
}

var BVadodara婆度陀罗 feud.Barony = &婆度陀罗VadodaraBarony{}

func init() {
	f := BVadodara婆度陀罗.(*婆度陀罗VadodaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vadodara",
		TitleName: "婆度陀罗",
		TitleCode: "b_vadodara",
	}
}
