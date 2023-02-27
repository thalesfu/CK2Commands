package gurgi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍卢俱SharogBarony struct {
	feud.BaseBarony
}

var BSharog舍卢俱 feud.Barony = &舍卢俱SharogBarony{}

func init() {
    f := BSharog舍卢俱.(*舍卢俱SharogBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sharog",
		TitleName: "舍卢俱",
		TitleCode: "b_sharog",
	}
}
