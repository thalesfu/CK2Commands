package hereford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫里福德HerefordBarony struct {
	feud.BaseBarony
}

var BHereford赫里福德 feud.Barony = &赫里福德HerefordBarony{}

func init() {
	f := BHereford赫里福德.(*赫里福德HerefordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hereford",
		TitleName: "赫里福德",
		TitleCode: "b_hereford",
	}
}
