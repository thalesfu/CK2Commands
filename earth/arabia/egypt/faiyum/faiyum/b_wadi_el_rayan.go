package faiyum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷延谷Wadi_el_rayanBarony struct {
	feud.BaseBarony
}

var BWadi_el_rayan雷延谷 feud.Barony = &雷延谷Wadi_el_rayanBarony{}

func init() {
    f := BWadi_el_rayan雷延谷.(*雷延谷Wadi_el_rayanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wadi_el_rayan",
		TitleName: "雷延谷",
		TitleCode: "b_wadi_el_rayan",
	}
}
