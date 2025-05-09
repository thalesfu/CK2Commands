package feudal

import (
	"bufio"
	"github.com/thalesfu/CK2Commands/earth"
	"github.com/thalesfu/CK2Commands/utils"
	"github.com/thalesfu/paradoxtools/CK2/feud"
	"log"
)

type Title struct {
	Feuds  []feud.Feud
	Holder int
}

func GiveTitle(titles []*Title) {
	file, err := utils.OpenFile("title")

	if err != nil {
		log.Println(err)
		return
	}

	defer utils.CloseFile(file)

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, title := range titles {
		for _, feud := range title.Feuds {
			writeGiveTitle(writer, title.Holder, feud)
		}
	}
}

func BuildTitle() {
	var titles []*Title
	titles = append(titles, &Title{
		Holder: 2749760,
		Feuds: []feud.Feud{

			//earth.Turkestan图兰.KKhotan于阗().DKashgar可失合儿().CAksu阿克苏(),
			//earth.Turkestan图兰.KKhotan于阗().DKashgar可失合儿().CUchturpan乌什吐鲁番(),
			//earth.Turkestan图兰.KKhotan于阗().DKashgar可失合儿().CUchturpan乌什吐鲁番().BUqturpan乌什(),
			//earth.Turkestan图兰.KKhotan于阗().DKashgar可失合儿().CKashgar可失合儿().BKashgar可失合儿(),
			//earth.Turkestan图兰.KKhotan于阗().DKashgar可失合儿().CYopurga岳普湖(),
			//earth.Turkestan图兰.KKhotan于阗().DKashgar可失合儿().CYopurga岳普湖().BYopurga岳普湖(),
			//earth.Turkestan图兰.KKhotan于阗().DKashgar可失合儿().CArtux阿图什(),
			//earth.Turkestan图兰.KKhotan于阗().DKashgar可失合儿().CArtux阿图什().BAtush阿图什(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CKhotan于阗(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CYarkand鸦儿看(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CKarghalik喀格勒克(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CKeriya克里雅(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CCadota精绝(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CCherchen车尔臣(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CCherchen车尔臣().BCherchen车尔臣(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CCherchen车尔臣().BZuomo左末(),
			//earth.Turkestan图兰.KKhotan于阗().DKumul哈密().CLopnor罗布泊(),
			//earth.Turkestan图兰.KKhotan于阗().DKumul哈密().CLopnor罗布泊().BLopnor罗布泊(),
			//earth.Turkestan图兰.KKhotan于阗().DKumul哈密().CLopnor罗布泊().BQitun七屯(),
			//earth.Turkestan图兰.KKhotan于阗().DKumul哈密().CCharkliq卡克里克().BCharkliq卡克里克(),
			//earth.Turkestan图兰.KKhotan于阗().DKumul哈密().CCharkliq卡克里克().BLop罗卜(),
			//earth.Turkestan图兰.KKhotan于阗().DKumul哈密().CCharkliq卡克里克().BMerdek麦德克(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CCherchen车尔臣(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CYarkand鸦儿看().BYarkand鸦儿看(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CKarghalik喀格勒克(),
			//earth.Turkestan图兰.KKhotan于阗().DKhotan于阗().CKarghalik喀格勒克().BKarghalik喀格勒克(),
			//earth.Turkestan图兰.KKhotan于阗().DKarashar喀喇沙尔().CKubera俱毗罗().BKubera俱毗罗(),
			//earth.Turkestan图兰.KKhiva河中().DKhiva希瓦(),
			//earth.Turkestan图兰.KKhiva河中().DKhiva希瓦().CUrgench马德米尼亚(),
			//earth.Turkestan图兰.KKhiva河中().DKhiva希瓦().CGurganj玉龙杰赤(),
			//earth.Turkestan图兰.KKhiva河中().DKhiva希瓦().CDashhowuz达什霍武兹().BDashhowuz达什霍武兹(),
			earth.Turkestan图兰.KKhiva河中().DKhuttal珂咄罗().CChaghaniyan斫汗那(),
			//earth.Turkestan图兰.KKhiva河中().DKhuttal珂咄罗().CVakhan镬侃(),
			//earth.Turkestan图兰.KKhiva河中().DFerghana费尔干纳().CKhaylam海拉姆(),
			//earth.Turkestan图兰.KKhiva河中().DFerghana费尔干纳(),
			//earth.Turkestan图兰.KKhiva河中().DFerghana费尔干纳().CFergana费尔干纳(),
			//earth.Turkestan图兰.KKhiva河中().DFerghana费尔干纳().CFergana费尔干纳().BUzkand讹迹邗(),
			//earth.Turkestan图兰.KKhiva河中().DFerghana费尔干纳().CKhaylam海拉姆(),
			//earth.Turkestan图兰.KKhiva河中().DFerghana费尔干纳().CKhojand苦盏(),
			//earth.Turkestan图兰.KKhiva河中().DSamarkand撒马尔罕().CBukhara布哈拉(),
			//earth.Turkestan图兰.KZhetysu七河().DChuy垂河(),
			//earth.Turkestan图兰.KZhetysu七河().DChuy垂河().CTalas怛罗斯(),
			//earth.Turkestan图兰.KZhetysu七河().DChuy垂河().CTalas怛罗斯().BShelji希勒吉(),
			//earth.Turkestan图兰.KZhetysu七河().DChuy垂河().CChuy裴罗将军城(),
			//earth.Turkestan图兰.KZhetysu七河().DChuy垂河().CChuy裴罗将军城().BBalasagun裴罗将军城(),
			//earth.Turkestan图兰.KZhetysu七河().DChuy垂河().CChuy裴罗将军城().BBishkek比什凯克(),
			//earth.Turkestan图兰.KZhetysu七河().DChuy垂河().CBarskhan拔塞干(),
			//earth.Turkestan图兰.KZhetysu七河().DChuy垂河().CBarskhan拔塞干().BBarskhan拔塞干(),
			//earth.Turkestan图兰.KZhetysu七河().DIli伊丽().CIli伊丽(),
			//earth.Turkestan图兰.KZhetysu七河().DIli伊丽().CAlmaliq阿力麻里(),
			//earth.Turkestan图兰.KTurkestan乌古斯().DUsturt乌斯秋尔特(),
			//earth.Turkestan图兰.KTurkestan乌古斯().DUsturt乌斯秋尔特().CKusbulak库斯布拉克().BKusbulak库斯布拉克(),
			//earth.Tibet吐蕃.KKashmir迦湿弥罗().DPamir播密().CPamir播密().BKala_panja喀喇喷赤(),
			//earth.Tibet吐蕃.KKashmir迦湿弥罗().DPamir播密().CTashkurgan塔什库尔干(),
			//earth.Tibet吐蕃.KKashmir迦湿弥罗().DKashmir迦湿弥罗().CGilgit吉尔吉特(),
			//earth.Tibet吐蕃.KKashmir迦湿弥罗().DKashmir迦湿弥罗().CGilgit吉尔吉特().BMinawar弥那伐尔(),
			//earth.Persia波斯帝国.KPersia波斯().DMafaza玛法扎(),
			//earth.Persia波斯帝国.KPersia波斯().DMafaza玛法扎().CLut玛法扎(),
			//earth.Persia波斯帝国.KPersia波斯().DMafaza玛法扎().CLut玛法扎().BDayhouk代胡克(),
			//earth.Persia波斯帝国.KPersia波斯().DMafaza玛法扎().CLut玛法扎().BTabas塔巴斯(),
			//earth.Persia波斯帝国.KPersia波斯().DMafaza玛法扎().CYazd亚兹德(),
			//earth.Persia波斯帝国.KPersia波斯().DMafaza玛法扎().CYazd亚兹德().BYazd亚兹德(),
			//earth.Persia波斯帝国.KPersia波斯().DMafaza玛法扎().CYazd亚兹德().BArdakan阿尔达坎(),
			//earth.Persia波斯帝国.KPersia波斯().DMafaza玛法扎().CYazd亚兹德().BMeybod梅博德(),
			//earth.Persia波斯帝国.KKhorasan呼罗珊().DBalkh缚喝().CBalkh缚喝(),
			//earth.Persia波斯帝国.KKhorasan呼罗珊().DBalkh缚喝().CBalkh缚喝().BTiliatepe蒂拉丘地(),
			//earth.Persia波斯帝国.KKhorasan呼罗珊().DBalkh缚喝().CBalkh缚喝().BAlkhanoum艾哈努姆(),
			//earth.Persia波斯帝国.KKhorasan呼罗珊(),
			//earth.Persia波斯帝国.KDaylam德莱木().DDihistan大益斯坦().CKara_kum法拉瓦().BFarava法拉瓦(),
			//earth.Persia波斯帝国.KDaylam德莱木().DDihistan大益斯坦().CKara_kum法拉瓦(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DBhakkar珀格尔(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DBhakkar珀格尔().CKandail甘代尔(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DBhakkar珀格尔().CKandail甘代尔().BKandail甘代尔(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DBhakkar珀格尔().CBhakkar珀格尔(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DBhakkar珀格尔().CAror阿卢梨(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DBhakkar珀格尔().CSibi尸毗(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DBhakkar珀格尔().CSibi尸毗().BSibi尸毗(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DBhakkar珀格尔().CQuzdar古兹达尔().BQuzdar古兹达尔(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DBhakkar珀格尔().CAror阿卢梨().BAror阿卢梨(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DSauvira粟毗罗().CDebul提部罗(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DSauvira粟毗罗().CMansura曼苏拉(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DSauvira粟毗罗().CMansura曼苏拉().BNerunkot内兰乔特(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DSauvira粟毗罗().CMansura曼苏拉().BMansura曼苏拉(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DSauvira粟毗罗().CSonda孙陀(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DSauvira粟毗罗().CSonda孙陀().BSonda孙陀(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DSauvira粟毗罗().CSonda孙陀().BKaroonjar迦卢恩奢(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DSauvira粟毗罗().CRanikot兰尼科特(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DSauvira粟毗罗().CRanikot兰尼科特().BPaat波罗(),
			//earth.Rajastan罗阇萨傥那.KSindh信度().DSauvira粟毗罗().CSiwistan室毗湿檀那().BGuja古奢(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DGurjara_mandala瞿折罗曼荼罗().CKhetaka契吒迦(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DGurjara_mandala瞿折罗曼荼罗().CKhetaka契吒迦().BTarapur多罗城(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DGurjara_mandala瞿折罗曼荼罗().CKhetaka契吒迦().BDholka陀尔迦(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DGurjara_mandala瞿折罗曼荼罗().CKhetaka契吒迦().BKhetaka契吒迦(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DAnartta阿捺多(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DAnartta阿捺多().CDhamalpur陀摩罗补罗(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DAnartta阿捺多().CDhamalpur陀摩罗补罗().BDhamalpur陀摩罗补罗(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DAnartta阿捺多().CDhamalpur陀摩罗补罗().BMorvi摩尔比(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DAnartta阿捺多().CDhamalpur陀摩罗补罗().BLakhota罗诃陀(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DAnartta阿捺多().CDhamalpur陀摩罗补罗().BDhrol陀罗(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DAnartta阿捺多().CDhamalpur陀摩罗补罗().BKhasta迦多(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DAnartta阿捺多().CDvaraka堕罗迦(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DAnartta阿捺多().CDvaraka堕罗迦().BBhanvad般跋(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DAnartta阿捺多().CDvaraka堕罗迦().BLalpur那尔补罗(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DAnartta阿捺多().CDvaraka堕罗迦().BGomati戈摩蒂(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DAnartta阿捺多().CDvaraka堕罗迦().BDwarakadheesh堕罗迦地舍(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DAnartta阿捺多().CKutch契吒(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DAnartta阿捺多().CKutch契吒().BDhaneti陀内提(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DAnartta阿捺多().CKutch契吒().BAnjar阿迦尔(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘().CValabhi伐腊毗(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘().CValabhi伐腊毗().BSihor锡霍尔(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘().CValabhi伐腊毗().BShatrunjaya设咄路阇耶(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘().CValabhi伐腊毗().BGundigar古蒂迦(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘().CValabhi伐腊毗().BPalatina波利旦那(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘().CSomnath须门那(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘().CSomnath须门那().BVeraval韦拉沃尔(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘().CSomnath须门那().BDiu刁元(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘().CSomnath须门那().BDelvada提伐陀(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘().CSomnath须门那().BSomnath苏摩那他(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘().CBhumilka菩弥伽(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘().CBhumilka菩弥伽().BBhavnat般那特(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘().CBhumilka菩弥伽().BUperkot优波罗拘吒(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘().CBhumilka菩弥伽().BGirnar吉尔纳尔(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘().CBhumilka菩弥伽().BBhumilka菩弥伽(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘().CVardhamana筏陀摩那(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DSaurashtra苏剌侘().CVardhamana筏陀摩那().BAmarvalli阿摩伐梨(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DLata罗吒(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DLata罗吒().CVadodara婆度陀罗(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DLata罗吒().CVadodara婆度陀罗().BKayavarohan迦耶伐卢汉拿(),
			//earth.Rajastan罗阇萨傥那.KGujarat瞿折罗().DLata罗吒().CVadodara婆度陀罗().BAnkotakka安拘吒迦(),
			//earth.Rajastan罗阇萨傥那.KRajputana罗阇弗多那(),
			//earth.Rajastan罗阇萨傥那.KRajputana罗阇弗多那().DStravani萨怛罗婆尼(),
			//earth.Rajastan罗阇萨傥那.KRajputana罗阇弗多那().DStravani萨怛罗婆尼().CLudrava律陀罗婆(),
			//earth.Rajastan罗阇萨傥那.KRajputana罗阇弗多那().DStravani萨怛罗婆尼().CLudrava律陀罗婆().BLudrava律陀罗婆(),
		},
	})
	titles = append(titles, &Title{
		Holder: 2740856, //yin
		Feuds: []feud.Feud{
			earth.Persia波斯帝国.KPersia波斯().DMafaza玛法扎().CLut玛法扎(),
		},
	})
	//titles = append(titles, &Title{
	//	Holder: 2729898, //zhang
	//	Feuds: []feud.Feud{
	//		earth.Persia波斯帝国.KAfghanistan迦布罗().DZabulistan社护罗萨他那().CGhazna伽色尼(),
	//		//earth.Persia波斯帝国.KPersia波斯().DMafaza玛法扎().CYazd亚兹德().BMeybod梅博德(),
	//	},
	//})
	//titles = append(titles, &Title{
	//	Holder: 2688515,
	//	Feuds: []feud.Feud{
	//		earth.Persia波斯帝国.KAfghanistan迦布罗().DKabul迦布罗().CKunduz昆都士(),
	//	},
	//})
	GiveTitle(titles)
}
