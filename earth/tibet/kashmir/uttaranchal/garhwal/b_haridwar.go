package garhwal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃梨堕罗HaridwarBarony struct {
	feud.BaseBarony
}

var BHaridwar诃梨堕罗 feud.Barony = &诃梨堕罗HaridwarBarony{}

func init() {
	f := BHaridwar诃梨堕罗.(*诃梨堕罗HaridwarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haridwar",
		TitleName: "诃梨堕罗",
		TitleCode: "b_haridwar",
	}
}
