package upper_burgundy

import (
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/upper_burgundy/aargau"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/upper_burgundy/bern"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/upper_burgundy/neuchatel"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/upper_burgundy/vaud"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Upper_burgundyDuke interface {
    feud.Duke
    CAargau巴塞尔() 	aargau.AargauCounty
    CBern伯尔尼() 	bern.BernCounty
    CNeuchatel纳沙泰尔() 	neuchatel.NeuchatelCounty
    CVaud沃() 	vaud.VaudCounty
}

type 上勃艮第Upper_burgundyDuke struct {
	feud.BaseDuke
	巴塞尔Aargau 	aargau.AargauCounty
	伯尔尼Bern 	bern.BernCounty
	纳沙泰尔Neuchatel 	neuchatel.NeuchatelCounty
	沃Vaud 	vaud.VaudCounty
}

func (d *上勃艮第Upper_burgundyDuke) CAargau巴塞尔() aargau.AargauCounty {
	return d.巴塞尔Aargau
}
    
func (d *上勃艮第Upper_burgundyDuke) CBern伯尔尼() bern.BernCounty {
	return d.伯尔尼Bern
}
    
func (d *上勃艮第Upper_burgundyDuke) CNeuchatel纳沙泰尔() neuchatel.NeuchatelCounty {
	return d.纳沙泰尔Neuchatel
}
    
func (d *上勃艮第Upper_burgundyDuke) CVaud沃() vaud.VaudCounty {
	return d.沃Vaud
}
    
var DUpper_burgundy上勃艮第 Upper_burgundyDuke = &上勃艮第Upper_burgundyDuke{}

func init() {
	f := DUpper_burgundy上勃艮第.(*上勃艮第Upper_burgundyDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "upper_burgundy",
		TitleName: "上勃艮第",
		TitleCode: "d_upper_burgundy",
		Counties:  map[string]feud.County{},
	}

	f.巴塞尔Aargau = aargau.CAargau巴塞尔
	f.巴塞尔Aargau.SetParent(f)
	
	f.伯尔尼Bern = bern.CBern伯尔尼
	f.伯尔尼Bern.SetParent(f)
	
	f.纳沙泰尔Neuchatel = neuchatel.CNeuchatel纳沙泰尔
	f.纳沙泰尔Neuchatel.SetParent(f)
	
	f.沃Vaud = vaud.CVaud沃
	f.沃Vaud.SetParent(f)
	
}
