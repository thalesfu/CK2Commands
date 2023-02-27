package nyitra

import (
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/nyitra/gemer"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/nyitra/nitra"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/nyitra/orava"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/nyitra/trencin"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NyitraDuke interface {
    feud.Duke
    CGemer格梅尔() 	gemer.GemerCounty
    CNitra尼特拉() 	nitra.NitraCounty
    COrava兹沃伦() 	orava.OravaCounty
    CTrencin特伦钦() 	trencin.TrencinCounty
}

type 尼特劳NyitraDuke struct {
	feud.BaseDuke
	格梅尔Gemer 	gemer.GemerCounty
	尼特拉Nitra 	nitra.NitraCounty
	兹沃伦Orava 	orava.OravaCounty
	特伦钦Trencin 	trencin.TrencinCounty
}

func (d *尼特劳NyitraDuke) CGemer格梅尔() gemer.GemerCounty {
	return d.格梅尔Gemer
}
    
func (d *尼特劳NyitraDuke) CNitra尼特拉() nitra.NitraCounty {
	return d.尼特拉Nitra
}
    
func (d *尼特劳NyitraDuke) COrava兹沃伦() orava.OravaCounty {
	return d.兹沃伦Orava
}
    
func (d *尼特劳NyitraDuke) CTrencin特伦钦() trencin.TrencinCounty {
	return d.特伦钦Trencin
}
    
var DNyitra尼特劳 NyitraDuke = &尼特劳NyitraDuke{}

func init() {
	f := DNyitra尼特劳.(*尼特劳NyitraDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "nyitra",
		TitleName: "尼特劳",
		TitleCode: "d_nyitra",
		Counties:  map[string]feud.County{},
	}

	f.格梅尔Gemer = gemer.CGemer格梅尔
	f.格梅尔Gemer.SetParent(f)
	
	f.尼特拉Nitra = nitra.CNitra尼特拉
	f.尼特拉Nitra.SetParent(f)
	
	f.兹沃伦Orava = orava.COrava兹沃伦
	f.兹沃伦Orava.SetParent(f)
	
	f.特伦钦Trencin = trencin.CTrencin特伦钦
	f.特伦钦Trencin.SetParent(f)
	
}
