package gwent

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/gwent/brycheiniog"
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/gwent/glamorgan"
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/gwent/gwent"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GwentDuke interface {
    feud.Duke
    CBrycheiniog布里黑尼奥格() 	brycheiniog.BrycheiniogCounty
    CGlamorgan格拉摩根() 	glamorgan.GlamorganCounty
    CGwent格温特() 	gwent.GwentCounty
}

type 格温特GwentDuke struct {
	feud.BaseDuke
	布里黑尼奥格Brycheiniog 	brycheiniog.BrycheiniogCounty
	格拉摩根Glamorgan 	glamorgan.GlamorganCounty
	格温特Gwent 	gwent.GwentCounty
}

func (d *格温特GwentDuke) CBrycheiniog布里黑尼奥格() brycheiniog.BrycheiniogCounty {
	return d.布里黑尼奥格Brycheiniog
}
    
func (d *格温特GwentDuke) CGlamorgan格拉摩根() glamorgan.GlamorganCounty {
	return d.格拉摩根Glamorgan
}
    
func (d *格温特GwentDuke) CGwent格温特() gwent.GwentCounty {
	return d.格温特Gwent
}
    
var DGwent格温特 GwentDuke = &格温特GwentDuke{}

func init() {
	f := DGwent格温特.(*格温特GwentDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "gwent",
		TitleName: "格温特",
		TitleCode: "d_gwent",
		Counties:  map[string]feud.County{},
	}

	f.布里黑尼奥格Brycheiniog = brycheiniog.CBrycheiniog布里黑尼奥格
	f.布里黑尼奥格Brycheiniog.SetParent(f)
	
	f.格拉摩根Glamorgan = glamorgan.CGlamorgan格拉摩根
	f.格拉摩根Glamorgan.SetParent(f)
	
	f.格温特Gwent = gwent.CGwent格温特
	f.格温特Gwent.SetParent(f)
	
}
