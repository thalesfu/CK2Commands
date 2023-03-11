package ushytsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌希察UshytsiaBarony struct {
	feud.BaseBarony
}

var BUshytsia乌希察 feud.Barony = &乌希察UshytsiaBarony{}

func init() {
    f := BUshytsia乌希察.(*乌希察UshytsiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ushytsia",
		TitleName: "乌希察",
		TitleCode: "b_ushytsia",
	}
}
