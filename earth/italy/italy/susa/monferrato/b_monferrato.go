package monferrato

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙费拉托MonferratoBarony struct {
	feud.BaseBarony
}

var BMonferrato蒙费拉托 feud.Barony = &蒙费拉托MonferratoBarony{}

func init() {
    f := BMonferrato蒙费拉托.(*蒙费拉托MonferratoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monferrato",
		TitleName: "蒙费拉托",
		TitleCode: "b_monferrato",
	}
}
