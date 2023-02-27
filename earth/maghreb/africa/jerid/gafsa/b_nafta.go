package gafsa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳夫塔NaftaBarony struct {
	feud.BaseBarony
}

var BNafta纳夫塔 feud.Barony = &纳夫塔NaftaBarony{}

func init() {
    f := BNafta纳夫塔.(*纳夫塔NaftaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nafta",
		TitleName: "纳夫塔",
		TitleCode: "b_nafta",
	}
}
