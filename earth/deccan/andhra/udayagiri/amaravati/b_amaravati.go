package amaravati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿摩罗伐底AmaravatiBarony struct {
	feud.BaseBarony
}

var BAmaravati阿摩罗伐底 feud.Barony = &阿摩罗伐底AmaravatiBarony{}

func init() {
	f := BAmaravati阿摩罗伐底.(*阿摩罗伐底AmaravatiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amaravati",
		TitleName: "阿摩罗伐底",
		TitleCode: "b_amaravati",
	}
}
