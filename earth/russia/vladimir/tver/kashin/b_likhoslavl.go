package kashin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利霍斯拉夫尔LikhoslavlBarony struct {
	feud.BaseBarony
}

var BLikhoslavl利霍斯拉夫尔 feud.Barony = &利霍斯拉夫尔LikhoslavlBarony{}

func init() {
    f := BLikhoslavl利霍斯拉夫尔.(*利霍斯拉夫尔LikhoslavlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "likhoslavl",
		TitleName: "利霍斯拉夫尔",
		TitleCode: "b_likhoslavl",
	}
}
