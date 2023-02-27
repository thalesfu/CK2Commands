package amaravati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿赫巴拉姆AhobalamBarony struct {
	feud.BaseBarony
}

var BAhobalam阿赫巴拉姆 feud.Barony = &阿赫巴拉姆AhobalamBarony{}

func init() {
    f := BAhobalam阿赫巴拉姆.(*阿赫巴拉姆AhobalamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ahobalam",
		TitleName: "阿赫巴拉姆",
		TitleCode: "b_ahobalam",
	}
}
