package internal

import (
	"bufio"
	"fmt"
	"time"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/quickfix"
	"github.com/shopspring/decimal"

	"os"
	"strconv"
	"strings"

	fix40nos "github.com/quickfixgo/fix40/newordersingle"
	fix41nos "github.com/quickfixgo/fix41/newordersingle"
	fix42nos "github.com/quickfixgo/fix42/newordersingle"
	fix43nos "github.com/quickfixgo/fix43/newordersingle"
	fix44nos "github.com/quickfixgo/fix44/newordersingle"
	fix50nos "github.com/quickfixgo/fix50/newordersingle"

	fix40cxl "github.com/quickfixgo/fix40/ordercancelrequest"
	fix41cxl "github.com/quickfixgo/fix41/ordercancelrequest"
	fix42cxl "github.com/quickfixgo/fix42/ordercancelrequest"
	fix43cxl "github.com/quickfixgo/fix43/ordercancelrequest"
	fix44cxl "github.com/quickfixgo/fix44/ordercancelrequest"
	fix50cxl "github.com/quickfixgo/fix50/ordercancelrequest"

	fix42mdr "github.com/quickfixgo/fix42/marketdatarequest"
	fix43mdr "github.com/quickfixgo/fix43/marketdatarequest"
	fix44mdr "github.com/quickfixgo/fix44/marketdatarequest"
	fix50mdr "github.com/quickfixgo/fix50/marketdatarequest"
)

func queryString(fieldName string) string {
	fmt.Printf("%v: ", fieldName)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return scanner.Text()
}

func queryDecimal(fieldName string) decimal.Decimal {
	val, err := decimal.NewFromString(queryString(fieldName))
	if err != nil {
		panic(err)
	}

	return val
}

func queryFieldChoices(fieldName string, choices []string, values []string) string {
	for i, choice := range choices {
		fmt.Printf("%v) %v\n", i+1, choice)
	}

	choiceStr := queryString(fieldName)
	choice, err := strconv.Atoi(choiceStr)
	if err != nil || choice < 1 || choice > len(choices) {
		panic(fmt.Errorf("Invalid %v: %v", fieldName, choice))
	}

	if values == nil {
		return choiceStr
	}

	return values[choice-1]
}

func QueryAction(neworder string) (string) {
	neworder2 := strings.Split(neworder, " ")
	action := neworder2[0]
	action = action[1:len(action)]
	return action

}

func queryVersion(neworder string) (string, error) {
	neworder2 := strings.Split(neworder, " ")
	version := neworder2[1]

	switch version {
	case "1":
		fmt.Println("Version: FIX 4.0")
		return quickfix.BeginStringFIX40, nil
	case "2":
		fmt.Println("Version: FIX 4.1")
		return quickfix.BeginStringFIX41, nil
	case "3":
		fmt.Println("Version: FIX 4.2")
		return quickfix.BeginStringFIX42, nil
	case "4":
		fmt.Println("Version: FIX 4.3")
		return quickfix.BeginStringFIX43, nil
	case "5":
		fmt.Println("Version: FIX 4.4")
		return quickfix.BeginStringFIX44, nil
	case "6":
		fmt.Println("Version: FIX 1.1")
		return quickfix.BeginStringFIXT11, nil
	}

	return "", fmt.Errorf("unknown BeginString choice: %v", version)
}

func queryOrigClOrdID() field.OrigClOrdIDField {
	return field.NewOrigClOrdID(("OrigClOrdID"))
}

func queryClOrdID(clordid string) field.ClOrdIDField {
	fmt.Println("ClOrdID: "+clordid)
	return field.NewClOrdID(clordid)
}

