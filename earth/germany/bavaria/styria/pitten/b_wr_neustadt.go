package pitten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维也纳新城Wr_neustadtBarony struct {
	feud.BaseBarony
}

var BWr_neustadt维也纳新城 feud.Barony = &维也纳新城Wr_neustadtBarony{}

func init() {
    f := BWr_neustadt维也纳新城.(*维也纳新城Wr_neustadtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wr_neustadt",
		TitleName: "维也纳新城",
		TitleCode: "b_wr_neustadt",
	}
}
