package shewa

import (
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/shewa/ankober"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/shewa/antalo"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ShewaDuke interface {
	feud.Duke
	CAnkober安科贝尔() ankober.AnkoberCounty
	CAntalo安塔洛() antalo.AntaloCounty
}

type 绍阿ShewaDuke struct {
	feud.BaseDuke
	安科贝尔Ankober ankober.AnkoberCounty
	安塔洛Antalo   antalo.AntaloCounty
}

func (d *绍阿ShewaDuke) CAnkober安科贝尔() ankober.AnkoberCounty {
	return d.安科贝尔Ankober
}

func (d *绍阿ShewaDuke) CAntalo安塔洛() antalo.AntaloCounty {
	return d.安塔洛Antalo
}

var DShewa绍阿 ShewaDuke = &绍阿ShewaDuke{}

func init() {
	f := DShewa绍阿.(*绍阿ShewaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "shewa",
		TitleName: "绍阿",
		TitleCode: "d_shewa",
		Counties:  map[string]feud.County{},
	}

	f.安科贝尔Ankober = ankober.CAnkober安科贝尔
	f.安科贝尔Ankober.SetParent(f)

	f.安塔洛Antalo = antalo.CAntalo安塔洛
	f.安塔洛Antalo.SetParent(f)

}
