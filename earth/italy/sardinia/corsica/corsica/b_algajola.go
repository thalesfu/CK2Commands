package corsica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔加约拉AlgajolaBarony struct {
	feud.BaseBarony
}

var BAlgajola阿尔加约拉 feud.Barony = &阿尔加约拉AlgajolaBarony{}

func init() {
	f := BAlgajola阿尔加约拉.(*阿尔加约拉AlgajolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "algajola",
		TitleName: "阿尔加约拉",
		TitleCode: "b_algajola",
	}
}
