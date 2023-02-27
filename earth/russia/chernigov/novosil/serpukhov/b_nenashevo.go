package serpukhov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅纳舍沃NenashevoBarony struct {
	feud.BaseBarony
}

var BNenashevo涅纳舍沃 feud.Barony = &涅纳舍沃NenashevoBarony{}

func init() {
    f := BNenashevo涅纳舍沃.(*涅纳舍沃NenashevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nenashevo",
		TitleName: "涅纳舍沃",
		TitleCode: "b_nenashevo",
	}
}
