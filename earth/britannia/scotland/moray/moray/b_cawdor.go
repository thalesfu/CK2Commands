package moray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考德CawdorBarony struct {
	feud.BaseBarony
}

var BCawdor考德 feud.Barony = &考德CawdorBarony{}

func init() {
    f := BCawdor考德.(*考德CawdorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cawdor",
		TitleName: "考德",
		TitleCode: "b_cawdor",
	}
}
