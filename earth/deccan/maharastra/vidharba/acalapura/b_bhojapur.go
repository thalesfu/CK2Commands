package acalapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菩阇补罗BhojapurBarony struct {
	feud.BaseBarony
}

var BBhojapur菩阇补罗 feud.Barony = &菩阇补罗BhojapurBarony{}

func init() {
    f := BBhojapur菩阇补罗.(*菩阇补罗BhojapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhojapur",
		TitleName: "菩阇补罗",
		TitleCode: "b_bhojapur",
	}
}
