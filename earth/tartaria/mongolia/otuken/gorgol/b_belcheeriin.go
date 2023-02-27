package gorgol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔彻林BelcheeriinBarony struct {
	feud.BaseBarony
}

var BBelcheeriin贝尔彻林 feud.Barony = &贝尔彻林BelcheeriinBarony{}

func init() {
    f := BBelcheeriin贝尔彻林.(*贝尔彻林BelcheeriinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belcheeriin",
		TitleName: "贝尔彻林",
		TitleCode: "b_belcheeriin",
	}
}
