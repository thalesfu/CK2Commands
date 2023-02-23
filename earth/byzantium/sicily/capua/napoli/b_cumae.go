package napoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库迈CumaeBarony struct {
	feud.BaseBarony
}

var BCumae库迈 feud.Barony = &库迈CumaeBarony{}

func init() {
	f := BCumae库迈.(*库迈CumaeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cumae",
		TitleName: "库迈",
		TitleCode: "b_cumae",
	}
}
