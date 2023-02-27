package alqusair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古胡孙AbughusunBarony struct {
	feud.BaseBarony
}

var BAbughusun古胡孙 feud.Barony = &古胡孙AbughusunBarony{}

func init() {
    f := BAbughusun古胡孙.(*古胡孙AbughusunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abughusun",
		TitleName: "古胡孙",
		TitleCode: "b_abughusun",
	}
}
