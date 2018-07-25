package internal

import (
	"fmt"
	"time"
	"bufio"
	"os"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/quickfix"
	"github.com/shopspring/decimal"

	"strings"

	fix50nos "github.com/quickfixgo/fix50/newordersingle"

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

//NOT DEFINED FOR STOP PRICES YET 
//func queryStopPx() field.StopPxField {
//	return field.NewStopPx(queryDecimal("Stop Price"), 2)
//}

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

type header interface {
	Set(f quickfix.FieldWriter) *quickfix.FieldMap
}

func queryNewOrderSingle50(neworder string) (msg *quickfix.Message) {
	var ordType field.OrdTypeField

	neworder2 := strings.Split(neworder, " ")
	clordid := neworder2[0]
	clordid = clordid[1:len(clordid)]
	price := neworder2[1]
	symbol := neworder2[2]
	orderqty := neworder2[3]
	side := neworder2[4]
	ordtype := neworder2[5]
	timeinforce := neworder2[6]
	senderid := neworder2[7]
	targetid := neworder2[8]
	targetsubid := neworder2[9][0:(len(neworder2[9])-1)]

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

func QueryEnterOrder(neworder string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	order := queryNewOrderSingle50(neworder)
	return quickfix.Send(order)
}