package kumara_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苾曼伽罗BetmangalaBarony struct {
	feud.BaseBarony
}

var BBetmangala苾曼伽罗 feud.Barony = &苾曼伽罗BetmangalaBarony{}

func init() {
    f := BBetmangala苾曼伽罗.(*苾曼伽罗BetmangalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "betmangala",
		TitleName: "苾曼伽罗",
		TitleCode: "b_betmangala",
	}
}
