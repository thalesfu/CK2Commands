package bikrampur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩陀波甘MeddappakkamBarony struct {
	feud.BaseBarony
}

var BMeddappakkam摩陀波甘 feud.Barony = &摩陀波甘MeddappakkamBarony{}

func init() {
	f := BMeddappakkam摩陀波甘.(*摩陀波甘MeddappakkamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "meddappakkam",
		TitleName: "摩陀波甘",
		TitleCode: "b_meddappakkam",
	}
}
