package draa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰纳奈曼泽Jenane_el_menzehBarony struct {
	feud.BaseBarony
}

var BJenane_el_menzeh杰纳奈曼泽 feud.Barony = &杰纳奈曼泽Jenane_el_menzehBarony{}

func init() {
    f := BJenane_el_menzeh杰纳奈曼泽.(*杰纳奈曼泽Jenane_el_menzehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jenane_el_menzeh",
		TitleName: "杰纳奈曼泽",
		TitleCode: "b_jenane_el_menzeh",
	}
}
