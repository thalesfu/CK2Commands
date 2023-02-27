package gent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 登德尔蒙德DendermondeBarony struct {
	feud.BaseBarony
}

var BDendermonde登德尔蒙德 feud.Barony = &登德尔蒙德DendermondeBarony{}

func init() {
    f := BDendermonde登德尔蒙德.(*登德尔蒙德DendermondeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dendermonde",
		TitleName: "登德尔蒙德",
		TitleCode: "b_dendermonde",
	}
}
