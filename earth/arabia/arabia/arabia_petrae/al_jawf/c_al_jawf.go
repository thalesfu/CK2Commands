package al_jawf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Al_jawfCounty interface {
    feud.County
    BAl_jawf焦夫() 	feud.Barony
    BAladan阿丹() 	feud.Barony
    BDumat_al_jandal杜玛占达() 	feud.Barony
    BQurayyat古里牙() 	feud.Barony
    BRadifah拉迪法() 	feud.Barony
    BSakakah塞卡凯() 	feud.Barony
    BTuwayr图韦尔() 	feud.Barony
    BTyer图耶() 	feud.Barony
}

type 焦夫Al_jawfCounty struct {
	feud.BaseCounty
	焦夫Al_jawf 	feud.Barony
	阿丹Aladan 	feud.Barony
	杜玛占达Dumat_al_jandal 	feud.Barony
	古里牙Qurayyat 	feud.Barony
	拉迪法Radifah 	feud.Barony
	塞卡凯Sakakah 	feud.Barony
	图韦尔Tuwayr 	feud.Barony
	图耶Tyer 	feud.Barony
}

func (c *焦夫Al_jawfCounty) BAl_jawf焦夫() feud.Barony {
	return c.焦夫Al_jawf
}
    
func (c *焦夫Al_jawfCounty) BAladan阿丹() feud.Barony {
	return c.阿丹Aladan
}
    
func (c *焦夫Al_jawfCounty) BDumat_al_jandal杜玛占达() feud.Barony {
	return c.杜玛占达Dumat_al_jandal
}
    
func (c *焦夫Al_jawfCounty) BQurayyat古里牙() feud.Barony {
	return c.古里牙Qurayyat
}
    
func (c *焦夫Al_jawfCounty) BRadifah拉迪法() feud.Barony {
	return c.拉迪法Radifah
}
    
func (c *焦夫Al_jawfCounty) BSakakah塞卡凯() feud.Barony {
	return c.塞卡凯Sakakah
}
    
func (c *焦夫Al_jawfCounty) BTuwayr图韦尔() feud.Barony {
	return c.图韦尔Tuwayr
}
    
func (c *焦夫Al_jawfCounty) BTyer图耶() feud.Barony {
	return c.图耶Tyer
}
    
var CAl_jawf焦夫 Al_jawfCounty = &焦夫Al_jawfCounty{}

func init() {
	f := CAl_jawf焦夫.(*焦夫Al_jawfCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "717",
		Title:     "al_jawf",
		TitleName: "焦夫",
		TitleCode: "c_al_jawf",
		Baronies:  map[string]feud.Barony{},
	}

	f.焦夫Al_jawf = BAl_jawf焦夫
	f.焦夫Al_jawf.SetParent(f)
	
	f.阿丹Aladan = BAladan阿丹
	f.阿丹Aladan.SetParent(f)
	
	f.杜玛占达Dumat_al_jandal = BDumat_al_jandal杜玛占达
	f.杜玛占达Dumat_al_jandal.SetParent(f)
	
	f.古里牙Qurayyat = BQurayyat古里牙
	f.古里牙Qurayyat.SetParent(f)
	
	f.拉迪法Radifah = BRadifah拉迪法
	f.拉迪法Radifah.SetParent(f)
	
	f.塞卡凯Sakakah = BSakakah塞卡凯
	f.塞卡凯Sakakah.SetParent(f)
	
	f.图韦尔Tuwayr = BTuwayr图韦尔
	f.图韦尔Tuwayr.SetParent(f)
	
	f.图耶Tyer = BTyer图耶
	f.图耶Tyer.SetParent(f)
	
}
