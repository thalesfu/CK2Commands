package wallachia

import (
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/wallachia/calarasi"
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/wallachia/campulung"
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/wallachia/tirgoviste"
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/wallachia/turnu"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WallachiaDuke interface {
	feud.Duke
	CCalarasi克勒拉希() calarasi.CalarasiCounty
	CCampulung肯普隆格() campulung.CampulungCounty
	CTirgoviste特尔戈维什泰() tirgoviste.TirgovisteCounty
	CTurnu图尔努() turnu.TurnuCounty
}

type 蒙泰尼亚WallachiaDuke struct {
	feud.BaseDuke
	克勒拉希Calarasi     calarasi.CalarasiCounty
	肯普隆格Campulung    campulung.CampulungCounty
	特尔戈维什泰Tirgoviste tirgoviste.TirgovisteCounty
	图尔努Turnu         turnu.TurnuCounty
}

func (d *蒙泰尼亚WallachiaDuke) CCalarasi克勒拉希() calarasi.CalarasiCounty {
	return d.克勒拉希Calarasi
}

func (d *蒙泰尼亚WallachiaDuke) CCampulung肯普隆格() campulung.CampulungCounty {
	return d.肯普隆格Campulung
}

func (d *蒙泰尼亚WallachiaDuke) CTirgoviste特尔戈维什泰() tirgoviste.TirgovisteCounty {
	return d.特尔戈维什泰Tirgoviste
}

func (d *蒙泰尼亚WallachiaDuke) CTurnu图尔努() turnu.TurnuCounty {
	return d.图尔努Turnu
}

var DWallachia蒙泰尼亚 WallachiaDuke = &蒙泰尼亚WallachiaDuke{}

func init() {
	f := DWallachia蒙泰尼亚.(*蒙泰尼亚WallachiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "wallachia",
		TitleName: "蒙泰尼亚",
		TitleCode: "d_wallachia",
		Counties:  map[string]feud.County{},
	}

	f.克勒拉希Calarasi = calarasi.CCalarasi克勒拉希
	f.克勒拉希Calarasi.SetParent(f)

	f.肯普隆格Campulung = campulung.CCampulung肯普隆格
	f.肯普隆格Campulung.SetParent(f)

	f.特尔戈维什泰Tirgoviste = tirgoviste.CTirgoviste特尔戈维什泰
	f.特尔戈维什泰Tirgoviste.SetParent(f)

	f.图尔努Turnu = turnu.CTurnu图尔努
	f.图尔努Turnu.SetParent(f)

}
