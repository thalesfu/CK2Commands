package bornu

import (
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/bornu/bauchi"
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/bornu/bornu"
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/bornu/kagha"
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/bornu/makari"
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/bornu/masseniya"
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/bornu/munio"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BornuDuke interface {
	feud.Duke
	CBauchi包奇() bauchi.BauchiCounty
	CBornu博尔努() bornu.BornuCounty
	CKagha卡加() kagha.KaghaCounty
	CMakari马卡里() makari.MakariCounty
	CMasseniya马塞尼亚() masseniya.MasseniyaCounty
	CMunio穆尼奥() munio.MunioCounty
}

type 博尔努BornuDuke struct {
	feud.BaseDuke
	包奇Bauchi      bauchi.BauchiCounty
	博尔努Bornu      bornu.BornuCounty
	卡加Kagha       kagha.KaghaCounty
	马卡里Makari     makari.MakariCounty
	马塞尼亚Masseniya masseniya.MasseniyaCounty
	穆尼奥Munio      munio.MunioCounty
}

func (d *博尔努BornuDuke) CBauchi包奇() bauchi.BauchiCounty {
	return d.包奇Bauchi
}

func (d *博尔努BornuDuke) CBornu博尔努() bornu.BornuCounty {
	return d.博尔努Bornu
}

func (d *博尔努BornuDuke) CKagha卡加() kagha.KaghaCounty {
	return d.卡加Kagha
}

func (d *博尔努BornuDuke) CMakari马卡里() makari.MakariCounty {
	return d.马卡里Makari
}

func (d *博尔努BornuDuke) CMasseniya马塞尼亚() masseniya.MasseniyaCounty {
	return d.马塞尼亚Masseniya
}

func (d *博尔努BornuDuke) CMunio穆尼奥() munio.MunioCounty {
	return d.穆尼奥Munio
}

var DBornu博尔努 BornuDuke = &博尔努BornuDuke{}

func init() {
	f := DBornu博尔努.(*博尔努BornuDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "bornu",
		TitleName: "博尔努",
		TitleCode: "d_bornu",
		Counties:  map[string]feud.County{},
	}

	f.包奇Bauchi = bauchi.CBauchi包奇
	f.包奇Bauchi.SetParent(f)

	f.博尔努Bornu = bornu.CBornu博尔努
	f.博尔努Bornu.SetParent(f)

	f.卡加Kagha = kagha.CKagha卡加
	f.卡加Kagha.SetParent(f)

	f.马卡里Makari = makari.CMakari马卡里
	f.马卡里Makari.SetParent(f)

	f.马塞尼亚Masseniya = masseniya.CMasseniya马塞尼亚
	f.马塞尼亚Masseniya.SetParent(f)

	f.穆尼奥Munio = munio.CMunio穆尼奥
	f.穆尼奥Munio.SetParent(f)

}
