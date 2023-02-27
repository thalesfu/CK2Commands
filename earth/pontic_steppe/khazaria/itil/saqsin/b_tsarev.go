package saqsin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 察列夫TsarevBarony struct {
	feud.BaseBarony
}

var BTsarev察列夫 feud.Barony = &察列夫TsarevBarony{}

func init() {
    f := BTsarev察列夫.(*察列夫TsarevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsarev",
		TitleName: "察列夫",
		TitleCode: "b_tsarev",
	}
}
