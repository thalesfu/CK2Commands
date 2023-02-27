package acalapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿摩罗伐底AmravatiBarony struct {
	feud.BaseBarony
}

var BAmravati阿摩罗伐底 feud.Barony = &阿摩罗伐底AmravatiBarony{}

func init() {
    f := BAmravati阿摩罗伐底.(*阿摩罗伐底AmravatiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amravati",
		TitleName: "阿摩罗伐底",
		TitleCode: "b_amravati",
	}
}
