package tana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫纳斯特尔斯基MonastyrskyBarony struct {
	feud.BaseBarony
}

var BMonastyrsky莫纳斯特尔斯基 feud.Barony = &莫纳斯特尔斯基MonastyrskyBarony{}

func init() {
    f := BMonastyrsky莫纳斯特尔斯基.(*莫纳斯特尔斯基MonastyrskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monastyrsky",
		TitleName: "莫纳斯特尔斯基",
		TitleCode: "b_monastyrsky",
	}
}
