package thuringen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施马尔卡尔登SchmalkaldenBarony struct {
	feud.BaseBarony
}

var BSchmalkalden施马尔卡尔登 feud.Barony = &施马尔卡尔登SchmalkaldenBarony{}

func init() {
    f := BSchmalkalden施马尔卡尔登.(*施马尔卡尔登SchmalkaldenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schmalkalden",
		TitleName: "施马尔卡尔登",
		TitleCode: "b_schmalkalden",
	}
}
