package pronsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PronskCounty interface {
    feud.County
    BGryaznoye格里亚兹诺耶() 	feud.Barony
    BLikharevskoye_gorodishch利哈列夫斯科耶戈罗季希() 	feud.Barony
    BMikhaylov米哈伊洛夫() 	feud.Barony
    BOktyabrsky十月镇() 	feud.Barony
    BPavelets帕韦列茨() 	feud.Barony
    BPronsk普龙斯克() 	feud.Barony
    BSkopin斯科平() 	feud.Barony
}

type 普龙斯克PronskCounty struct {
	feud.BaseCounty
	格里亚兹诺耶Gryaznoye 	feud.Barony
	利哈列夫斯科耶戈罗季希Likharevskoye_gorodishch 	feud.Barony
	米哈伊洛夫Mikhaylov 	feud.Barony
	十月镇Oktyabrsky 	feud.Barony
	帕韦列茨Pavelets 	feud.Barony
	普龙斯克Pronsk 	feud.Barony
	斯科平Skopin 	feud.Barony
}

func (c *普龙斯克PronskCounty) BGryaznoye格里亚兹诺耶() feud.Barony {
	return c.格里亚兹诺耶Gryaznoye
}
    
func (c *普龙斯克PronskCounty) BLikharevskoye_gorodishch利哈列夫斯科耶戈罗季希() feud.Barony {
	return c.利哈列夫斯科耶戈罗季希Likharevskoye_gorodishch
}
    
func (c *普龙斯克PronskCounty) BMikhaylov米哈伊洛夫() feud.Barony {
	return c.米哈伊洛夫Mikhaylov
}
    
func (c *普龙斯克PronskCounty) BOktyabrsky十月镇() feud.Barony {
	return c.十月镇Oktyabrsky
}
    
func (c *普龙斯克PronskCounty) BPavelets帕韦列茨() feud.Barony {
	return c.帕韦列茨Pavelets
}
    
func (c *普龙斯克PronskCounty) BPronsk普龙斯克() feud.Barony {
	return c.普龙斯克Pronsk
}
    
func (c *普龙斯克PronskCounty) BSkopin斯科平() feud.Barony {
	return c.斯科平Skopin
}
    
var CPronsk普龙斯克 PronskCounty = &普龙斯克PronskCounty{}

func init() {
	f := CPronsk普龙斯克.(*普龙斯克PronskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "577",
		Title:     "pronsk",
		TitleName: "普龙斯克",
		TitleCode: "c_pronsk",
		Baronies:  map[string]feud.Barony{},
	}

	f.格里亚兹诺耶Gryaznoye = BGryaznoye格里亚兹诺耶
	f.格里亚兹诺耶Gryaznoye.SetParent(f)
	
	f.利哈列夫斯科耶戈罗季希Likharevskoye_gorodishch = BLikharevskoye_gorodishch利哈列夫斯科耶戈罗季希
	f.利哈列夫斯科耶戈罗季希Likharevskoye_gorodishch.SetParent(f)
	
	f.米哈伊洛夫Mikhaylov = BMikhaylov米哈伊洛夫
	f.米哈伊洛夫Mikhaylov.SetParent(f)
	
	f.十月镇Oktyabrsky = BOktyabrsky十月镇
	f.十月镇Oktyabrsky.SetParent(f)
	
	f.帕韦列茨Pavelets = BPavelets帕韦列茨
	f.帕韦列茨Pavelets.SetParent(f)
	
	f.普龙斯克Pronsk = BPronsk普龙斯克
	f.普龙斯克Pronsk.SetParent(f)
	
	f.斯科平Skopin = BSkopin斯科平
	f.斯科平Skopin.SetParent(f)
	
}
