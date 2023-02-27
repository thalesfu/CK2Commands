package neuchatel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉博讷维尔LabonnevilleBarony struct {
	feud.BaseBarony
}

var BLabonneville拉博讷维尔 feud.Barony = &拉博讷维尔LabonnevilleBarony{}

func init() {
    f := BLabonneville拉博讷维尔.(*拉博讷维尔LabonnevilleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "labonneville",
		TitleName: "拉博讷维尔",
		TitleCode: "b_labonneville",
	}
}
