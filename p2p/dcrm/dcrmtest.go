package dcrm

import (
	"fmt"
	//get args
	"flag"
	"github.com/fusion/go-fusion/log"
	"math/big"
	"os"
	"time"
)

//call define
func call(msg interface{}) {
	fmt.Printf("\n\ncall: msg=%v\n", msg)
}

func dcrmcall(msg interface{}) (<-chan string) {
	ch := make(chan string, 8000)
	fmt.Printf("\ndcrmcall: msg=%v\n", msg)
	ch <- "jklaafd"
	return ch
}

func StartTest() {
	fmt.Printf("\n\nDCRM P2P test ...\n\n")
	//callback
	RegisterCallback(call)
	RegisterDcrmCallback(dcrmcall)

	verbosity := flag.Int("verbosity", int(log.LvlInfo), "log verbosity (0-9)")
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(*verbosity))
	log.Root().SetHandler(glogger)

	time.Sleep(time.Duration(20) * time.Second)

	for {
		//var count, peers = 0, ""
		//get selfID, peers
		//count, peers := GetEnodes()
		//fmt.Printf("\ncount: %v, peers: %v\n", count, peers)

		//count, peers = GetGroup()
		//fmt.Printf("\n\ngroup: %v, peers: %v\n", count, peers)

		//if count >= 2 {
		//	enode1 := strings.Split(peers, "dcrmmsg")
		//	enode2 := strings.Split(enode1[1], "dcrmmsg")
		//	fmt.Printf("enode1: %+v\nenode2: %+v\n", enode1[0], enode2[0])
		//	p := "hhhhhhhhhh"
		//	SendToPeer(enode1[0], &p)
		//	SendToPeer(enode2[0], &p)
		//}
		//group: 4, peers: enode://74ec982620b1a9929b19e1373e74347289d43b8f6cd96dd03af8b72799a75139d601338dbb48e0786a304b28f35325407a0535625e1f8ead6f9292aeda0b4fd5@10.192.32.92:1236dcrmmsgenode://c25d9eb7e5100fc533a6507b0a2a1e1df027caa861d0c3b6ea1e6ade0fd17f3e932d1198dac2d13d02fdd894601b403d2ebe1a040d90eb409b7b68aad8c02e90@10.192.32.92:1237dcrmmsgenode://4538612f7d2b63aeea0adc96d550d3fa346c9abcdbde27e623dfb7a5c2977fdee0116047a826c89f20ef1d4fc44c6bba077e0039aa05b26608f081128a2780e6@10.192.32.92:1234dcrmmsgenode://3b2fd28db22477b9d3c5d51b12c36dc8f6af1e5b287e0de7353033ac6f3bd3ee3ae3d10256cc9af9d7af169749dfa09feaad03b398ba6c4c2a7ee559c11f8b13@10.192.32.92:1235
		//dcrm.SendMsg("test.........................")

		//fmt.Printf("\n\n\n----------\nBroatcastToGroup 1300 ...\n")
		//BroatcastToGroup("ggggggggggg, 149538132774657726528707636360336072285983621283716341866783381657518370346851019858260832456828530008033228638611352429959811432602530865173229170389257564378211923685950755531505238837058174086920335663658007299016833930579190142733302787192962349263819218016641353162586263128086590197536551686649470858150350618785456971418349930832104986121457695940172259894785398687533391786467867726631938596486511367901516568978112757142419630344178113667535121256560750392040658835275511942192701455632856375360094654110506609442227926625226950851473547011845863799213473989320072877761242218539582755795868000263697836998474090588662414706888454874809674266627030310050419787442102412880785593702679692045546544678324521807290617899168862623088851045574478150024227834019720456951861075082019330308080142876318264220503114622414715335720216012228777904424800272212487065178597668582265177598601386221891513790771444030710626346787584945068392174449499454227939574351840134966158644998994253071165243081522707017502892682948657806254474909058265538220056912463355107659552522125535221371768932967262577221596192288246964658379224806800363453895132533539966950472006685897686030983111842328749543010305869598146974505445468053138588172334973")
		//time.Sleep(time.Duration(5) * time.Second)
		//fmt.Printf("\n----------\nBroatcastToGroup >2000 ...\n")
		//BroatcastToGroup("ggggggggggg, 149538132774657726528707636360336072285983621283716341866783381657518370346851019858260832456828530008033228638611352429959811432602530865173229170389257564378211923685950755531505238837058174086920335663658007299016833930579190142733302787192962349263819218016641353162586263128086590197536551686649470858150350618785456971418349930832104986121457695940172259894785398687533391786467867726631938596486511367901516568978112757142419630344178113667535121256560750392040658835275511942192701455632856375360094654110506609442227926625226950851473547011845863799213473989320072877761242218539582755795868000263697836998474090588662414706888454874809674266627030310050419787442102412880785593702679692045546544678324521807290617899168862623088851045574478150024227834019720456951861075082019330308080142876318264220503114622414715335720216012228777904424800272212487065178597668582265177598601386221891513790771444030710626346787584945068392174449499454227939574351840134966158644998994253071165243081522707017502892682948657806254474909058265538220056912463355107659552522125535221371768932967262577221596192288246964658379224806800363453895132533539966950472006685897686030983111842328749543010305869598146974505445468053138588172334973--149538132774657726528707636360336072285983621283716341866783381657518370346851019858260832456828530008033228638611352429959811432602530865173229170389257564378211923685950755531505238837058174086920335663658007299016833930579190142733302787192962349263819218016641353162586263128086590197536551686649470858150350618785456971418349930832104986121457695940172259894785398687533391786467867726631938596486511367901516568978112757142419630344178113667535121256560750392040658835275511942192701455632856375360094654110506609442227926625226950851473547011845863799213473989320072877761242218539582755795868000263697836998474090588662414706888454874809674266627030310050419787442102412880785593702679692045546544678324521807290617899168862623088851045574478150024227834019720456951861075082019330308080142876318264220503114622414715335720216012228777904424800272212487065178597668582265177598601386221891513790771444030710626346787584945068392174449499454227939574351840134966158644998994253071165243081522707017502892682948657806254474909058265538220056912463355107659552522125535221371768932967262577221596192288246964658379224806800363453895132533539966950472006685897686030983111842328749543010305869598146974505445468053138588172334973gg")
		time.Sleep(time.Duration(5) * time.Second)
		fmt.Printf("\n\n\n----------\nBroatcast 1300 ...\n")
		Broatcast("b 1200, 149538132774657726528707636360336072285983621283716341866783381657518370346851019858260832456828530008033228638611352429959811432602530865173229170389257564378211923685950755531505238837058174086920335663658007299016833930579190142733302787192962349263819218016641353162586263128086590197536551686649470858150350618785456971418349930832104986121457695940172259894785398687533391786467867726631938596486511367901516568978112757142419630344178113667535121256560750392040658835275511942192701455632856375360094654110506609442227926625226950851473547011845863799213473989320072877761242218539582755795868000263697836998474090588662414706888454874809674266627030310050419787442102412880785593702679692045546544678324521807290617899168862623088851045574478150024227834019720456951861075082019330308080142876318264220503114622414715335720216012228777904424800272212487065178597668582265177598601386221891513790771444030710626346787584945068392174449499454227939574351840134966158644998994253071165243081522707017502892682948657806254474909058265538220056912463355107659552522125535221371768932967262577221596192288246964658379224806800363453895132533539966950472006685897686030983111842328749543010305869598146974505445468053138588172334973")
		time.Sleep(time.Duration(10) * time.Second)
		fmt.Printf("\n\n\n----------\nBroatcast >2000 ...\n")
		Broatcast("b 2300, 149538132774657726528707636360336072285983621283716341866783381657518370346851019858260832456828530008033228638611352429959811432602530865173229170389257564378211923685950755531505238837058174086920335663658007299016833930579190142733302787192962349263819218016641353162586263128086590197536551686649470858150350618785456971418349930832104986121457695940172259894785398687533391786467867726631938596486511367901516568978112757142419630344178113667535121256560750392040658835275511942192701455632856375360094654110506609442227926625226950851473547011845863799213473989320072877761242218539582755795868000263697836998474090588662414706888454874809674266627030310050419787442102412880785593702679692045546544678324521807290617899168862623088851045574478150024227834019720456951861075082019330308080142876318264220503114622414715335720216012228777904424800272212487065178597668582265177598601386221891513790771444030710626346787584945068392174449499454227939574351840134966158644998994253071165243081522707017502892682948657806254474909058265538220056912463355107659552522125535221371768932967262577221596192288246964658379224806800363453895132533539966950472006685897686030983111842328749543010305869598146974505445468053138588172334973--149538132774657726528707636360336072285983621283716341866783381657518370346851019858260832456828530008033228638611352429959811432602530865173229170389257564378211923685950755531505238837058174086920335663658007299016833930579190142733302787192962349263819218016641353162586263128086590197536551686649470858150350618785456971418349930832104986121457695940172259894785398687533391786467867726631938596486511367901516568978112757142419630344178113667535121256560750392040658835275511942192701455632856375360094654110506609442227926625226950851473547011845863799213473989320072877761242218539582755795868000263697836998474090588662414706888454874809674266627030310050419787442102412880785593702679692045546544678324521807290617899168862623088851045574478150024227834019720456951861075082019330308080142876318264220503114622414715335720216012228777904424800272212487065178597668582265177598601386221891513790771444030710626346787584945068392174449499454227939574351840134966158644998994253071165243081522707017502892682948657806254474909058265538220056912463355107659552522125535221371768932967262577221596192288246964658379224806800363453895132533539966950472006685897686030983111842328749543010305869598146974505445468053138588172334973gg")
		//time.Sleep(time.Duration(2) * time.Second)
		//Broatcast("hhhhhhwwwwwwjjjj, 20181108")

		//fmt.Printf("\nSendToGroup 1300 ...\n")
		//ret := SendToDcrmGroup("hhhhhhhhhhhhh")
		//fmt.Printf("send hhhhhhh, get = %+v\n", ret)
		//time.Sleep(time.Duration(5) * time.Second)
		//rett := SendToDcrmGroup("149538132774657726528707636360336072285983621283716341866783381657518370346851019858260832456828530008033228638611352429959811432602530865173229170389257564378211923685950755531505238837058174086920335663658007299016833930579190142733302787192962349263819218016641353162586263128086590197536551686649470858150350618785456971418349930832104986121457695940172259894785398687533391786467867726631938596486511367901516568978112757142419630344178113667535121256560750392040658835275511942192701455632856375360094654110506609442227926625226950851473547011845863799213473989320072877761242218539582755795868000263697836998474090588662414706888454874809674266627030310050419787442102412880785593702679692045546544678324521807290617899168862623088851045574478150024227834019720456951861075082019330308080142876318264220503114622414715335720216012228777904424800272212487065178597668582265177598601386221891513790771444030710626346787584945068392174449499454227939574351840134966158644998994253071165243081522707017502892682948657806254474909058265538220056912463355107659552522125535221371768932967262577221596192288246964658379224806800363453895132533539966950472006685897686030983111842328749543010305869598146974505445468053138588172334973")
		//fmt.Printf("send > 800, get = %+v\n", rett)
		//time.Sleep(time.Duration(5) * time.Second)
		//fmt.Printf("\nSendToGroup 2000 ...\n")
		//rettt := SendToDcrmGroup("149538132774657726528707636360336072285983621283716341866783381657518370346851019858260832456828530008033228638611352429959811432602530865173229170389257564378211923685950755531505238837058174086920335663658007299016833930579190142733302787192962349263819218016641353162586263128086590197536551686649470858150350618785456971418349930832104986121457695940172259894785398687533391786467867726631938596486511367901516568978112757142419630344178113667535121256560750392040658835275511942192701455632856375360094654110506609442227926625226950851473547011845863799213473989320072877761242218539582755795868000263697836998474090588662414706888454874809674266627030310050419787442102412880785593702679692045546544678324521807290617899168862623088851045574478150024227834019720456951861075082019330308080142876318264220503114622414715335720216012228777904424800272212487065178597668582265177598601386221891513790771444030710626346787584945068392174449499454227939574351840134966158644998994253071165243081522707017502892682948657806254474909058265538220056912463355107659552522125535221371768932967262577221596192288246964658379224806800363453895132533539966950472006685897686030983111842328749543010305869598146974505445468053138588172334973--149538132774657726528707636360336072285983621283716341866783381657518370346851019858260832456828530008033228638611352429959811432602530865173229170389257564378211923685950755531505238837058174086920335663658007299016833930579190142733302787192962349263819218016641353162586263128086590197536551686649470858150350618785456971418349930832104986121457695940172259894785398687533391786467867726631938596486511367901516568978112757142419630344178113667535121256560750392040658835275511942192701455632856375360094654110506609442227926625226950851473547011845863799213473989320072877761242218539582755795868000263697836998474090588662414706888454874809674266627030310050419787442102412880785593702679692045546544678324521807290617899168862623088851045574478150024227834019720456951861075082019330308080142876318264220503114622414715335720216012228777904424800272212487065178597668582265177598601386221891513790771444030710626346787584945068392174449499454227939574351840134966158644998994253071165243081522707017502892682948657806254474909058265538220056912463355107659552522125535221371768932967262577221596192288246964658379224806800363453895132533539966950472006685897686030983111842328749543010305869598146974505445468053138588172334973--sg")
		//fmt.Printf("send > 800, get = %+v\n", rettt)
		break
	}
	select {}

	time.Sleep(time.Duration(5) * time.Second)

	//send
	//ms := msg()
	//SendMsg(ms)

	select {}
}

func msg() string {
	ss := ""
	for i := 0; i < 20; i++ {
		a, _ := new(big.Int).SetString("149538132774657726528707636360336072285983621283716341866783381657518370346851019858260832456828530008033228638611352429959811432602530865173229170389257564378211923685950755531505238837058174086920335663658007299016833930579190142733302787192962349263819218016641353162586263128086590197536551686649470858150350618785456971418349930832104986121457695940172259894785398687533391786467867726631938596486511367901516568978112757142419630344178113667535121256560750392040658835275511942192701455632856375360094654110506609442227926625226950851473547011845863799213473989320072877761242218539582755795868000263697836998474090588662414706888454874809674266627030310050419787442102412880785593702679692045546544678324521807290617899168862623088851045574478150024227834019720456951861075082019330308080142876318264220503114622414715335720216012228777904424800272212487065178597668582265177598601386221891513790771444030710626346787584945068392174449499454227939574351840134966158644998994253071165243081522707017502892682948657806254474909058265538220056912463355107659552522125535221371768932967262577221596192288246964658379224806800363453895132533539966950472006685897686030983111842328749543010305869598146974505445468053138588172334973", 10)

		s := a.Bytes()
		ss += string(s[:])
		ss += "dcrmmsg"
	}
	return ss
}
