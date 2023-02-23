package tyrone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓吉文DungivenBarony struct {
	feud.BaseBarony
}

var BDungiven邓吉文 feud.Barony = &邓吉文DungivenBarony{}

func init() {
	f := BDungiven邓吉文.(*邓吉文DungivenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dungiven",
		TitleName: "邓吉文",
		TitleCode: "b_dungiven",
	}
}
