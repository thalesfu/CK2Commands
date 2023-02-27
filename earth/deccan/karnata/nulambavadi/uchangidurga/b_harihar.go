package uchangidurga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫里赫尔HariharBarony struct {
	feud.BaseBarony
}

var BHarihar赫里赫尔 feud.Barony = &赫里赫尔HariharBarony{}

func init() {
    f := BHarihar赫里赫尔.(*赫里赫尔HariharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harihar",
		TitleName: "赫里赫尔",
		TitleCode: "b_harihar",
	}
}