func querySide(side string) field.SideField {

	switch side {
	case "1":
		fmt.Println("Side: Buy")
		side := string(enum.Side_BUY)
		return field.NewSide(enum.Side(side)) 
	case "2":
		fmt.Println("Side: Sell")
		side := string(enum.Side_SELL)
		return field.NewSide(enum.Side(side))
	case "3":
		fmt.Println("Side: Sell Short")
		side := string(enum.Side_SELL_SHORT)
		return field.NewSide(enum.Side(side))
	case "4":
		fmt.Println("Side: Sell Short Exempt")
		side := string(enum.Side_SELL_SHORT_EXEMPT)
		return field.NewSide(enum.Side(side))
	case "5":
		fmt.Println("Side: Cross")
		side := string(enum.Side_CROSS)
		return field.NewSide(enum.Side(side))
	case "6":
		fmt.Println("Side: Cross Short")
		side := string(enum.Side_CROSS_SHORT)
		return field.NewSide(enum.Side(side))
	case "7":
		fmt.Println("Side: Cross Short Exempt")
		side := "A"
		return field.NewSide(enum.Side(side))
	}

	return field.NewSide(enum.Side(side))

}

func queryOrdType(f *field.OrdTypeField, ordtype string) field.OrdTypeField {

	switch ordtype {
	case "1":
		fmt.Println("OrderType: Market")
		ordtype := string(enum.OrdType_MARKET)
		f.FIXString = quickfix.FIXString(ordtype)
		return *f
	case "2":
		fmt.Println("OrderType: Limit")
		ordtype := string(enum.OrdType_LIMIT)
		f.FIXString = quickfix.FIXString(ordtype)
		return *f
	case "3":
		fmt.Println("OrderType: Stop")
		ordtype := string(enum.OrdType_STOP)
		f.FIXString = quickfix.FIXString(ordtype)
		return *f
	case "4":
		fmt.Println("OrderType: Stop Limit")
		ordtype := string(enum.OrdType_STOP_LIMIT)
		f.FIXString = quickfix.FIXString(ordtype)
		return *f
	}

	f.FIXString = quickfix.FIXString(ordtype)
	return *f
}

func querySymbol(symbol string) field.SymbolField {
	fmt.Println("Symbol: "+symbol)
	return field.NewSymbol(symbol)
}

func queryOrderQty(orderqty string) field.OrderQtyField {
	fmt.Println("OrderQty: "+orderqty)
	orderqty2, _ := decimal.NewFromString(orderqty)
	return field.NewOrderQty(orderqty2, 2)
}

func queryTimeInForce(timeinforce string) field.TimeInForceField {

	switch timeinforce {
	case "1":
		fmt.Println("TimeInForce: Day")
		timeinforce := string(enum.TimeInForce_DAY)
		return field.NewTimeInForce(enum.TimeInForce(timeinforce))
	case "2":
		fmt.Println("TimeInForce: IOC")
		timeinforce := string(enum.TimeInForce_IMMEDIATE_OR_CANCEL)
		return field.NewTimeInForce(enum.TimeInForce(timeinforce))
	case "3":
		fmt.Println("TimeInForce: OPG")
		timeinforce := string(enum.TimeInForce_AT_THE_OPENING)
		return field.NewTimeInForce(enum.TimeInForce(timeinforce))
	case "4":
		fmt.Println("TimeInForce: GTC")
		timeinforce := string(enum.TimeInForce_GOOD_TILL_CANCEL)
		return field.NewTimeInForce(enum.TimeInForce(timeinforce))
	case "5":
		fmt.Println("TimeInForce: GTX")
		timeinforce := string(enum.TimeInForce_GOOD_TILL_CROSSING)
		return field.NewTimeInForce(enum.TimeInForce(timeinforce))
	}
	return field.NewTimeInForce(enum.TimeInForce(timeinforce))
}

func queryPrice(price string) field.PriceField {
	fmt.Println("Price: "+price)
	price2, _ := decimal.NewFromString(price)
	return field.NewPrice(price2, 2)
}

func queryStopPx() field.StopPxField {
	return field.NewStopPx(queryDecimal("Stop Price"), 2)
}

func querySenderCompID(senderid string) field.SenderCompIDField {
	fmt.Println("SenderCompID: "+senderid)
	return field.NewSenderCompID(senderid)
}

func queryTargetCompID(targetid string) field.TargetCompIDField {
	fmt.Println("TargetCompID: "+targetid)
	return field.NewTargetCompID(targetid)
}

