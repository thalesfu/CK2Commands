package canda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加德奇罗利GadchiroliBarony struct {
	feud.BaseBarony
}

var BGadchiroli加德奇罗利 feud.Barony = &加德奇罗利GadchiroliBarony{}

func init() {
    f := BGadchiroli加德奇罗利.(*加德奇罗利GadchiroliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gadchiroli",
		TitleName: "加德奇罗利",
		TitleCode: "b_gadchiroli",
	}
}
