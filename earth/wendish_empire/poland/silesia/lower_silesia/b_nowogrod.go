package lower_silesia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺沃格鲁德NowogrodBarony struct {
	feud.BaseBarony
}

var BNowogrod诺沃格鲁德 feud.Barony = &诺沃格鲁德NowogrodBarony{}

func init() {
    f := BNowogrod诺沃格鲁德.(*诺沃格鲁德NowogrodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nowogrod",
		TitleName: "诺沃格鲁德",
		TitleCode: "b_nowogrod",
	}
}
