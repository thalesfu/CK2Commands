package bilyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努尔拉特NurlatBarony struct {
	feud.BaseBarony
}

var BNurlat努尔拉特 feud.Barony = &努尔拉特NurlatBarony{}

func init() {
    f := BNurlat努尔拉特.(*努尔拉特NurlatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nurlat",
		TitleName: "努尔拉特",
		TitleCode: "b_nurlat",
	}
}
