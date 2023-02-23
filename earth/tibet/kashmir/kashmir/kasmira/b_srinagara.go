package kasmira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 室利那伽罗SrinagaraBarony struct {
	feud.BaseBarony
}

var BSrinagara室利那伽罗 feud.Barony = &室利那伽罗SrinagaraBarony{}

func init() {
	f := BSrinagara室利那伽罗.(*室利那伽罗SrinagaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "srinagara",
		TitleName: "室利那伽罗",
		TitleCode: "b_srinagara",
	}
}
