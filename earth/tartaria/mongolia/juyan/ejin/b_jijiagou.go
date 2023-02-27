package ejin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纪家沟JijiagouBarony struct {
	feud.BaseBarony
}

var BJijiagou纪家沟 feud.Barony = &纪家沟JijiagouBarony{}

func init() {
    f := BJijiagou纪家沟.(*纪家沟JijiagouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jijiagou",
		TitleName: "纪家沟",
		TitleCode: "b_jijiagou",
	}
}
