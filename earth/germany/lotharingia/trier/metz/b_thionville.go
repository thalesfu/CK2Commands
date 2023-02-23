package metz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂永维勒ThionvilleBarony struct {
	feud.BaseBarony
}

var BThionville蒂永维勒 feud.Barony = &蒂永维勒ThionvilleBarony{}

func init() {
	f := BThionville蒂永维勒.(*蒂永维勒ThionvilleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thionville",
		TitleName: "蒂永维勒",
		TitleCode: "b_thionville",
	}
}