func queryTargetSubID(targetsubid string) field.TargetSubIDField {
	fmt.Println("TargetSubID: "+targetsubid)
	return field.NewTargetSubID(targetsubid)
}

func queryHeader(h header, senderid string, targetid string, targetsubid string) {
	h.Set(querySenderCompID(senderid))
	h.Set(queryTargetCompID(targetid))
	h.Set(queryTargetSubID(targetsubid))
}

type header interface {
	Set(f quickfix.FieldWriter) *quickfix.FieldMap
}

func queryConfirm(prompt string) bool {
	fmt.Println()
	fmt.Printf("%v?: ", prompt)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return strings.ToUpper(scanner.Text()) == "Y"
}

func queryNewOrderSingle40(neworder string) fix40nos.NewOrderSingle {
	var ordType field.OrdTypeField
	neworder2 := strings.Split(neworder, " ")
	clordid := neworder2[2]
	price := neworder2[3]
	symbol := neworder2[4]
	orderqty := neworder2[5]
	side := neworder2[6]
	ordtype := neworder2[7]
	timeinforce := neworder2[8]
	senderid := neworder2[9]
	targetid := neworder2[10]
	targetsubid := neworder2[11][0:(len(neworder2[11])-1)]

	order := fix40nos.New(queryClOrdID(clordid), field.NewHandlInst("1"), querySymbol(symbol), querySide(side), queryOrderQty(orderqty), queryOrdType(&ordType, ordtype))

	switch ordType.Value() {
	case enum.OrdType_LIMIT, enum.OrdType_STOP_LIMIT:
		order.Set(queryPrice(price))
	}

	switch ordType.Value() {
	case enum.OrdType_STOP, enum.OrdType_STOP_LIMIT:
		order.Set(queryStopPx())
	}

	order.Set(queryTimeInForce(timeinforce))
	h := order.Header.Header
	h.Set(querySenderCompID(senderid))
	h.Set(queryTargetCompID(targetid))
	h.Set(queryTargetSubID(targetsubid))

	return order
}

func queryNewOrderSingle41(neworder string) (msg *quickfix.Message) {
	var ordType field.OrdTypeField
	neworder2 := strings.Split(neworder, " ")
	clordid := neworder2[2]
	price := neworder2[3]
	symbol := neworder2[4]
	orderqty := neworder2[5]
	side := neworder2[6]
	ordtype := neworder2[7]
	timeinforce := neworder2[8]
	senderid := neworder2[9]
	targetid := neworder2[10]
	targetsubid := neworder2[11][0:(len(neworder2[11])-1)]

	order := fix41nos.New(queryClOrdID(clordid), field.NewHandlInst("1"), querySymbol(symbol), querySide(side), queryOrdType(&ordType, ordtype))
	order.Set(queryOrderQty(orderqty))

	switch ordType.Value() {
	case enum.OrdType_LIMIT, enum.OrdType_STOP_LIMIT:
		order.Set(queryPrice(price))
	}

	switch ordType.Value() {
	case enum.OrdType_STOP, enum.OrdType_STOP_LIMIT:
		order.Set(queryStopPx())
	}

	order.Set(queryTimeInForce(timeinforce))
	msg = order.ToMessage()
	h := &msg.Header
	h.Set(querySenderCompID(senderid))
	h.Set(queryTargetCompID(targetid))
	h.Set(queryTargetSubID(targetsubid))

	return
}

