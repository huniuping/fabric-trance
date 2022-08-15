package assemble

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	assemble_sdk "raft-fabric-project/application/sdk/assemble"
	"strings"
)

func Query(ctx *gin.Context) {
	assemble_product_id := ctx.Query("assemble_product_id")

	chaincode_name := "assemblecc"
	fnc := "query"
	args := [][]byte{[]byte(assemble_product_id)}
	rsp, err := assemble_sdk.ChannelQuery(chaincode_name, fnc, args)
	if err != nil||string(rsp.Payload)=="" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "查询失败",
			"data": nil,
		})
		return
	}
	fmt.Println("=============")
	xx:=string(rsp.Payload)
	fmt.Println(xx)
	map_data := make(map[string]interface{})
	json.Unmarshal([]byte(string(rsp.Payload)), &map_data)
	fmt.Println(map_data)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"meg":  "查询成功",
		"data": map_data,
	})
	return
}

func QueryAssembleDetailByID(ctx *gin.Context) {
	assemble_product_id := ctx.Query("assemble_product_id")
	chaincode_name := "assemblecc"
	fnc := "queryAssembleDetailByID"
	args := [][]byte{[]byte(assemble_product_id)}
	rsp, err := assemble_sdk.ChannelQuery(chaincode_name, fnc, args)
	if err != nil||string(rsp.Payload)=="" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "查询失败",
			"data": nil,
		})
		return
	}
	payload_s:=string(rsp.Payload)
	payloads :=strings.Split(payload_s,";;")
	map_data:=map[string]interface{}{}
	assembles:=[]map[string]interface{}{}
	for i := 0; i < len(payloads); i++ {
		if payloads[i] ==""{
			continue
		}
		temp_map := make(map[string]interface{})//一行所有数据存在这一个map中
		json.Unmarshal([]byte(payloads[i]), &temp_map)
		assembles=append(assembles,temp_map)
	}
	map_data["assemblesHistory"]=assembles
	fmt.Println(map_data)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"meg":  "查询成功",
		"data": map_data,
	})
	return
}

func QueryByworkOrderId(ctx *gin.Context) {
	work_order_Id := ctx.Query("work_order_Id")
	chaincode_name := "assemblecc"
	fnc := "queryByContractId"
	args := [][]byte{[]byte(work_order_Id)}
	rsp, err := assemble_sdk.ChannelQuery(chaincode_name, fnc, args)
	if err != nil||string(rsp.Payload)=="" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "查询失败",
			"data": nil,
		})
		return
	}
	payload_s:=string(rsp.Payload)
	payloads :=strings.Split(payload_s,";;")

	map_data:=map[string]interface{}{}
	work_order:=[]map[string]interface{}{}

	for i := 0; i < len(payloads); i++ {
		temp_map := make(map[string]interface{})//一行所有数据存在这一个map中
		json.Unmarshal([]byte(payloads[i]), &temp_map)
		work_order=append(work_order,temp_map)
	}
	map_data["work_order_Id"]=work_order_Id
	fmt.Println(map_data)

	fmt.Println(map_data)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"meg":  "查询成功",
		"data": map_data,

	})
	return
}

func Delete(ctx *gin.Context) {
	assemble_product_id := ctx.PostForm("assemble_product_id")
	chaincode_name := "assemblecc"
	fnc := "delete"
	args := [][]byte{[]byte(assemble_product_id)}
	_, err := assemble_sdk.ChannelExecute(chaincode_name, fnc, args)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "删除成功",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"meg":  "删除失败",
	})
	return
}

func Set(ctx *gin.Context) {
	assemble_product_id := ctx.PostForm("assemble_product_id")
	assemble_product_name := ctx.PostForm("assemble_product_name")
	work_order_id := ctx.PostForm("work_order_id")
	assemble_line_id := ctx.PostForm("assemble_line_id")
	date := ctx.PostForm("date")
	process_list := ctx.PostForm("process_list")
	technology := ctx.PostForm("technology")
	chaincode_name := "assemblecc"
	fnc := "set"
	args := [][]byte{
		[]byte(assemble_product_id),
		[]byte(assemble_product_name),
		[]byte(work_order_id),
		[]byte(assemble_line_id),
		[]byte(date),
		[]byte(process_list),
		[]byte(technology)}
	_, err := assemble_sdk.ChannelExecute(chaincode_name, fnc, args)
	fmt.Println("====================")
	fmt.Println(err)
	if error.Error(err) == "Client Status Code: (5) TIMEOUT. Description: request timed out or been cancelled" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "已经添加成功",
		})
		return
	} else if error.Error(err) == "Transaction processing for endorser [localhost:11051]: Chaincode status Code: (500) UNKNOWN. Description: 该产品标识号已存在" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "该组装产品已存在",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "添加失败",
		})
		return
	}
}

func Update(ctx *gin.Context) {
	assemble_product_id := ctx.PostForm("assemble_product_id")
	assemble_product_name := ctx.PostForm("assemble_product_name")
	work_order_id := ctx.PostForm("work_order_id")
	assemble_line_id := ctx.PostForm("assemble_line_id")
	date := ctx.PostForm("date")
	process_list := ctx.PostForm("process_list")
	technology := ctx.PostForm("technology")
	chaincode_name := "assemblecc"
	fnc := "update"
	args := [][]byte{
		[]byte(assemble_product_id),
		[]byte(assemble_product_name),
		[]byte(work_order_id),
		[]byte(assemble_line_id),
		[]byte(date),
		[]byte(process_list),
		[]byte(technology)}
	_, err := assemble_sdk.ChannelExecute(chaincode_name, fnc, args)
	fmt.Println("====================")
	fmt.Println(err)
	if error.Error(err) == "Client Status Code: (5) TIMEOUT. Description: request timed out or been cancelled" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "已经更新成功",
		})
		return
	} else if error.Error(err) == "Transaction processing for endorser [localhost:11051]: Chaincode status Code: (500) UNKNOWN. Description: 未找到需要更新的记录" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "该组装产品已存在",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "更新失败",
		})
		return
	}
}
