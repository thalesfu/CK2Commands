package vidin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯夫尔利格SrvljigBarony struct {
	feud.BaseBarony
}

var BSrvljig斯夫尔利格 feud.Barony = &斯夫尔利格SrvljigBarony{}

func init() {
	f := BSrvljig斯夫尔利格.(*斯夫尔利格SrvljigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "srvljig",
		TitleName: "斯夫尔利格",
		TitleCode: "b_srvljig",
	}
}
