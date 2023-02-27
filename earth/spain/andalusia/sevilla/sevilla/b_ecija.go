package sevilla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃西哈EcijaBarony struct {
	feud.BaseBarony
}

var BEcija埃西哈 feud.Barony = &埃西哈EcijaBarony{}

func init() {
    f := BEcija埃西哈.(*埃西哈EcijaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ecija",
		TitleName: "埃西哈",
		TitleCode: "b_ecija",
	}
}
