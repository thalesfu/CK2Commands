package poznanskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦乌奇WalczBarony struct {
	feud.BaseBarony
}

var BWalcz瓦乌奇 feud.Barony = &瓦乌奇WalczBarony{}

func init() {
    f := BWalcz瓦乌奇.(*瓦乌奇WalczBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "walcz",
		TitleName: "瓦乌奇",
		TitleCode: "b_walcz",
	}
}
