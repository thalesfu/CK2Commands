package usturt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克_库都克AkkudukBarony struct {
	feud.BaseBarony
}

var BAkkuduk阿克_库都克 feud.Barony = &阿克_库都克AkkudukBarony{}

func init() {
    f := BAkkuduk阿克_库都克.(*阿克_库都克AkkudukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akkuduk",
		TitleName: "阿克_库都克",
		TitleCode: "b_akkuduk",
	}
}
