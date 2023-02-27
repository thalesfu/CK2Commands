package cephalonia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯法利尼亚KefaloniaBarony struct {
	feud.BaseBarony
}

var BKefalonia凯法利尼亚 feud.Barony = &凯法利尼亚KefaloniaBarony{}

func init() {
    f := BKefalonia凯法利尼亚.(*凯法利尼亚KefaloniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kefalonia",
		TitleName: "凯法利尼亚",
		TitleCode: "b_kefalonia",
	}
}
