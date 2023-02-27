package tagadur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多伽头罗TagadurBarony struct {
	feud.BaseBarony
}

var BTagadur多伽头罗 feud.Barony = &多伽头罗TagadurBarony{}

func init() {
    f := BTagadur多伽头罗.(*多伽头罗TagadurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tagadur",
		TitleName: "多伽头罗",
		TitleCode: "b_tagadur",
	}
}
