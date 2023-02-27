package satyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罽罗多俱波KiratakupaBarony struct {
	feud.BaseBarony
}

var BKiratakupa罽罗多俱波 feud.Barony = &罽罗多俱波KiratakupaBarony{}

func init() {
    f := BKiratakupa罽罗多俱波.(*罽罗多俱波KiratakupaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiratakupa",
		TitleName: "罽罗多俱波",
		TitleCode: "b_kiratakupa",
	}
}
