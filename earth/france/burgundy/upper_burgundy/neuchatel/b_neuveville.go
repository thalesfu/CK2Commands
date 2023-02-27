package neuchatel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 讷维尔NeuvevilleBarony struct {
	feud.BaseBarony
}

var BNeuveville讷维尔 feud.Barony = &讷维尔NeuvevilleBarony{}

func init() {
    f := BNeuveville讷维尔.(*讷维尔NeuvevilleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neuveville",
		TitleName: "讷维尔",
		TitleCode: "b_neuveville",
	}
}
