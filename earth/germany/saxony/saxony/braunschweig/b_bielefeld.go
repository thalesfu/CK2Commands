package braunschweig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比勒费尔德BielefeldBarony struct {
	feud.BaseBarony
}

var BBielefeld比勒费尔德 feud.Barony = &比勒费尔德BielefeldBarony{}

func init() {
	f := BBielefeld比勒费尔德.(*比勒费尔德BielefeldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bielefeld",
		TitleName: "比勒费尔德",
		TitleCode: "b_bielefeld",
	}
}
