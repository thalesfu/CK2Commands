package ubagan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朱马古利ZhumagulBarony struct {
	feud.BaseBarony
}

var BZhumagul朱马古利 feud.Barony = &朱马古利ZhumagulBarony{}

func init() {
    f := BZhumagul朱马古利.(*朱马古利ZhumagulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhumagul",
		TitleName: "朱马古利",
		TitleCode: "b_zhumagul",
	}
}
