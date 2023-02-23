package worcester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊夫舍姆EveshamBarony struct {
	feud.BaseBarony
}

var BEvesham伊夫舍姆 feud.Barony = &伊夫舍姆EveshamBarony{}

func init() {
	f := BEvesham伊夫舍姆.(*伊夫舍姆EveshamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "evesham",
		TitleName: "伊夫舍姆",
		TitleCode: "b_evesham",
	}
}
