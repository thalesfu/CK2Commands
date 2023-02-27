package potapi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 劫吠梨钵胝那摩KaveripattinamBarony struct {
	feud.BaseBarony
}

var BKaveripattinam劫吠梨钵胝那摩 feud.Barony = &劫吠梨钵胝那摩KaveripattinamBarony{}

func init() {
    f := BKaveripattinam劫吠梨钵胝那摩.(*劫吠梨钵胝那摩KaveripattinamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaveripattinam",
		TitleName: "劫吠梨钵胝那摩",
		TitleCode: "b_kaveripattinam",
	}
}
