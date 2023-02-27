package tihama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布斯AbsBarony struct {
	feud.BaseBarony
}

var BAbs阿布斯 feud.Barony = &阿布斯AbsBarony{}

func init() {
    f := BAbs阿布斯.(*阿布斯AbsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abs",
		TitleName: "阿布斯",
		TitleCode: "b_abs",
	}
}
