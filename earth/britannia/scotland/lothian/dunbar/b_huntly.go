package dunbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亨特利HuntlyBarony struct {
	feud.BaseBarony
}

var BHuntly亨特利 feud.Barony = &亨特利HuntlyBarony{}

func init() {
    f := BHuntly亨特利.(*亨特利HuntlyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huntly",
		TitleName: "亨特利",
		TitleCode: "b_huntly",
	}
}
