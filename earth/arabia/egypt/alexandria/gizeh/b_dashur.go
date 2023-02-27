package gizeh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代赫舒尔DashurBarony struct {
	feud.BaseBarony
}

var BDashur代赫舒尔 feud.Barony = &代赫舒尔DashurBarony{}

func init() {
    f := BDashur代赫舒尔.(*代赫舒尔DashurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dashur",
		TitleName: "代赫舒尔",
		TitleCode: "b_dashur",
	}
}
