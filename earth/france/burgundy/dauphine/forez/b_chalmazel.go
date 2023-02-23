package forez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙尔马泽尔ChalmazelBarony struct {
	feud.BaseBarony
}

var BChalmazel沙尔马泽尔 feud.Barony = &沙尔马泽尔ChalmazelBarony{}

func init() {
	f := BChalmazel沙尔马泽尔.(*沙尔马泽尔ChalmazelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chalmazel",
		TitleName: "沙尔马泽尔",
		TitleCode: "b_chalmazel",
	}
}
