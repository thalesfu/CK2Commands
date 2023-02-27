package charkliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔干KarganBarony struct {
	feud.BaseBarony
}

var BKargan库尔干 feud.Barony = &库尔干KarganBarony{}

func init() {
    f := BKargan库尔干.(*库尔干KarganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kargan",
		TitleName: "库尔干",
		TitleCode: "b_kargan",
	}
}
