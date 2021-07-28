package main

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	Account         string        `json:"account"`
	Created         int64         `json:"created"`
	Data            *EventData    `json:"data"`
	ID              string        `json:"id"`
	Livemode        bool          `json:"livemode"`
	PendingWebhooks int64         `json:"pending_webhooks"`
	Request         *EventRequest `json:"request"`
	Type            string        `json:"type"`
}

type EventData struct {
	// Object is a raw mapping of the API resource contained in the event.
	// Although marked with json:"-", it's still populated independently by
	// a custom UnmarshalJSON implementation.
	Object             map[string]interface{} `json:"-"`
	PreviousAttributes map[string]interface{} `json:"previous_attributes"`
	Raw                json.RawMessage        `json:"object"`
}

type EventRequest struct {
	// ID is the request ID of the request that created an event, if the event
	// was created by a request.
	ID string `json:"id"`

	// IdempotencyKey is the idempotency key of the request that created an
	// event, if the event was created by a request and if an idempotency key
	// was specified for that request.
	IdempotencyKey string `json:"idempotency_key"`
}

func main1() {
	a := `{"account":"","created":1620916236,"data":{"previous_attributes":null,"object":{"id":"pi_1IqfYcCxfCjFfGO0DaOTreyh","object":"payment_intent","amount":3000,"amount_capturable":0,"amount_received":3000,"application":null,"application_fee_amount":null,"canceled_at":null,"cancellation_reason":null,"capture_method":"automatic","charges":{"object":"list","data":[{"id":"ch_1IqfYdCxfCjFfGO04pTWmSf6","object":"charge","amount":3000,"amount_captured":3000,"amount_refunded":0,"application":null,"application_fee":null,"application_fee_amount":null,"balance_transaction":"txn_1IqfYeCxfCjFfGO0xPat2ulR","billing_details":{"address":{"city":null,"country":null,"line1":null,"line2":null,"postal_code":null,"state":null},"email":null,"name":"zpj","phone":null},"calculated_statement_descriptor":"Stripe","captured":true,"created":1620916235,"currency":"usd","customer":"cus_JTckySjy74tZDG","description":null,"destination":null,"dispute":null,"disputed":false,"failure_code":null,"failure_message":null,"fraud_details":{},"invoice":null,"livemode":false,"metadata":{"order_id":"c6ef4e4b-b3f7-11eb-97fa-023a99d3797d"},"on_behalf_of":null,"order":null,"outcome":{"network_status":"approved_by_network","reason":null,"risk_level":"normal","risk_score":22,"seller_message":"Payment complete.","type":"authorized"},"paid":true,"payment_intent":"pi_1IqfYcCxfCjFfGO0DaOTreyh","payment_method":"pm_1IqfY9CxfCjFfGO0CEKLmyOs","payment_method_details":{"card":{"brand":"visa","checks":{"address_line1_check":null,"address_postal_code_check":null,"cvc_check":"pass"},"country":"US","exp_month":12,"exp_year":2025,"fingerprint":"zh87OO9FQ0qUwFtw","funding":"credit","installments":null,"last4":"4242","network":"visa","three_d_secure":null,"wallet":null},"type":"card"},"receipt_email":null,"receipt_number":null,"receipt_url":"https://pay.stripe.com/receipts/acct_1HeATTCxfCjFfGO0/ch_1IqfYdCxfCjFfGO04pTWmSf6/rcpt_JTcoI3oheDX98ZLNLA94XakDW5DXG9o","refunded":false,"refunds":{"object":"list","data":[],"has_more":false,"total_count":0,"url":"/v1/charges/ch_1IqfYdCxfCjFfGO04pTWmSf6/refunds"},"review":null,"shipping":null,"source":null,"source_transfer":null,"statement_descriptor":null,"statement_descriptor_suffix":null,"status":"succeeded","transfer_data":null,"transfer_group":null}],"has_more":false,"total_count":1,"url":"/v1/charges?payment_intent=pi_1IqfYcCxfCjFfGO0DaOTreyh"},"client_secret":"pi_1IqfYcCxfCjFfGO0DaOTreyh_secret_0Y5f1jaXNi9MhP4dIlYRPC7fs","confirmation_method":"automatic","created":1620916234,"currency":"usd","customer":"cus_JTckySjy74tZDG","description":null,"invoice":null,"last_payment_error":null,"livemode":false,"metadata":{"order_id":"c6ef4e4b-b3f7-11eb-97fa-023a99d3797d"},"next_action":null,"on_behalf_of":null,"payment_method":"pm_1IqfY9CxfCjFfGO0CEKLmyOs","payment_method_options":{"card":{"installments":null,"network":null,"request_three_d_secure":"automatic"}},"payment_method_types":["card"],"receipt_email":null,"review":null,"setup_future_usage":null,"shipping":null,"source":null,"statement_descriptor":null,"statement_descriptor_suffix":null,"status":"succeeded","transfer_data":null,"transfer_group":null}},"id":"evt_1IqfYeCxfCjFfGO0Yjm3j8dE","livemode":false,"pending_webhooks":2,"request":{"id":"req_VF7qXtjJVbDPGu","idempotency_key":""},"type":"payment_intent.succeeded"}`
	event := new(Event)
	err := json.Unmarshal([]byte(a), event)
	fmt.Println("111", err, event)
	if event.Type != "payment_intent.succeeded" {
		fmt.Println("111")
	}
	maa,err := json.Marshal(event.Data.Raw)

	fmt.Println("aaa",string(maa),err)
	obj := make(map[string]interface{},0)
	err = json.Unmarshal(maa,&obj)
	fmt.Println("----",err)
	event.Data.Object = obj
	object := event.Data.Object
	metadata, ok := object["metadata"].(map[string]interface{})
	fmt.Println("metadata", metadata, ok)

	//TODO:更新订单状态为已支付，订单ID获取方法
	orderID, ok := metadata["order_id"].(string)

	fmt.Println("orderID", orderID, ok)
	secrect, ok := object["client_secret"].(string)
	fmt.Println("secrect", secrect, ok)
	type Charges struct {
		ID string `json:"id"`
	}

	type ChargeRequestList struct {
		Data []Charges `json:"data"`
	}

	type ChargeRequestResp struct {
		Data []Charges `json:"data"`
	}
	type TestData struct {
		Charges ChargeRequestResp `json:"charges"`
	}
	chargs1,err := json.Marshal(object["charges"])
	fmt.Println("=========",string(chargs1),err)

	var Ord TestData

	err = json.Unmarshal(event.Data.Raw,&Ord)
	fmt.Println(Ord,Ord.Charges.Data,err)

	charges, ok := object["charges"].(ChargeRequestResp)
	fmt.Println("+++++++",charges,ok)
	charges2, ok := object["charges"].(map[string]interface{})
	fmt.Println("*******",charges2["data"],ok)
	charges3,ok := charges2["data"].([]interface{})
	fmt.Println("*******",charges3,ok)
	for _,v:=range charges3{
		charges4,ok := v.(map[string]interface{})
		fmt.Println("*******",charges4["id"],ok)
	}

}

func main()  {
	main1()
	//b := `{"test":{"a":111,"b":222}}`
	//
	//type Test1 struct {
	//	Test interface{} `json:"test"`
	//}
	//var a Test1
	//err := json.Unmarshal([]byte(b),&a)
	//fmt.Println(err)
	//
	//type t2 struct{
	//	A int `json:"a"`
	//	B int `json:"b"`
	//}
	//
	//aa,ok:= a.Test.(t2)
	//fmt.Println(aa,ok)

}
