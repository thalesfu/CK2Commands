package claudiopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比提尼乌姆BithyniumBarony struct {
	feud.BaseBarony
}

var BBithynium比提尼乌姆 feud.Barony = &比提尼乌姆BithyniumBarony{}

func init() {
    f := BBithynium比提尼乌姆.(*比提尼乌姆BithyniumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bithynium",
		TitleName: "比提尼乌姆",
		TitleCode: "b_bithynium",
	}
}
