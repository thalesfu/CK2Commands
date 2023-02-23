package nevers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔特奈CourtenayBarony struct {
	feud.BaseBarony
}

var BCourtenay库尔特奈 feud.Barony = &库尔特奈CourtenayBarony{}

func init() {
	f := BCourtenay库尔特奈.(*库尔特奈CourtenayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "courtenay",
		TitleName: "库尔特奈",
		TitleCode: "b_courtenay",
	}
}
