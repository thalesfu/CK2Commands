package nafusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞杰南SejenaneBarony struct {
	feud.BaseBarony
}

var BSejenane塞杰南 feud.Barony = &塞杰南SejenaneBarony{}

func init() {
	f := BSejenane塞杰南.(*塞杰南SejenaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sejenane",
		TitleName: "塞杰南",
		TitleCode: "b_sejenane",
	}
}
