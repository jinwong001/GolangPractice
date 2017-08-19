package main

import (
	"fmt"
	"encoding/json"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

type Info struct {
	CampName   string `json:"camp_name,omitempty"`
}

func main() {
	pos, neg := adder(), adder()
	for i := 10; i < 11; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	 str:=`[{"camp_name":"小太阳麻食单品减至券2 20170810","today_used_user_cnt":"0","biz_date":"20170815","today_used_cnt":"0","today_campaign_new_user_cnt":"0","today_taken_cnt":"0","today_campaign_trade_amt":"0","today_taken_user_cnt":"0"},{"camp_name":"小太阳麻食单品减至券2 20170810","total_relative_period_pid_repay_trade_cnt":"126","total_relative_period_pid_loss_trade_amt":"2","today_used_user_cnt":"0","total_relative_period_pid_trade_cnt":"171","total_camp_period_pid_loss_trade_cnt":"3","biz_date":"20170814","total_used_cnt":"2","today_used_cnt":"0","total_relative_period_pid_repay_trade_user_cnt":"33","total_camp_period_pid_loss_order_amt":"3","total_taken_user_cnt":"1","total_relative_period_pid_repay_trade_amt":"4561","total_campaign_order_amt":"8","total_relative_period_pid_loss_trade_user_cnt":"2","total_taken_repay_user_cnt":"0","total_campaign_new_user_cnt":"0","total_camp_period_pid_loss_trade_amt":"3","total_campaign_trade_cnt":"2","today_campaign_new_user_cnt":"0","total_camp_period_pid_loss_trade_user_cnt":"1","total_taken_cnt":"2","total_relative_period_pid_loss_order_amt":"2","total_taken_repay_order_amt":"0","total_relative_period_pid_loss_trade_cnt":"2","total_relative_period_pid_trade_user_cnt":"45","today_taken_cnt":"0","total_used_user_cnt":"1","total_relative_period_pid_repay_order_amt":"4561","total_taken_repay_trade_amt":"0","today_campaign_trade_amt":"0","total_campaign_trade_amt":"2","today_taken_user_cnt":"0","total_taken_repay_trade_cnt":"0"}]`

	var indicatorInfos []Info
	err:=json.Unmarshal([]byte(str),&indicatorInfos)
	fmt.Println("response:", str)

	//result := sdk.AlipassInstanceResult{}
	//err = json.Unmarshal([]byte(resp.Result), &result)
	if err==nil{
		fmt.Println("response:", len(indicatorInfos))
	}

}
