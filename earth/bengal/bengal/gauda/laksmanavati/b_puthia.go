package laksmanavati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 补提阿PuthiaBarony struct {
	feud.BaseBarony
}

var BPuthia补提阿 feud.Barony = &补提阿PuthiaBarony{}

func init() {
    f := BPuthia补提阿.(*补提阿PuthiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "puthia",
		TitleName: "补提阿",
		TitleCode: "b_puthia",
	}
}
