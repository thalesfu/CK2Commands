package karashar

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/karashar/kara_khoja"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/karashar/karashar"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/karashar/kubera"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/karashar/kucha"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/karashar/luntai"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/karashar/mingoi"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KarasharDuke interface {
    feud.Duke
    CKara_khoja哈剌火州() 	kara_khoja.Kara_khojaCounty
    CKarashar喀喇沙尔() 	karashar.KarasharCounty
    CKubera俱毗罗() 	kubera.KuberaCounty
    CKucha苦叉() 	kucha.KuchaCounty
    CLuntai轮台() 	luntai.LuntaiCounty
    CMingoi明屋() 	mingoi.MingoiCounty
}

type 喀喇沙尔KarasharDuke struct {
	feud.BaseDuke
	哈剌火州Kara_khoja 	kara_khoja.Kara_khojaCounty
	喀喇沙尔Karashar 	karashar.KarasharCounty
	俱毗罗Kubera 	kubera.KuberaCounty
	苦叉Kucha 	kucha.KuchaCounty
	轮台Luntai 	luntai.LuntaiCounty
	明屋Mingoi 	mingoi.MingoiCounty
}

func (d *喀喇沙尔KarasharDuke) CKara_khoja哈剌火州() kara_khoja.Kara_khojaCounty {
	return d.哈剌火州Kara_khoja
}
    
func (d *喀喇沙尔KarasharDuke) CKarashar喀喇沙尔() karashar.KarasharCounty {
	return d.喀喇沙尔Karashar
}
    
func (d *喀喇沙尔KarasharDuke) CKubera俱毗罗() kubera.KuberaCounty {
	return d.俱毗罗Kubera
}
    
func (d *喀喇沙尔KarasharDuke) CKucha苦叉() kucha.KuchaCounty {
	return d.苦叉Kucha
}
    
func (d *喀喇沙尔KarasharDuke) CLuntai轮台() luntai.LuntaiCounty {
	return d.轮台Luntai
}
    
func (d *喀喇沙尔KarasharDuke) CMingoi明屋() mingoi.MingoiCounty {
	return d.明屋Mingoi
}
    
var DKarashar喀喇沙尔 KarasharDuke = &喀喇沙尔KarasharDuke{}

func init() {
	f := DKarashar喀喇沙尔.(*喀喇沙尔KarasharDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "karashar",
		TitleName: "喀喇沙尔",
		TitleCode: "d_karashar",
		Counties:  map[string]feud.County{},
	}

	f.哈剌火州Kara_khoja = kara_khoja.CKara_khoja哈剌火州
	f.哈剌火州Kara_khoja.SetParent(f)
	
	f.喀喇沙尔Karashar = karashar.CKarashar喀喇沙尔
	f.喀喇沙尔Karashar.SetParent(f)
	
	f.俱毗罗Kubera = kubera.CKubera俱毗罗
	f.俱毗罗Kubera.SetParent(f)
	
	f.苦叉Kucha = kucha.CKucha苦叉
	f.苦叉Kucha.SetParent(f)
	
	f.轮台Luntai = luntai.CLuntai轮台
	f.轮台Luntai.SetParent(f)
	
	f.明屋Mingoi = mingoi.CMingoi明屋
	f.明屋Mingoi.SetParent(f)
	
}
