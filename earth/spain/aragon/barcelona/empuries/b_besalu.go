package empuries

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝萨卢BesaluBarony struct {
	feud.BaseBarony
}

var BBesalu贝萨卢 feud.Barony = &贝萨卢BesaluBarony{}

func init() {
    f := BBesalu贝萨卢.(*贝萨卢BesaluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "besalu",
		TitleName: "贝萨卢",
		TitleCode: "b_besalu",
	}
}
