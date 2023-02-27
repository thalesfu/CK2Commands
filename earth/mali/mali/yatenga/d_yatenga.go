package yatenga

import (
	"github.com/thalesfu/CK2Commands/earth/mali/mali/yatenga/bobo_dyulasso"
	"github.com/thalesfu/CK2Commands/earth/mali/mali/yatenga/wagadougou"
	"github.com/thalesfu/CK2Commands/earth/mali/mali/yatenga/yatenga"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YatengaDuke interface {
    feud.Duke
    CBobo_dyulasso博博迪乌拉索() 	bobo_dyulasso.Bobo_dyulassoCounty
    CWagadougou瓦加杜古() 	wagadougou.WagadougouCounty
    CYatenga亚滕加() 	yatenga.YatengaCounty
}

type 亚滕加YatengaDuke struct {
	feud.BaseDuke
	博博迪乌拉索Bobo_dyulasso 	bobo_dyulasso.Bobo_dyulassoCounty
	瓦加杜古Wagadougou 	wagadougou.WagadougouCounty
	亚滕加Yatenga 	yatenga.YatengaCounty
}

func (d *亚滕加YatengaDuke) CBobo_dyulasso博博迪乌拉索() bobo_dyulasso.Bobo_dyulassoCounty {
	return d.博博迪乌拉索Bobo_dyulasso
}
    
func (d *亚滕加YatengaDuke) CWagadougou瓦加杜古() wagadougou.WagadougouCounty {
	return d.瓦加杜古Wagadougou
}
    
func (d *亚滕加YatengaDuke) CYatenga亚滕加() yatenga.YatengaCounty {
	return d.亚滕加Yatenga
}
    
var DYatenga亚滕加 YatengaDuke = &亚滕加YatengaDuke{}

func init() {
	f := DYatenga亚滕加.(*亚滕加YatengaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "yatenga",
		TitleName: "亚滕加",
		TitleCode: "d_yatenga",
		Counties:  map[string]feud.County{},
	}

	f.博博迪乌拉索Bobo_dyulasso = bobo_dyulasso.CBobo_dyulasso博博迪乌拉索
	f.博博迪乌拉索Bobo_dyulasso.SetParent(f)
	
	f.瓦加杜古Wagadougou = wagadougou.CWagadougou瓦加杜古
	f.瓦加杜古Wagadougou.SetParent(f)
	
	f.亚滕加Yatenga = yatenga.CYatenga亚滕加
	f.亚滕加Yatenga.SetParent(f)
	
}
