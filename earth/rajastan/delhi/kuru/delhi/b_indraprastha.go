package delhi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因陀罗钵罗萨他IndraprasthaBarony struct {
	feud.BaseBarony
}

var BIndraprastha因陀罗钵罗萨他 feud.Barony = &因陀罗钵罗萨他IndraprasthaBarony{}

func init() {
    f := BIndraprastha因陀罗钵罗萨他.(*因陀罗钵罗萨他IndraprasthaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "indraprastha",
		TitleName: "因陀罗钵罗萨他",
		TitleCode: "b_indraprastha",
	}
}