func queryNewOrderSingle42(neworder string) (msg *quickfix.Message) {
	var ordType field.OrdTypeField
	neworder2 := strings.Split(neworder, " ")
	clordid := neworder2[2]
	price := neworder2[3]
	symbol := neworder2[4]
	orderqty := neworder2[5]
	side := neworder2[6]
	ordtype := neworder2[7]
	timeinforce := neworder2[8]
	senderid := neworder2[9]
	targetid := neworder2[10]
	targetsubid := neworder2[11][0:(len(neworder2[11])-1)]

	order := fix42nos.New(queryClOrdID(clordid), field.NewHandlInst("1"), querySymbol(symbol), querySide(side), field.NewTransactTime(time.Now()), queryOrdType(&ordType, ordtype))
	order.Set(queryOrderQty(orderqty))

	switch ordType.Value() {
	case enum.OrdType_LIMIT, enum.OrdType_STOP_LIMIT:
		order.Set(queryPrice(price))
	}

	switch ordType.Value() {
	case enum.OrdType_STOP, enum.OrdType_STOP_LIMIT:
		order.Set(queryStopPx())
	}

	order.Set(queryTimeInForce(timeinforce))
	msg = order.ToMessage()
	h := &msg.Header
	h.Set(querySenderCompID(senderid))
	h.Set(queryTargetCompID(targetid))
	h.Set(queryTargetSubID(targetsubid))

	return
}

func queryNewOrderSingle43(neworder string) (msg *quickfix.Message) {
	var ordType field.OrdTypeField
	neworder2 := strings.Split(neworder, " ")
	clordid := neworder2[2]
	price := neworder2[3]
	symbol := neworder2[4]
	orderqty := neworder2[5]
	side := neworder2[6]
	ordtype := neworder2[7]
	timeinforce := neworder2[8]
	senderid := neworder2[9]
	targetid := neworder2[10]
	targetsubid := neworder2[11][0:(len(neworder2[11])-1)]
	
	order := fix43nos.New(queryClOrdID(clordid), field.NewHandlInst("1"), querySide(side), field.NewTransactTime(time.Now()), queryOrdType(&ordType, ordtype))
	order.Set(querySymbol(symbol))
	order.Set(queryOrderQty(orderqty))

	switch ordType.Value() {
	case enum.OrdType_LIMIT, enum.OrdType_STOP_LIMIT:
		order.Set(queryPrice(price))
	}

	switch ordType.Value() {
	case enum.OrdType_STOP, enum.OrdType_STOP_LIMIT:
		order.Set(queryStopPx())
	}

	order.Set(queryTimeInForce(timeinforce))
	msg = order.ToMessage()
	h := &msg.Header
	h.Set(querySenderCompID(senderid))
	h.Set(queryTargetCompID(targetid))
	h.Set(queryTargetSubID(targetsubid))

	return
}

func queryNewOrderSingle44(neworder string) (msg *quickfix.Message) {
	var ordType field.OrdTypeField

	neworder2 := strings.Split(neworder, " ")
	clordid := neworder2[2]
	price := neworder2[3]
	symbol := neworder2[4]
	orderqty := neworder2[5]
	side := neworder2[6]
	ordtype := neworder2[7]
	timeinforce := neworder2[8]
	senderid := neworder2[9]
	targetid := neworder2[10]
	targetsubid := neworder2[11][0:(len(neworder2[11])-1)]

	order := fix44nos.New(queryClOrdID(clordid), querySide(side), field.NewTransactTime(time.Now()), queryOrdType(&ordType, ordtype))
	order.SetHandlInst("1")
	order.Set(querySymbol(symbol))
	order.Set(queryOrderQty(orderqty))
	order.Set(queryTimeInForce(timeinforce))

	switch ordType.Value() {
	case enum.OrdType_LIMIT, enum.OrdType_STOP_LIMIT:
		order.Set(queryPrice(price))
	}

	//switch ordType.Value() {
	//case enum.OrdType_STOP, enum.OrdType_STOP_LIMIT:
	//	order.Set(queryStopPx())
	//}

	msg = order.ToMessage()
	h := &msg.Header
	h.Set(querySenderCompID(senderid))
	h.Set(queryTargetCompID(targetid))
	h.Set(queryTargetSubID(targetsubid))
	

	return
}


