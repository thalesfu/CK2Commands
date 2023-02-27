package borovichi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普切夫扎PchevzhaBarony struct {
	feud.BaseBarony
}

var BPchevzha普切夫扎 feud.Barony = &普切夫扎PchevzhaBarony{}

func init() {
    f := BPchevzha普切夫扎.(*普切夫扎PchevzhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pchevzha",
		TitleName: "普切夫扎",
		TitleCode: "b_pchevzha",
	}
}
