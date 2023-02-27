package tell_atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅彻尔克MechreckBarony struct {
	feud.BaseBarony
}

var BMechreck梅彻尔克 feud.Barony = &梅彻尔克MechreckBarony{}

func init() {
    f := BMechreck梅彻尔克.(*梅彻尔克MechreckBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mechreck",
		TitleName: "梅彻尔克",
		TitleCode: "b_mechreck",
	}
}
