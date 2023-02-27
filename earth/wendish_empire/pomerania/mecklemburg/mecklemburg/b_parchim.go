package mecklemburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕尔希姆ParchimBarony struct {
	feud.BaseBarony
}

var BParchim帕尔希姆 feud.Barony = &帕尔希姆ParchimBarony{}

func init() {
    f := BParchim帕尔希姆.(*帕尔希姆ParchimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "parchim",
		TitleName: "帕尔希姆",
		TitleCode: "b_parchim",
	}
}
