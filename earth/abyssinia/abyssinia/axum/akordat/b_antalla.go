package akordat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安塔拉AntallaBarony struct {
	feud.BaseBarony
}

var BAntalla安塔拉 feud.Barony = &安塔拉AntallaBarony{}

func init() {
    f := BAntalla安塔拉.(*安塔拉AntallaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "antalla",
		TitleName: "安塔拉",
		TitleCode: "b_antalla",
	}
}
