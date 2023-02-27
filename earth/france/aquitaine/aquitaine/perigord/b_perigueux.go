package perigord

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩里格PerigueuxBarony struct {
	feud.BaseBarony
}

var BPerigueux佩里格 feud.Barony = &佩里格PerigueuxBarony{}

func init() {
    f := BPerigueux佩里格.(*佩里格PerigueuxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "perigueux",
		TitleName: "佩里格",
		TitleCode: "b_perigueux",
	}
}
