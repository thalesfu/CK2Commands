package krain

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥埃尔斯佩格AuerspergBarony struct {
	feud.BaseBarony
}

var BAuersperg奥埃尔斯佩格 feud.Barony = &奥埃尔斯佩格AuerspergBarony{}

func init() {
    f := BAuersperg奥埃尔斯佩格.(*奥埃尔斯佩格AuerspergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "auersperg",
		TitleName: "奥埃尔斯佩格",
		TitleCode: "b_auersperg",
	}
}
