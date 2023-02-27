package travunia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TravuniaCounty interface {
    feud.County
    BBogorodice博戈罗迪采() 	feud.Barony
    BDracevica德拉切维察() 	feud.Barony
    BGripuli格里普利() 	feud.Barony
    BKotor科托尔() 	feud.Barony
    BSv_stefan圣斯特凡() 	feud.Barony
    BTravunia特拉武尼亚() 	feud.Barony
    BTrebinje特雷比涅() 	feud.Barony
}

type 特拉武尼亚TravuniaCounty struct {
	feud.BaseCounty
	博戈罗迪采Bogorodice 	feud.Barony
	德拉切维察Dracevica 	feud.Barony
	格里普利Gripuli 	feud.Barony
	科托尔Kotor 	feud.Barony
	圣斯特凡Sv_stefan 	feud.Barony
	特拉武尼亚Travunia 	feud.Barony
	特雷比涅Trebinje 	feud.Barony
}

func (c *特拉武尼亚TravuniaCounty) BBogorodice博戈罗迪采() feud.Barony {
	return c.博戈罗迪采Bogorodice
}
    
func (c *特拉武尼亚TravuniaCounty) BDracevica德拉切维察() feud.Barony {
	return c.德拉切维察Dracevica
}
    
func (c *特拉武尼亚TravuniaCounty) BGripuli格里普利() feud.Barony {
	return c.格里普利Gripuli
}
    
func (c *特拉武尼亚TravuniaCounty) BKotor科托尔() feud.Barony {
	return c.科托尔Kotor
}
    
func (c *特拉武尼亚TravuniaCounty) BSv_stefan圣斯特凡() feud.Barony {
	return c.圣斯特凡Sv_stefan
}
    
func (c *特拉武尼亚TravuniaCounty) BTravunia特拉武尼亚() feud.Barony {
	return c.特拉武尼亚Travunia
}
    
func (c *特拉武尼亚TravuniaCounty) BTrebinje特雷比涅() feud.Barony {
	return c.特雷比涅Trebinje
}
    
var CTravunia特拉武尼亚 TravuniaCounty = &特拉武尼亚TravuniaCounty{}

func init() {
	f := CTravunia特拉武尼亚.(*特拉武尼亚TravuniaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1972",
		Title:     "travunia",
		TitleName: "特拉武尼亚",
		TitleCode: "c_travunia",
		Baronies:  map[string]feud.Barony{},
	}

	f.博戈罗迪采Bogorodice = BBogorodice博戈罗迪采
	f.博戈罗迪采Bogorodice.SetParent(f)
	
	f.德拉切维察Dracevica = BDracevica德拉切维察
	f.德拉切维察Dracevica.SetParent(f)
	
	f.格里普利Gripuli = BGripuli格里普利
	f.格里普利Gripuli.SetParent(f)
	
	f.科托尔Kotor = BKotor科托尔
	f.科托尔Kotor.SetParent(f)
	
	f.圣斯特凡Sv_stefan = BSv_stefan圣斯特凡
	f.圣斯特凡Sv_stefan.SetParent(f)
	
	f.特拉武尼亚Travunia = BTravunia特拉武尼亚
	f.特拉武尼亚Travunia.SetParent(f)
	
	f.特雷比涅Trebinje = BTrebinje特雷比涅
	f.特雷比涅Trebinje.SetParent(f)
	
}
