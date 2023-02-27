package mongolia

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/abakan"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/altay"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/barkul"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/baygal"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/beshbaliq"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/gobi_altay"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/ikh_bogd"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/juyan"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/kara_khorum"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/khangai"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/otuken"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/uvs"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MongoliaKingdom interface {
    feud.Kingdom
    DAbakan阿巴坎() 	abakan.AbakanDuke
    DAltay阿尔泰() 	altay.AltayDuke
    DBarkul婆悉厥() 	barkul.BarkulDuke
    DBaygal贝加尔() 	baygal.BaygalDuke
    DBeshbaliq别失八里() 	beshbaliq.BeshbaliqDuke
    DGobi_altay戈壁阿尔泰() 	gobi_altay.Gobi_altayDuke
    DIkh_bogd也客孛黑答() 	ikh_bogd.Ikh_bogdDuke
    DJuyan居延() 	juyan.JuyanDuke
    DKara_khorum哈剌和林() 	kara_khorum.Kara_khorumDuke
    DKhangai杭海() 	khangai.KhangaiDuke
    DOtuken于都斤() 	otuken.OtukenDuke
    DUvs乌布苏() 	uvs.UvsDuke
}

type 蒙古MongoliaKingdom struct {
	feud.BaseKingdom
	阿巴坎Abakan 	abakan.AbakanDuke
	阿尔泰Altay 	altay.AltayDuke
	婆悉厥Barkul 	barkul.BarkulDuke
	贝加尔Baygal 	baygal.BaygalDuke
	别失八里Beshbaliq 	beshbaliq.BeshbaliqDuke
	戈壁阿尔泰Gobi_altay 	gobi_altay.Gobi_altayDuke
	也客孛黑答Ikh_bogd 	ikh_bogd.Ikh_bogdDuke
	居延Juyan 	juyan.JuyanDuke
	哈剌和林Kara_khorum 	kara_khorum.Kara_khorumDuke
	杭海Khangai 	khangai.KhangaiDuke
	于都斤Otuken 	otuken.OtukenDuke
	乌布苏Uvs 	uvs.UvsDuke
}

func (k *蒙古MongoliaKingdom) DAbakan阿巴坎() abakan.AbakanDuke {
	return k.阿巴坎Abakan
}
    
func (k *蒙古MongoliaKingdom) DAltay阿尔泰() altay.AltayDuke {
	return k.阿尔泰Altay
}
    
func (k *蒙古MongoliaKingdom) DBarkul婆悉厥() barkul.BarkulDuke {
	return k.婆悉厥Barkul
}
    
func (k *蒙古MongoliaKingdom) DBaygal贝加尔() baygal.BaygalDuke {
	return k.贝加尔Baygal
}
    
func (k *蒙古MongoliaKingdom) DBeshbaliq别失八里() beshbaliq.BeshbaliqDuke {
	return k.别失八里Beshbaliq
}
    
func (k *蒙古MongoliaKingdom) DGobi_altay戈壁阿尔泰() gobi_altay.Gobi_altayDuke {
	return k.戈壁阿尔泰Gobi_altay
}
    
func (k *蒙古MongoliaKingdom) DIkh_bogd也客孛黑答() ikh_bogd.Ikh_bogdDuke {
	return k.也客孛黑答Ikh_bogd
}
    
func (k *蒙古MongoliaKingdom) DJuyan居延() juyan.JuyanDuke {
	return k.居延Juyan
}
    
func (k *蒙古MongoliaKingdom) DKara_khorum哈剌和林() kara_khorum.Kara_khorumDuke {
	return k.哈剌和林Kara_khorum
}
    
func (k *蒙古MongoliaKingdom) DKhangai杭海() khangai.KhangaiDuke {
	return k.杭海Khangai
}
    
func (k *蒙古MongoliaKingdom) DOtuken于都斤() otuken.OtukenDuke {
	return k.于都斤Otuken
}
    
func (k *蒙古MongoliaKingdom) DUvs乌布苏() uvs.UvsDuke {
	return k.乌布苏Uvs
}
    
var KMongolia蒙古 MongoliaKingdom = &蒙古MongoliaKingdom{}

func init() {
	f := KMongolia蒙古.(*蒙古MongoliaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "mongolia",
		TitleName: "蒙古",
		TitleCode: "k_mongolia",
		Dukes:  map[string]feud.Duke{},
	}

	f.阿巴坎Abakan = abakan.DAbakan阿巴坎
	f.阿巴坎Abakan.SetParent(f)
	
	f.阿尔泰Altay = altay.DAltay阿尔泰
	f.阿尔泰Altay.SetParent(f)
	
	f.婆悉厥Barkul = barkul.DBarkul婆悉厥
	f.婆悉厥Barkul.SetParent(f)
	
	f.贝加尔Baygal = baygal.DBaygal贝加尔
	f.贝加尔Baygal.SetParent(f)
	
	f.别失八里Beshbaliq = beshbaliq.DBeshbaliq别失八里
	f.别失八里Beshbaliq.SetParent(f)
	
	f.戈壁阿尔泰Gobi_altay = gobi_altay.DGobi_altay戈壁阿尔泰
	f.戈壁阿尔泰Gobi_altay.SetParent(f)
	
	f.也客孛黑答Ikh_bogd = ikh_bogd.DIkh_bogd也客孛黑答
	f.也客孛黑答Ikh_bogd.SetParent(f)
	
	f.居延Juyan = juyan.DJuyan居延
	f.居延Juyan.SetParent(f)
	
	f.哈剌和林Kara_khorum = kara_khorum.DKara_khorum哈剌和林
	f.哈剌和林Kara_khorum.SetParent(f)
	
	f.杭海Khangai = khangai.DKhangai杭海
	f.杭海Khangai.SetParent(f)
	
	f.于都斤Otuken = otuken.DOtuken于都斤
	f.于都斤Otuken.SetParent(f)
	
	f.乌布苏Uvs = uvs.DUvs乌布苏
	f.乌布苏Uvs.SetParent(f)
	
}