func queryNewOrderSingle50(neworder string) (msg *quickfix.Message) {
	var ordType field.OrdTypeField

	neworder2 := strings.Split(neworder, " ")
	clordid := neworder2[2]
	price := neworder2[3]
	symbol := neworder2[4]
	orderqty := neworder2[5]
	side := neworder2[6]
	ordtype := neworder2[7]
	timeinforce := neworder2[8]
	senderid := neworder2[9]
	targetid := neworder2[10]
	targetsubid := neworder2[11][0:(len(neworder2[11])-1)]

	order := fix50nos.New(queryClOrdID(clordid), querySide(side), field.NewTransactTime(time.Now()), queryOrdType(&ordType, ordtype))
	order.SetHandlInst("1")
	order.Set(querySymbol(symbol))
	order.Set(queryOrderQty(orderqty))
	order.Set(queryTimeInForce(timeinforce))

	switch ordType.Value() {
	case enum.OrdType_LIMIT, enum.OrdType_STOP_LIMIT:
		order.Set(queryPrice(price))
	}

	//switch ordType.Value() {
	//case enum.OrdType_STOP, enum.OrdType_STOP_LIMIT:
	//	order.Set(queryStopPx())
	//}

	msg = order.ToMessage()
	h := &msg.Header
	h.Set(querySenderCompID(senderid))
	h.Set(queryTargetCompID(targetid))
	h.Set(queryTargetSubID(targetsubid))
	

	return
}

func queryOrderCancelRequest40(neworder string) (msg *quickfix.Message) {
	neworder2 := strings.Split(neworder, " ")
	clordid := neworder2[2]
	symbol := neworder2[4]
	orderqty := neworder2[5]
	side := neworder2[6]
	senderid := neworder2[9]
	targetid := neworder2[10]
	targetsubid := neworder2[11][0:(len(neworder2[11])-1)]

	cancel := fix40cxl.New(queryOrigClOrdID(), queryClOrdID(clordid), field.NewCxlType("F"), querySymbol(symbol), querySide(side), queryOrderQty(orderqty))
	msg = cancel.ToMessage()
	queryHeader(&msg.Header, senderid, targetid, targetsubid)
	return
}

func queryOrderCancelRequest41(neworder string) (msg *quickfix.Message) {
	neworder2 := strings.Split(neworder, " ")
	clordid := neworder2[2]
	symbol := neworder2[4]
	orderqty := neworder2[5]
	side := neworder2[6]
	senderid := neworder2[9]
	targetid := neworder2[10]
	targetsubid := neworder2[11][0:(len(neworder2[11])-1)]

	cancel := fix41cxl.New(queryOrigClOrdID(), queryClOrdID(clordid), querySymbol(symbol), querySide(side))
	cancel.Set(queryOrderQty(orderqty))
	msg = cancel.ToMessage()
	queryHeader(&msg.Header, senderid, targetid, targetsubid)
	return
}

func queryOrderCancelRequest42(neworder string) (msg *quickfix.Message) {
	neworder2 := strings.Split(neworder, " ")
	clordid := neworder2[2]
	symbol := neworder2[4]
	orderqty := neworder2[5]
	side := neworder2[6]
	senderid := neworder2[9]
	targetid := neworder2[10]
	targetsubid := neworder2[11][0:(len(neworder2[11])-1)]

	cancel := fix42cxl.New(queryOrigClOrdID(), queryClOrdID(clordid), querySymbol(symbol), querySide(side), field.NewTransactTime(time.Now()))
	cancel.Set(queryOrderQty(orderqty))
	msg = cancel.ToMessage()
	queryHeader(&msg.Header, senderid, targetid, targetsubid)
	return
}

func queryOrderCancelRequest43(neworder string) (msg *quickfix.Message) {
	neworder2 := strings.Split(neworder, " ")
	clordid := neworder2[2]
	symbol := neworder2[4]
	orderqty := neworder2[5]
	side := neworder2[6]
	senderid := neworder2[9]
	targetid := neworder2[10]
	targetsubid := neworder2[11][0:(len(neworder2[11])-1)]

	cancel := fix43cxl.New(queryOrigClOrdID(), queryClOrdID(clordid), querySide(side), field.NewTransactTime(time.Now()))
	cancel.Set(querySymbol(symbol))
	cancel.Set(queryOrderQty(orderqty))
	msg = cancel.ToMessage()
	queryHeader(&msg.Header, senderid, targetid, targetsubid)
	return
}

