package poitiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙孔图尔MoncontourBarony struct {
	feud.BaseBarony
}

var BMoncontour蒙孔图尔 feud.Barony = &蒙孔图尔MoncontourBarony{}

func init() {
	f := BMoncontour蒙孔图尔.(*蒙孔图尔MoncontourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moncontour",
		TitleName: "蒙孔图尔",
		TitleCode: "b_moncontour",
	}
}
