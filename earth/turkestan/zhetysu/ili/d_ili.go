package ili

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/zhetysu/ili/almaliq"
	"github.com/thalesfu/CK2Commands/earth/turkestan/zhetysu/ili/ili"
	"github.com/thalesfu/CK2Commands/earth/turkestan/zhetysu/ili/qayaliq"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IliDuke interface {
	feud.Duke
	CAlmaliq阿力麻里() almaliq.AlmaliqCounty
	CIli伊丽() ili.IliCounty
	CQayaliq海押立() qayaliq.QayaliqCounty
}

type 伊丽IliDuke struct {
	feud.BaseDuke
	阿力麻里Almaliq almaliq.AlmaliqCounty
	伊丽Ili       ili.IliCounty
	海押立Qayaliq  qayaliq.QayaliqCounty
}

func (d *伊丽IliDuke) CAlmaliq阿力麻里() almaliq.AlmaliqCounty {
	return d.阿力麻里Almaliq
}

func (d *伊丽IliDuke) CIli伊丽() ili.IliCounty {
	return d.伊丽Ili
}

func (d *伊丽IliDuke) CQayaliq海押立() qayaliq.QayaliqCounty {
	return d.海押立Qayaliq
}

var DIli伊丽 IliDuke = &伊丽IliDuke{}

func init() {
	f := DIli伊丽.(*伊丽IliDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ili",
		TitleName: "伊丽",
		TitleCode: "d_ili",
		Counties:  map[string]feud.County{},
	}

	f.阿力麻里Almaliq = almaliq.CAlmaliq阿力麻里
	f.阿力麻里Almaliq.SetParent(f)

	f.伊丽Ili = ili.CIli伊丽
	f.伊丽Ili.SetParent(f)

	f.海押立Qayaliq = qayaliq.CQayaliq海押立
	f.海押立Qayaliq.SetParent(f)

}