func queryOrderCancelRequest44(neworder string) (msg *quickfix.Message) {
	neworder2 := strings.Split(neworder, " ")
	clordid := neworder2[2]
	symbol := neworder2[4]
	orderqty := neworder2[5]
	side := neworder2[6]
	senderid := neworder2[9]
	targetid := neworder2[10]
	targetsubid := neworder2[11][0:(len(neworder2[11])-1)]

	cancel := fix44cxl.New(queryOrigClOrdID(), queryClOrdID(clordid), querySide(side), field.NewTransactTime(time.Now()))
	cancel.Set(querySymbol(symbol))
	cancel.Set(queryOrderQty(orderqty))

	msg = cancel.ToMessage()
	queryHeader(&msg.Header, senderid, targetid, targetsubid)
	return
}

func queryOrderCancelRequest50(neworder string) (msg *quickfix.Message) {
	neworder2 := strings.Split(neworder, " ")
	clordid := neworder2[2]
	symbol := neworder2[4]
	orderqty := neworder2[5]
	side := neworder2[6]
	senderid := neworder2[9]
	targetid := neworder2[10]
	targetsubid := neworder2[11][0:(len(neworder2[11])-1)]

	cancel := fix50cxl.New(queryOrigClOrdID(), queryClOrdID(clordid), querySide(side), field.NewTransactTime(time.Now()))
	cancel.Set(querySymbol(symbol))
	cancel.Set(queryOrderQty(orderqty))
	msg = cancel.ToMessage()
	queryHeader(&msg.Header, senderid, targetid, targetsubid)
	return
}

func queryMarketDataRequest42(neworder string) fix42mdr.MarketDataRequest {
	neworder2 := strings.Split(neworder, " ")
	senderid := neworder2[9]
	targetid := neworder2[10]
	targetsubid := neworder2[11][0:(len(neworder2[11])-1)]

	request := fix42mdr.New(field.NewMDReqID("MARKETDATAID"),
		field.NewSubscriptionRequestType(enum.SubscriptionRequestType_SNAPSHOT),
		field.NewMarketDepth(0),
	)

	entryTypes := fix42mdr.NewNoMDEntryTypesRepeatingGroup()
	entryTypes.Add().SetMDEntryType(enum.MDEntryType_BID)
	request.SetNoMDEntryTypes(entryTypes)

	relatedSym := fix42mdr.NewNoRelatedSymRepeatingGroup()
	relatedSym.Add().SetSymbol("LNUX")
	request.SetNoRelatedSym(relatedSym)

	queryHeader(request.Header, senderid, targetid, targetsubid)
	return request
}

func queryMarketDataRequest43(neworder string) fix43mdr.MarketDataRequest {
	neworder2 := strings.Split(neworder, " ")
	senderid := neworder2[9]
	targetid := neworder2[10]
	targetsubid := neworder2[11][0:(len(neworder2[11])-1)]

	request := fix43mdr.New(field.NewMDReqID("MARKETDATAID"),
		field.NewSubscriptionRequestType(enum.SubscriptionRequestType_SNAPSHOT),
		field.NewMarketDepth(0),
	)

	entryTypes := fix43mdr.NewNoMDEntryTypesRepeatingGroup()
	entryTypes.Add().SetMDEntryType(enum.MDEntryType_BID)
	request.SetNoMDEntryTypes(entryTypes)

	relatedSym := fix43mdr.NewNoRelatedSymRepeatingGroup()
	relatedSym.Add().SetSymbol("LNUX")
	request.SetNoRelatedSym(relatedSym)

	queryHeader(request.Header, senderid, targetid, targetsubid)
	return request
}

