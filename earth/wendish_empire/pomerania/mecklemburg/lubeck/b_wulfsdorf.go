package lubeck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伍尔夫斯多夫WulfsdorfBarony struct {
	feud.BaseBarony
}

var BWulfsdorf伍尔夫斯多夫 feud.Barony = &伍尔夫斯多夫WulfsdorfBarony{}

func init() {
    f := BWulfsdorf伍尔夫斯多夫.(*伍尔夫斯多夫WulfsdorfBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wulfsdorf",
		TitleName: "伍尔夫斯多夫",
		TitleCode: "b_wulfsdorf",
	}
}
