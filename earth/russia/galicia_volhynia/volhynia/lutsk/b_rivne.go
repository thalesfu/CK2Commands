package lutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗夫诺RivneBarony struct {
	feud.BaseBarony
}

var BRivne罗夫诺 feud.Barony = &罗夫诺RivneBarony{}

func init() {
    f := BRivne罗夫诺.(*罗夫诺RivneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rivne",
		TitleName: "罗夫诺",
		TitleCode: "b_rivne",
	}
}
