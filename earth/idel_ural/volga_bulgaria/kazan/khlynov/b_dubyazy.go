package khlynov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜布亚济DubyazyBarony struct {
	feud.BaseBarony
}

var BDubyazy杜布亚济 feud.Barony = &杜布亚济DubyazyBarony{}

func init() {
    f := BDubyazy杜布亚济.(*杜布亚济DubyazyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dubyazy",
		TitleName: "杜布亚济",
		TitleCode: "b_dubyazy",
	}
}
