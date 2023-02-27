package xainza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班戈BaingoinBarony struct {
	feud.BaseBarony
}

var BBaingoin班戈 feud.Barony = &班戈BaingoinBarony{}

func init() {
    f := BBaingoin班戈.(*班戈BaingoinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baingoin",
		TitleName: "班戈",
		TitleCode: "b_baingoin",
	}
}
