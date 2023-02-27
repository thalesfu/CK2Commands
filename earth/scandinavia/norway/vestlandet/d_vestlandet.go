package vestlandet

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/vestlandet/bergenshus"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/vestlandet/forde"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/vestlandet/romsdal"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/vestlandet/sogn"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/vestlandet/sunnmore"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VestlandetDuke interface {
    feud.Duke
    CBergenshus卑尔根胡斯() 	bergenshus.BergenshusCounty
    CForde弗勒() 	forde.FordeCounty
    CRomsdal鲁姆斯达尔() 	romsdal.RomsdalCounty
    CSogn松恩() 	sogn.SognCounty
    CSunnmore孙默勒() 	sunnmore.SunnmoreCounty
}

type 西挪威VestlandetDuke struct {
	feud.BaseDuke
	卑尔根胡斯Bergenshus 	bergenshus.BergenshusCounty
	弗勒Forde 	forde.FordeCounty
	鲁姆斯达尔Romsdal 	romsdal.RomsdalCounty
	松恩Sogn 	sogn.SognCounty
	孙默勒Sunnmore 	sunnmore.SunnmoreCounty
}

func (d *西挪威VestlandetDuke) CBergenshus卑尔根胡斯() bergenshus.BergenshusCounty {
	return d.卑尔根胡斯Bergenshus
}
    
func (d *西挪威VestlandetDuke) CForde弗勒() forde.FordeCounty {
	return d.弗勒Forde
}
    
func (d *西挪威VestlandetDuke) CRomsdal鲁姆斯达尔() romsdal.RomsdalCounty {
	return d.鲁姆斯达尔Romsdal
}
    
func (d *西挪威VestlandetDuke) CSogn松恩() sogn.SognCounty {
	return d.松恩Sogn
}
    
func (d *西挪威VestlandetDuke) CSunnmore孙默勒() sunnmore.SunnmoreCounty {
	return d.孙默勒Sunnmore
}
    
var DVestlandet西挪威 VestlandetDuke = &西挪威VestlandetDuke{}

func init() {
	f := DVestlandet西挪威.(*西挪威VestlandetDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "vestlandet",
		TitleName: "西挪威",
		TitleCode: "d_vestlandet",
		Counties:  map[string]feud.County{},
	}

	f.卑尔根胡斯Bergenshus = bergenshus.CBergenshus卑尔根胡斯
	f.卑尔根胡斯Bergenshus.SetParent(f)
	
	f.弗勒Forde = forde.CForde弗勒
	f.弗勒Forde.SetParent(f)
	
	f.鲁姆斯达尔Romsdal = romsdal.CRomsdal鲁姆斯达尔
	f.鲁姆斯达尔Romsdal.SetParent(f)
	
	f.松恩Sogn = sogn.CSogn松恩
	f.松恩Sogn.SetParent(f)
	
	f.孙默勒Sunnmore = sunnmore.CSunnmore孙默勒
	f.孙默勒Sunnmore.SetParent(f)
	
}
