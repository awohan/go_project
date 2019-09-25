package controllers

import (
	"BastionPay/bas-api/apibackend"
	. "BastionPay/bas-base/log/zap"
	"BastionPay/baspay-recharge/api"
	"BastionPay/baspay-recharge/base"
	"BastionPay/baspay-recharge/comsumer"
	"BastionPay/baspay-recharge/config"
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/kataras/iris"
	"go.uber.org/zap"
)

type (
	CommonNotify struct {
		Controllers
	}
)

func (this *CommonNotify) TransferCallBackCommon(ctx iris.Context) {
	params := new(api.NotifyRequest)

	err := ctx.ReadJSON(params)
	if err != nil {
		ZapLog().Error("param err", zap.Error(err))
		return
	}
	ZapLog().Info("[** notify params coming phone: **]", zap.Any("[** notify params coming phone : **]", *params.Phone))
	ZapLog().Info("[** notify params coming countryCode: **]", zap.Any("[** notify params coming countryCode : **]", *params.CountryCode))

	//api-key: 70123b81-9754-4fd4-b2cd-ecac5d3947cd
	if params.ApiKey == nil || *params.ApiKey != "70123b81-9754-4fd4-b2cd-ecac5d3947cd" {
		ZapLog().Error("api key nil or err")
		this.ExceptionSerive(ctx, 1000, "api key err")
		return
	}

	var userId string
	value, ok := params.UserId.(int)
	if ok {
		userId = fmt.Sprintf("%s", value)
		fmt.Println("uid ok,", userId)
	} else {
		userId, _ = params.UserId.(string)
		fmt.Println("uid not ok,", userId)
	}

	//拿对应的token
	if params.SponsorAccount == nil {
		ZapLog().Error("SponsorAccount nil or err")
		this.ExceptionSerive(ctx, 1000, "SponsorAccount err")
		return
	}
	//登录池的map
	//tokenAndRefreshToken := comsumer.GLoginPoolTasker.AccountMap[*params.SponsorAccount]
	//loginPoolInfo := config.GAccountMgr.Get(*params.SponsorAccount)

	if params.UserId == nil || params.UserId == "" || params.UserId == " " {
		fmt.Println("[** start go no uid *****")
		//这里用phone去查询 uid，如果有注册，转账
		//consumer.Task
		code := *params.CountryCode
		if len(code) <= 2 {
			ZapLog().Error("country code err", zap.Error(err))
			return
		}
		code = "+" + code[1:]
		fmt.Println("[** get uid params* phone code common]:", *params.Phone, code)
		fmt.Println("[** get uid params* Token common]:", *comsumer.GTasker.Token)
		uid, registime, err := comsumer.GTasker.GetUidByPhone(*params.Phone, code, *comsumer.GTasker.Token)

		if err != nil || uid == "" || registime == "" {
			ZapLog().Info("get uid by phone nil")
			this.Response(ctx, nil)
			return
		}
		fmt.Println("[** get uid by phone common*]:", uid)

		registTime, err := strconv.ParseInt(registime, 10, 64)
		if params.OffAt == nil {
			*params.OffAt = 9999999
		}
		if registTime/1000 > *params.OffAt {
			ZapLog().Error("activite time out, no transfer...", zap.Error(err))
			this.Response(ctx, nil)
			return
		}

		//根据通知的参数，先把trans_flag从0 设置为1，再充值
		robberUpdate, _ := json.Marshal(map[string]interface{}{
			"id":            *params.Id,
			"transfer_flag": 1,
		})

		res, err := base.HttpSend(config.GConfig.LuckDrawUrl.LuckDrawUrl+"/v1/ft/luckdraw/drawer/set-transferflag", bytes.NewBuffer(robberUpdate), "POST", nil)
		if err != nil {
			ZapLog().Sugar().Errorf("request red err[%v]", err)
		}
		fmt.Println("[** set flag res common***]", string(res))

		responseMsg := new(api.NotifyResponse)
		json.Unmarshal(res, responseMsg)

		if responseMsg.Code == apibackend.BASERR_ACTIVITY_FISSIONSHARE_ROBBER_TRANSFERFLAG_NOT_AFFECTED.Code() {
			this.Response(ctx, nil)
			return
		}

		if responseMsg.Code == 1000100 {
			ZapLog().Info("set flag param nil")
		}
		if responseMsg.Code != 0 {
			ZapLog().Error("set flag err")
		}

		//充值
		pwd := config.GConfig.Login.ZfPwd
		requestNo := time.Now().UnixNano()
		requestNoStr := strconv.FormatInt(requestNo, 10)
		rands := RandString(4)
		requestNoStr = requestNoStr + rands
		//userId := strconv.FormatInt(uid, 10)
		coin := fmt.Sprintf("%s", params.Coin)

		//fmt.Println("**symbol,pwd,token,userid,requestNo,coin**",params.Symbol, pwd, *comsumer.GTasker.Token, uid, requestNo, coin)
		status, err := comsumer.GTasker.Transfer.TransferCoin(*params.Symbol, pwd, *comsumer.GTasker.Token, uid, requestNoStr, coin)
		if err != nil || status != 2 {
			//充值失败，再设置回0
			robberUpdate, _ := json.Marshal(map[string]interface{}{
				"id":            *params.Id,
				"transfer_flag": 0,
			})

			res, err := base.HttpSend(config.GConfig.LuckDrawUrl.LuckDrawUrl+"/v1/ft/luckdraw/drawer/set-transferflag", bytes.NewBuffer(robberUpdate), "POST", nil)
			if err != nil {
				ZapLog().Sugar().Errorf("request red err[%v]", err)
			}
			fmt.Println("[** set flag again res ***]")

			responseMsg := new(api.NotifyResponse)
			json.Unmarshal(res, responseMsg)
			if responseMsg.Code != 0 {
				ZapLog().Error("set flag err")
				fmt.Println("[** set flag again fail....]")
			}
		}
		if err != nil {
			ZapLog().Sugar().Errorf("transfer coin err[%v]", err)
			this.ExceptionSerive(ctx, 1009, err.Error())
			return
		}
		if status == 2 {
			ZapLog().Sugar().Info("transfer succ")
			this.Response(ctx, " ")
			return
		}
		if status == 3 {
			ZapLog().Sugar().Info("transfer fail")
			this.ExceptionSerive(ctx, 1008, "transfer fail")
			return
		}
		if status == 1 {
			ZapLog().Sugar().Info("transfer running...")
			this.ExceptionSerive(ctx, 1007, "transfer running...")
			return
		}
		//只要不是返回的0，都是失败
		this.ExceptionSerive(ctx, 1008, "transfer fail")
		return
	}

	fmt.Println("[** start with uid common*****")
	//根据通知的参数，先把trans_flag从0 设置为1，再充值
	robberUpdate, _ := json.Marshal(map[string]interface{}{
		"id":            *params.Id,
		"transfer_flag": 1,
	})

	res, err := base.HttpSend(config.GConfig.LuckDrawUrl.LuckDrawUrl+"/v1/ft/luckdraw/drawer/set-transferflag", bytes.NewBuffer(robberUpdate), "POST", nil)
	if err != nil {
		ZapLog().Sugar().Errorf("request red err[%v]", err)
		this.ExceptionSerive(ctx, 1008, "set-transferflag fail")
	}
	fmt.Println("[** set flag with uid res common**]", string(res))

	responseMsg := new(api.NotifyResponse)
	json.Unmarshal(res, responseMsg)

	if responseMsg.Code == apibackend.BASERR_ACTIVITY_FISSIONSHARE_ROBBER_TRANSFERFLAG_NOT_AFFECTED.Code() {
		this.Response(ctx, nil)
		return
	}

	if responseMsg.Code == 1000100 {
		ZapLog().Info("set flag param nil")
		this.ExceptionSerive(ctx, 1008, "set-transferflag fail")
		return
	}
	if responseMsg.Code != 0 {
		ZapLog().Error("set flag err", zap.Error(err))
	}

	//充值
	pwd := config.GConfig.Login.ZfPwd
	requestNo := time.Now().UnixNano()
	requestNoStr := strconv.FormatInt(requestNo, 10)
	rands := RandString(4)
	requestNoStr = requestNoStr + rands
	//userId := strconv.FormatInt(*params.UserId, 10)
	//userId := fmt.Sprintf("%v",params.UserId)
	coin := fmt.Sprintf("%s", params.Coin)

	//fmt.Println("**symbol,pwd,token,userid,requestNo,coin**",params.Symbol, pwd, *comsumer.GTasker.Token, userId, requestNo, coin)
	status, err := comsumer.GTasker.Transfer.TransferCoin(*params.Symbol, pwd, config.GConfig.Login.ZfPwd, userId, requestNoStr, coin)
	if err != nil || status != 2 {
		//充值失败，再设置回0
		robberUpdate, _ := json.Marshal(map[string]interface{}{
			"id":            *params.Id,
			"transfer_flag": 0,
		})

		res, err := base.HttpSend(config.GConfig.LuckDrawUrl.LuckDrawUrl+"/v1/ft/luckdraw/drawer/set-transferflag", bytes.NewBuffer(robberUpdate), "POST", nil)
		if err != nil {
			ZapLog().Sugar().Errorf("request red err[%v]", err)
		}
		fmt.Println("[** set flag again res ***]")

		responseMsg := new(api.NotifyResponse)
		json.Unmarshal(res, responseMsg)
		if responseMsg.Code != 0 {
			ZapLog().Error("set flag err")
			fmt.Println("[** set flag again fail....]")
		}
	}
	if err != nil {
		ZapLog().Sugar().Errorf("transfer coin err[%v]", err)
		this.ExceptionSerive(ctx, 1009, err.Error())
		return
	}
	if status == 2 {
		this.Response(ctx, " ")
		return
	}
	if status == 3 {
		ZapLog().Sugar().Info("transfer fail")
		this.ExceptionSerive(ctx, 1008, "transfer fail")
		return
	}
	if status == 1 {
		ZapLog().Sugar().Info("transfer running...")
		this.ExceptionSerive(ctx, 1007, "transfer running...")
		return
	}
	fmt.Println("[** trans res status not normal....]")
	this.ExceptionSerive(ctx, 1010, "transfer err...")
	return
}

//sliceId := make([]int,0)
//sliceId = append(sliceId, params.Id)

//
//var r *rand.Rand
//func init() {
//	r = rand.New(rand.NewSource(time.Now().Unix()))
//}
//
//// RandString 生成随机字符串
//func RandString(len int) string {
//	bytes := make([]byte, len)
//	for i := 0; i < len; i++ {
//		b := r.Intn(26) + 65
//		bytes[i] = byte(b)
//	}
//	return string(bytes)
//}
