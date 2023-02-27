package gwynedd

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/gwynedd/anglesey"
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/gwynedd/gwynedd"
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/gwynedd/perfeddwlad"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GwyneddDuke interface {
    feud.Duke
    CAnglesey安格尔西() 	anglesey.AngleseyCounty
    CGwynedd圭内思() 	gwynedd.GwyneddCounty
    CPerfeddwlad佩菲杜拉德() 	perfeddwlad.PerfeddwladCounty
}

type 圭内思GwyneddDuke struct {
	feud.BaseDuke
	安格尔西Anglesey 	anglesey.AngleseyCounty
	圭内思Gwynedd 	gwynedd.GwyneddCounty
	佩菲杜拉德Perfeddwlad 	perfeddwlad.PerfeddwladCounty
}

func (d *圭内思GwyneddDuke) CAnglesey安格尔西() anglesey.AngleseyCounty {
	return d.安格尔西Anglesey
}
    
func (d *圭内思GwyneddDuke) CGwynedd圭内思() gwynedd.GwyneddCounty {
	return d.圭内思Gwynedd
}
    
func (d *圭内思GwyneddDuke) CPerfeddwlad佩菲杜拉德() perfeddwlad.PerfeddwladCounty {
	return d.佩菲杜拉德Perfeddwlad
}
    
var DGwynedd圭内思 GwyneddDuke = &圭内思GwyneddDuke{}

func init() {
	f := DGwynedd圭内思.(*圭内思GwyneddDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "gwynedd",
		TitleName: "圭内思",
		TitleCode: "d_gwynedd",
		Counties:  map[string]feud.County{},
	}

	f.安格尔西Anglesey = anglesey.CAnglesey安格尔西
	f.安格尔西Anglesey.SetParent(f)
	
	f.圭内思Gwynedd = gwynedd.CGwynedd圭内思
	f.圭内思Gwynedd.SetParent(f)
	
	f.佩菲杜拉德Perfeddwlad = perfeddwlad.CPerfeddwlad佩菲杜拉德
	f.佩菲杜拉德Perfeddwlad.SetParent(f)
	
}
