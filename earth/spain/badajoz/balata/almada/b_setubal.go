package almada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞图巴尔SetubalBarony struct {
	feud.BaseBarony
}

var BSetubal塞图巴尔 feud.Barony = &塞图巴尔SetubalBarony{}

func init() {
	f := BSetubal塞图巴尔.(*塞图巴尔SetubalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "setubal",
		TitleName: "塞图巴尔",
		TitleCode: "b_setubal",
	}
}
