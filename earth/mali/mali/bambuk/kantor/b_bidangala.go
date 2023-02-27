package kantor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比当加拉BidangalaBarony struct {
	feud.BaseBarony
}

var BBidangala比当加拉 feud.Barony = &比当加拉BidangalaBarony{}

func init() {
    f := BBidangala比当加拉.(*比当加拉BidangalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bidangala",
		TitleName: "比当加拉",
		TitleCode: "b_bidangala",
	}
}
