package kara_bogaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢韦尔维赫SevervykhBarony struct {
	feud.BaseBarony
}

var BSevervykh谢韦尔维赫 feud.Barony = &谢韦尔维赫SevervykhBarony{}

func init() {
    f := BSevervykh谢韦尔维赫.(*谢韦尔维赫SevervykhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "severvykh",
		TitleName: "谢韦尔维赫",
		TitleCode: "b_severvykh",
	}
}
