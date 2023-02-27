package kurumba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KurumbaCounty interface {
    feud.County
    BAribinda阿里宾达() 	feud.Barony
    BDori多里() 	feud.Barony
    BDouentza杜安扎() 	feud.Barony
    BGourel_bassirou古雷尔巴西鲁() 	feud.Barony
    BKissi基西() 	feud.Barony
    BOuario_ouario瓦里奥瓦里奥() 	feud.Barony
    BSeriwala塞里瓦拉() 	feud.Barony
}

type 古伦巴KurumbaCounty struct {
	feud.BaseCounty
	阿里宾达Aribinda 	feud.Barony
	多里Dori 	feud.Barony
	杜安扎Douentza 	feud.Barony
	古雷尔巴西鲁Gourel_bassirou 	feud.Barony
	基西Kissi 	feud.Barony
	瓦里奥瓦里奥Ouario_ouario 	feud.Barony
	塞里瓦拉Seriwala 	feud.Barony
}

func (c *古伦巴KurumbaCounty) BAribinda阿里宾达() feud.Barony {
	return c.阿里宾达Aribinda
}
    
func (c *古伦巴KurumbaCounty) BDori多里() feud.Barony {
	return c.多里Dori
}
    
func (c *古伦巴KurumbaCounty) BDouentza杜安扎() feud.Barony {
	return c.杜安扎Douentza
}
    
func (c *古伦巴KurumbaCounty) BGourel_bassirou古雷尔巴西鲁() feud.Barony {
	return c.古雷尔巴西鲁Gourel_bassirou
}
    
func (c *古伦巴KurumbaCounty) BKissi基西() feud.Barony {
	return c.基西Kissi
}
    
func (c *古伦巴KurumbaCounty) BOuario_ouario瓦里奥瓦里奥() feud.Barony {
	return c.瓦里奥瓦里奥Ouario_ouario
}
    
func (c *古伦巴KurumbaCounty) BSeriwala塞里瓦拉() feud.Barony {
	return c.塞里瓦拉Seriwala
}
    
var CKurumba古伦巴 KurumbaCounty = &古伦巴KurumbaCounty{}

func init() {
	f := CKurumba古伦巴.(*古伦巴KurumbaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1756",
		Title:     "kurumba",
		TitleName: "古伦巴",
		TitleCode: "c_kurumba",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿里宾达Aribinda = BAribinda阿里宾达
	f.阿里宾达Aribinda.SetParent(f)
	
	f.多里Dori = BDori多里
	f.多里Dori.SetParent(f)
	
	f.杜安扎Douentza = BDouentza杜安扎
	f.杜安扎Douentza.SetParent(f)
	
	f.古雷尔巴西鲁Gourel_bassirou = BGourel_bassirou古雷尔巴西鲁
	f.古雷尔巴西鲁Gourel_bassirou.SetParent(f)
	
	f.基西Kissi = BKissi基西
	f.基西Kissi.SetParent(f)
	
	f.瓦里奥瓦里奥Ouario_ouario = BOuario_ouario瓦里奥瓦里奥
	f.瓦里奥瓦里奥Ouario_ouario.SetParent(f)
	
	f.塞里瓦拉Seriwala = BSeriwala塞里瓦拉
	f.塞里瓦拉Seriwala.SetParent(f)
	
}
