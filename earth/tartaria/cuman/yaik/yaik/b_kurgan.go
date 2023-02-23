package yaik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔干KurganBarony struct {
	feud.BaseBarony
}

var BKurgan库尔干 feud.Barony = &库尔干KurganBarony{}

func init() {
	f := BKurgan库尔干.(*库尔干KurganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kurgan",
		TitleName: "库尔干",
		TitleCode: "b_kurgan",
	}
}
