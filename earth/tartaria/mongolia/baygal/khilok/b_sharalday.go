package khilok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙喇勒岱SharaldayBarony struct {
	feud.BaseBarony
}

var BSharalday沙喇勒岱 feud.Barony = &沙喇勒岱SharaldayBarony{}

func init() {
    f := BSharalday沙喇勒岱.(*沙喇勒岱SharaldayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sharalday",
		TitleName: "沙喇勒岱",
		TitleCode: "b_sharalday",
	}
}
