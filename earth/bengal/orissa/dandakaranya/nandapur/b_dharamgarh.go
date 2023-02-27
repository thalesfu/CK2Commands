package nandapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达摩姞利呬DharamgarhBarony struct {
	feud.BaseBarony
}

var BDharamgarh达摩姞利呬 feud.Barony = &达摩姞利呬DharamgarhBarony{}

func init() {
    f := BDharamgarh达摩姞利呬.(*达摩姞利呬DharamgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dharamgarh",
		TitleName: "达摩姞利呬",
		TitleCode: "b_dharamgarh",
	}
}
