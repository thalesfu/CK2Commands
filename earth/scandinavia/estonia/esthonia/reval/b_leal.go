package reval

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利胡拉LealBarony struct {
	feud.BaseBarony
}

var BLeal利胡拉 feud.Barony = &利胡拉LealBarony{}

func init() {
    f := BLeal利胡拉.(*利胡拉LealBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leal",
		TitleName: "利胡拉",
		TitleCode: "b_leal",
	}
}
