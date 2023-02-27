package bhera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝哈克马肯Behak_makenBarony struct {
	feud.BaseBarony
}

var BBehak_maken贝哈克马肯 feud.Barony = &贝哈克马肯Behak_makenBarony{}

func init() {
    f := BBehak_maken贝哈克马肯.(*贝哈克马肯Behak_makenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "behak_maken",
		TitleName: "贝哈克马肯",
		TitleCode: "b_behak_maken",
	}
}
