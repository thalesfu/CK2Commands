package varendra

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/varendra/madhupur"
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/varendra/pundravardhana"
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/varendra/suvarnagram"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VarendraDuke interface {
    feud.Duke
    CMadhupur摩度补罗() 	madhupur.MadhupurCounty
    CPundravardhana奔那伐弹那() 	pundravardhana.PundravardhanaCounty
    CSuvarnagram苏伐剌拿伽罗摩() 	suvarnagram.SuvarnagramCounty
}

type 婆连陀罗VarendraDuke struct {
	feud.BaseDuke
	摩度补罗Madhupur 	madhupur.MadhupurCounty
	奔那伐弹那Pundravardhana 	pundravardhana.PundravardhanaCounty
	苏伐剌拿伽罗摩Suvarnagram 	suvarnagram.SuvarnagramCounty
}

func (d *婆连陀罗VarendraDuke) CMadhupur摩度补罗() madhupur.MadhupurCounty {
	return d.摩度补罗Madhupur
}
    
func (d *婆连陀罗VarendraDuke) CPundravardhana奔那伐弹那() pundravardhana.PundravardhanaCounty {
	return d.奔那伐弹那Pundravardhana
}
    
func (d *婆连陀罗VarendraDuke) CSuvarnagram苏伐剌拿伽罗摩() suvarnagram.SuvarnagramCounty {
	return d.苏伐剌拿伽罗摩Suvarnagram
}
    
var DVarendra婆连陀罗 VarendraDuke = &婆连陀罗VarendraDuke{}

func init() {
	f := DVarendra婆连陀罗.(*婆连陀罗VarendraDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "varendra",
		TitleName: "婆连陀罗",
		TitleCode: "d_varendra",
		Counties:  map[string]feud.County{},
	}

	f.摩度补罗Madhupur = madhupur.CMadhupur摩度补罗
	f.摩度补罗Madhupur.SetParent(f)
	
	f.奔那伐弹那Pundravardhana = pundravardhana.CPundravardhana奔那伐弹那
	f.奔那伐弹那Pundravardhana.SetParent(f)
	
	f.苏伐剌拿伽罗摩Suvarnagram = suvarnagram.CSuvarnagram苏伐剌拿伽罗摩
	f.苏伐剌拿伽罗摩Suvarnagram.SetParent(f)
	
}
