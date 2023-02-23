package kano

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂加罗TigaroBarony struct {
	feud.BaseBarony
}

var BTigaro蒂加罗 feud.Barony = &蒂加罗TigaroBarony{}

func init() {
	f := BTigaro蒂加罗.(*蒂加罗TigaroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tigaro",
		TitleName: "蒂加罗",
		TitleCode: "b_tigaro",
	}
}
