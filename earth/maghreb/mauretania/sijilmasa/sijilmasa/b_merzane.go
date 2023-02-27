package sijilmasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅尔赞MerzaneBarony struct {
	feud.BaseBarony
}

var BMerzane梅尔赞 feud.Barony = &梅尔赞MerzaneBarony{}

func init() {
    f := BMerzane梅尔赞.(*梅尔赞MerzaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "merzane",
		TitleName: "梅尔赞",
		TitleCode: "b_merzane",
	}
}
