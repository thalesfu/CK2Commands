package charkliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朱倭ZhehorBarony struct {
	feud.BaseBarony
}

var BZhehor朱倭 feud.Barony = &朱倭ZhehorBarony{}

func init() {
    f := BZhehor朱倭.(*朱倭ZhehorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhehor",
		TitleName: "朱倭",
		TitleCode: "b_zhehor",
	}
}
