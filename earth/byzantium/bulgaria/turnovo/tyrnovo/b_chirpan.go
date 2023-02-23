package tyrnovo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇尔潘ChirpanBarony struct {
	feud.BaseBarony
}

var BChirpan奇尔潘 feud.Barony = &奇尔潘ChirpanBarony{}

func init() {
	f := BChirpan奇尔潘.(*奇尔潘ChirpanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chirpan",
		TitleName: "奇尔潘",
		TitleCode: "b_chirpan",
	}
}
