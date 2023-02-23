package como

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱科LeccoBarony struct {
	feud.BaseBarony
}

var BLecco莱科 feud.Barony = &莱科LeccoBarony{}

func init() {
	f := BLecco莱科.(*莱科LeccoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lecco",
		TitleName: "莱科",
		TitleCode: "b_lecco",
	}
}
