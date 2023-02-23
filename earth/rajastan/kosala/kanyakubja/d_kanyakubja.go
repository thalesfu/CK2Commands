package kanyakubja

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/kosala/kanyakubja/asni"
	"github.com/thalesfu/CK2Commands/earth/rajastan/kosala/kanyakubja/bharauli"
	"github.com/thalesfu/CK2Commands/earth/rajastan/kosala/kanyakubja/kalpi"
	"github.com/thalesfu/CK2Commands/earth/rajastan/kosala/kanyakubja/kanyakubja"
	"github.com/thalesfu/CK2Commands/earth/rajastan/kosala/kanyakubja/lakhnau"
	"github.com/thalesfu/CK2Commands/earth/rajastan/kosala/kanyakubja/prayaga"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KanyakubjaDuke interface {
	feud.Duke
	CAsni阿斯尼() asni.AsniCounty
	CBharauli婆楼梨() bharauli.BharauliCounty
	CKalpi伽罗毗() kalpi.KalpiCounty
	CKanyakubja曲女城() kanyakubja.KanyakubjaCounty
	CLakhnau落伽奈() lakhnau.LakhnauCounty
	CPrayaga钵罗耶伽() prayaga.PrayagaCounty
}

type 曲女城KanyakubjaDuke struct {
	feud.BaseDuke
	阿斯尼Asni       asni.AsniCounty
	婆楼梨Bharauli   bharauli.BharauliCounty
	伽罗毗Kalpi      kalpi.KalpiCounty
	曲女城Kanyakubja kanyakubja.KanyakubjaCounty
	落伽奈Lakhnau    lakhnau.LakhnauCounty
	钵罗耶伽Prayaga   prayaga.PrayagaCounty
}

func (d *曲女城KanyakubjaDuke) CAsni阿斯尼() asni.AsniCounty {
	return d.阿斯尼Asni
}

func (d *曲女城KanyakubjaDuke) CBharauli婆楼梨() bharauli.BharauliCounty {
	return d.婆楼梨Bharauli
}

func (d *曲女城KanyakubjaDuke) CKalpi伽罗毗() kalpi.KalpiCounty {
	return d.伽罗毗Kalpi
}

func (d *曲女城KanyakubjaDuke) CKanyakubja曲女城() kanyakubja.KanyakubjaCounty {
	return d.曲女城Kanyakubja
}

func (d *曲女城KanyakubjaDuke) CLakhnau落伽奈() lakhnau.LakhnauCounty {
	return d.落伽奈Lakhnau
}

func (d *曲女城KanyakubjaDuke) CPrayaga钵罗耶伽() prayaga.PrayagaCounty {
	return d.钵罗耶伽Prayaga
}

var DKanyakubja曲女城 KanyakubjaDuke = &曲女城KanyakubjaDuke{}

func init() {
	f := DKanyakubja曲女城.(*曲女城KanyakubjaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kanyakubja",
		TitleName: "曲女城",
		TitleCode: "d_kanyakubja",
		Counties:  map[string]feud.County{},
	}

	f.阿斯尼Asni = asni.CAsni阿斯尼
	f.阿斯尼Asni.SetParent(f)

	f.婆楼梨Bharauli = bharauli.CBharauli婆楼梨
	f.婆楼梨Bharauli.SetParent(f)

	f.伽罗毗Kalpi = kalpi.CKalpi伽罗毗
	f.伽罗毗Kalpi.SetParent(f)

	f.曲女城Kanyakubja = kanyakubja.CKanyakubja曲女城
	f.曲女城Kanyakubja.SetParent(f)

	f.落伽奈Lakhnau = lakhnau.CLakhnau落伽奈
	f.落伽奈Lakhnau.SetParent(f)

	f.钵罗耶伽Prayaga = prayaga.CPrayaga钵罗耶伽
	f.钵罗耶伽Prayaga.SetParent(f)

}
