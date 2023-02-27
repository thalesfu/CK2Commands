package viviers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧布纳AubenasBarony struct {
	feud.BaseBarony
}

var BAubenas欧布纳 feud.Barony = &欧布纳AubenasBarony{}

func init() {
    f := BAubenas欧布纳.(*欧布纳AubenasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aubenas",
		TitleName: "欧布纳",
		TitleCode: "b_aubenas",
	}
}
