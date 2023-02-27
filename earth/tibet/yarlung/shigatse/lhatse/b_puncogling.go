package lhatse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彭措林PuncoglingBarony struct {
	feud.BaseBarony
}

var BPuncogling彭措林 feud.Barony = &彭措林PuncoglingBarony{}

func init() {
    f := BPuncogling彭措林.(*彭措林PuncoglingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "puncogling",
		TitleName: "彭措林",
		TitleCode: "b_puncogling",
	}
}
