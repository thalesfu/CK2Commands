package balkhash

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙舒拜ShashubayBarony struct {
	feud.BaseBarony
}

var BShashubay沙舒拜 feud.Barony = &沙舒拜ShashubayBarony{}

func init() {
    f := BShashubay沙舒拜.(*沙舒拜ShashubayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shashubay",
		TitleName: "沙舒拜",
		TitleCode: "b_shashubay",
	}
}
