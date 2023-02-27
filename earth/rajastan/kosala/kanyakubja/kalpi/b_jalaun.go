package kalpi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶罗郁那JalaunBarony struct {
	feud.BaseBarony
}

var BJalaun耶罗郁那 feud.Barony = &耶罗郁那JalaunBarony{}

func init() {
    f := BJalaun耶罗郁那.(*耶罗郁那JalaunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jalaun",
		TitleName: "耶罗郁那",
		TitleCode: "b_jalaun",
	}
}
