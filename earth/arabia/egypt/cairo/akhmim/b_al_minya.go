package akhmim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 明亚Al_minyaBarony struct {
	feud.BaseBarony
}

var BAl_minya明亚 feud.Barony = &明亚Al_minyaBarony{}

func init() {
    f := BAl_minya明亚.(*明亚Al_minyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_minya",
		TitleName: "明亚",
		TitleCode: "b_al_minya",
	}
}
