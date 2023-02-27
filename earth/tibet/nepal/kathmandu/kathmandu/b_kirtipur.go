package kathmandu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诘帝补罗KirtipurBarony struct {
	feud.BaseBarony
}

var BKirtipur诘帝补罗 feud.Barony = &诘帝补罗KirtipurBarony{}

func init() {
    f := BKirtipur诘帝补罗.(*诘帝补罗KirtipurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirtipur",
		TitleName: "诘帝补罗",
		TitleCode: "b_kirtipur",
	}
}
