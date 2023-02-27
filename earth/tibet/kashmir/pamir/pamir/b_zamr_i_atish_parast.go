package pamir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜火城Zamr_i_atish_parastBarony struct {
	feud.BaseBarony
}

var BZamr_i_atish_parast拜火城 feud.Barony = &拜火城Zamr_i_atish_parastBarony{}

func init() {
    f := BZamr_i_atish_parast拜火城.(*拜火城Zamr_i_atish_parastBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zamr_i_atish_parast",
		TitleName: "拜火城",
		TitleCode: "b_zamr_i_atish_parast",
	}
}
