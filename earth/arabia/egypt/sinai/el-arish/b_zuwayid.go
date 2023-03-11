package el-arish

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祖瓦德ZuwayidBarony struct {
	feud.BaseBarony
}

var BZuwayid祖瓦德 feud.Barony = &祖瓦德ZuwayidBarony{}

func init() {
    f := BZuwayid祖瓦德.(*祖瓦德ZuwayidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zuwayid",
		TitleName: "祖瓦德",
		TitleCode: "b_zuwayid",
	}
}
