package capua

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/capua/capua"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/capua/napoli"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CapuaDuke interface {
	feud.Duke
	CCapua卡普阿() capua.CapuaCounty
	CNapoli那波利() napoli.NapoliCounty
}

type 卡普阿CapuaDuke struct {
	feud.BaseDuke
	卡普阿Capua  capua.CapuaCounty
	那波利Napoli napoli.NapoliCounty
}

func (d *卡普阿CapuaDuke) CCapua卡普阿() capua.CapuaCounty {
	return d.卡普阿Capua
}

func (d *卡普阿CapuaDuke) CNapoli那波利() napoli.NapoliCounty {
	return d.那波利Napoli
}

var DCapua卡普阿 CapuaDuke = &卡普阿CapuaDuke{}

func init() {
	f := DCapua卡普阿.(*卡普阿CapuaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "capua",
		TitleName: "卡普阿",
		TitleCode: "d_capua",
		Counties:  map[string]feud.County{},
	}

	f.卡普阿Capua = capua.CCapua卡普阿
	f.卡普阿Capua.SetParent(f)

	f.那波利Napoli = napoli.CNapoli那波利
	f.那波利Napoli.SetParent(f)

}