func queryMarketDataRequest44(neworder string) fix44mdr.MarketDataRequest {
	neworder2 := strings.Split(neworder, " ")
	senderid := neworder2[9]
	targetid := neworder2[10]
	targetsubid := neworder2[11][0:(len(neworder2[11])-1)]

	request := fix44mdr.New(field.NewMDReqID("MARKETDATAID"),
		field.NewSubscriptionRequestType(enum.SubscriptionRequestType_SNAPSHOT),
		field.NewMarketDepth(0),
	)

	entryTypes := fix44mdr.NewNoMDEntryTypesRepeatingGroup()
	entryTypes.Add().SetMDEntryType(enum.MDEntryType_BID)
	request.SetNoMDEntryTypes(entryTypes)

	relatedSym := fix44mdr.NewNoRelatedSymRepeatingGroup()
	relatedSym.Add().SetSymbol("LNUX")
	request.SetNoRelatedSym(relatedSym)

	queryHeader(request.Header, senderid, targetid, targetsubid)
	return request
}

func queryMarketDataRequest50(neworder string) fix50mdr.MarketDataRequest {
	neworder2 := strings.Split(neworder, " ")
	senderid := neworder2[9]
	targetid := neworder2[10]
	targetsubid := neworder2[11][0:(len(neworder2[11])-1)]

	request := fix50mdr.New(field.NewMDReqID("MARKETDATAID"),
		field.NewSubscriptionRequestType(enum.SubscriptionRequestType_SNAPSHOT),
		field.NewMarketDepth(0),
	)

	entryTypes := fix50mdr.NewNoMDEntryTypesRepeatingGroup()
	entryTypes.Add().SetMDEntryType(enum.MDEntryType_BID)
	request.SetNoMDEntryTypes(entryTypes)

	relatedSym := fix50mdr.NewNoRelatedSymRepeatingGroup()
	relatedSym.Add().SetSymbol("LNUX")
	request.SetNoRelatedSym(relatedSym)

	queryHeader(request.Header, senderid, targetid, targetsubid)
	return request
}

func QueryEnterOrder(neworder string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	var beginString string
	beginString, err = queryVersion(neworder)
	if err != nil {
		return err
	}

	var order quickfix.Messagable
	switch beginString {
	case quickfix.BeginStringFIX40:
		order = queryNewOrderSingle40(neworder)

	case quickfix.BeginStringFIX41:
		order = queryNewOrderSingle41(neworder)

	case quickfix.BeginStringFIX42:
		order = queryNewOrderSingle42(neworder)

	case quickfix.BeginStringFIX43:
		order = queryNewOrderSingle43(neworder)

	case quickfix.BeginStringFIX44:
		order = queryNewOrderSingle44(neworder)

	case quickfix.BeginStringFIXT11:
		order = queryNewOrderSingle50(neworder)
	}

	return quickfix.Send(order)
}

func QueryCancelOrder(neworder string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	var beginString string
	beginString, err = queryVersion(neworder)
	if err != nil {
		return err
	}

	var cxl *quickfix.Message
	switch beginString {
	case quickfix.BeginStringFIX40:
		cxl = queryOrderCancelRequest40(neworder)

	case quickfix.BeginStringFIX41:
		cxl = queryOrderCancelRequest41(neworder)

	case quickfix.BeginStringFIX42:
		cxl = queryOrderCancelRequest42(neworder)

	case quickfix.BeginStringFIX43:
		cxl = queryOrderCancelRequest43(neworder)

	case quickfix.BeginStringFIX44:
		cxl = queryOrderCancelRequest44(neworder)

	case quickfix.BeginStringFIXT11:
		cxl = queryOrderCancelRequest50(neworder)
	}

	if queryConfirm("Send Cancel") {
		return quickfix.Send(cxl)
	}

	return
}

func QueryMarketDataRequest(neworder string) error {
	beginString, err := queryVersion(neworder)
	if err != nil {
		return err
	}

	var req quickfix.Messagable
	switch beginString {
	case quickfix.BeginStringFIX42:
		req = queryMarketDataRequest42(neworder)

	case quickfix.BeginStringFIX43:
		req = queryMarketDataRequest43(neworder)

	case quickfix.BeginStringFIX44:
		req = queryMarketDataRequest44(neworder)

	case quickfix.BeginStringFIXT11:
		req = queryMarketDataRequest50(neworder)

	default:
		return fmt.Errorf("No test for version %v", beginString)
	}

	if queryConfirm("Send MarketDataRequest") {
		return quickfix.Send(req)
	}

	return nil
}
