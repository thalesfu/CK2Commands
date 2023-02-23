package lopnor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 北地家墩BeidijiadunBarony struct {
	feud.BaseBarony
}

var BBeidijiadun北地家墩 feud.Barony = &北地家墩BeidijiadunBarony{}

func init() {
	f := BBeidijiadun北地家墩.(*北地家墩BeidijiadunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beidijiadun",
		TitleName: "北地家墩",
		TitleCode: "b_beidijiadun",
	}
}
