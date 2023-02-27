package thalner

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆摩尔BhamerBarony struct {
	feud.BaseBarony
}

var BBhamer婆摩尔 feud.Barony = &婆摩尔BhamerBarony{}

func init() {
    f := BBhamer婆摩尔.(*婆摩尔BhamerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhamer",
		TitleName: "婆摩尔",
		TitleCode: "b_bhamer",
	}
}
