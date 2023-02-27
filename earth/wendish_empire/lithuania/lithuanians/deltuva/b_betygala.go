package deltuva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝蒂加拉BetygalaBarony struct {
	feud.BaseBarony
}

var BBetygala贝蒂加拉 feud.Barony = &贝蒂加拉BetygalaBarony{}

func init() {
    f := BBetygala贝蒂加拉.(*贝蒂加拉BetygalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "betygala",
		TitleName: "贝蒂加拉",
		TitleCode: "b_betygala",
	}
}
