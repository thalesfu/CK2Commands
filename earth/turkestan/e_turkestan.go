package turkestan

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan"
	"github.com/thalesfu/CK2Commands/earth/turkestan/turkestan"
	"github.com/thalesfu/CK2Commands/earth/turkestan/zhetysu"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TurkestanEmpire interface {
	feud.Empire
	KKhiva河中() khiva.KhivaKingdom
	KKhotan于阗() khotan.KhotanKingdom
	KTurkestan乌古斯() turkestan.TurkestanKingdom
	KZhetysu七河() zhetysu.ZhetysuKingdom
}

type 图兰TurkestanEmpire struct {
	feud.BaseEmpire
	河中Khiva      khiva.KhivaKingdom
	于阗Khotan     khotan.KhotanKingdom
	乌古斯Turkestan turkestan.TurkestanKingdom
	七河Zhetysu    zhetysu.ZhetysuKingdom
}

func (e *图兰TurkestanEmpire) KKhiva河中() khiva.KhivaKingdom {
	return e.河中Khiva
}

func (e *图兰TurkestanEmpire) KKhotan于阗() khotan.KhotanKingdom {
	return e.于阗Khotan
}

func (e *图兰TurkestanEmpire) KTurkestan乌古斯() turkestan.TurkestanKingdom {
	return e.乌古斯Turkestan
}

func (e *图兰TurkestanEmpire) KZhetysu七河() zhetysu.ZhetysuKingdom {
	return e.七河Zhetysu
}

var ETurkestan图兰 TurkestanEmpire = &图兰TurkestanEmpire{}

func init() {
	f := ETurkestan图兰.(*图兰TurkestanEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "turkestan",
		TitleName: "图兰",
		TitleCode: "e_turkestan",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.河中Khiva = khiva.KKhiva河中
	f.河中Khiva.SetParent(f)
	f.于阗Khotan = khotan.KKhotan于阗
	f.于阗Khotan.SetParent(f)
	f.乌古斯Turkestan = turkestan.KTurkestan乌古斯
	f.乌古斯Turkestan.SetParent(f)
	f.七河Zhetysu = zhetysu.KZhetysu七河
	f.七河Zhetysu.SetParent(f)
}
