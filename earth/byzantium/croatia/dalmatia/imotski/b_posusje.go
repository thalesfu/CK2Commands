package imotski

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波苏谢PosusjeBarony struct {
	feud.BaseBarony
}

var BPosusje波苏谢 feud.Barony = &波苏谢PosusjeBarony{}

func init() {
    f := BPosusje波苏谢.(*波苏谢PosusjeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "posusje",
		TitleName: "波苏谢",
		TitleCode: "b_posusje",
	}
}
