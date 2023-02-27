package asirgarh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯梨姞利呬AsirgarhBarony struct {
	feud.BaseBarony
}

var BAsirgarh阿斯梨姞利呬 feud.Barony = &阿斯梨姞利呬AsirgarhBarony{}

func init() {
    f := BAsirgarh阿斯梨姞利呬.(*阿斯梨姞利呬AsirgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asirgarh",
		TitleName: "阿斯梨姞利呬",
		TitleCode: "b_asirgarh",
	}
}
