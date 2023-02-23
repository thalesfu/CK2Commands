package ilam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代赫洛兰DehloranBarony struct {
	feud.BaseBarony
}

var BDehloran代赫洛兰 feud.Barony = &代赫洛兰DehloranBarony{}

func init() {
	f := BDehloran代赫洛兰.(*代赫洛兰DehloranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dehloran",
		TitleName: "代赫洛兰",
		TitleCode: "b_dehloran",
	}
}
