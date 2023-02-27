package delta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法里斯库尔FareskurBarony struct {
	feud.BaseBarony
}

var BFareskur法里斯库尔 feud.Barony = &法里斯库尔FareskurBarony{}

func init() {
    f := BFareskur法里斯库尔.(*法里斯库尔FareskurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fareskur",
		TitleName: "法里斯库尔",
		TitleCode: "b_fareskur",
	}
}
