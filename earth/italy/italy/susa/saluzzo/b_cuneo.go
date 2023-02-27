package saluzzo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库内奥CuneoBarony struct {
	feud.BaseBarony
}

var BCuneo库内奥 feud.Barony = &库内奥CuneoBarony{}

func init() {
    f := BCuneo库内奥.(*库内奥CuneoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cuneo",
		TitleName: "库内奥",
		TitleCode: "b_cuneo",
	}
}
