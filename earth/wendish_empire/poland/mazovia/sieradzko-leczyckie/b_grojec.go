package sieradzko-leczyckie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格罗杰茨GrojecBarony struct {
	feud.BaseBarony
}

var BGrojec格罗杰茨 feud.Barony = &格罗杰茨GrojecBarony{}

func init() {
    f := BGrojec格罗杰茨.(*格罗杰茨GrojecBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grojec",
		TitleName: "格罗杰茨",
		TitleCode: "b_grojec",
	}
}
