package toropets

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托罗佩茨ToropetsBarony struct {
	feud.BaseBarony
}

var BToropets托罗佩茨 feud.Barony = &托罗佩茨ToropetsBarony{}

func init() {
    f := BToropets托罗佩茨.(*托罗佩茨ToropetsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "toropets",
		TitleName: "托罗佩茨",
		TitleCode: "b_toropets",
	}
}
