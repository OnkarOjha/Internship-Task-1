package handler

import (
	
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/services/players"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func CreatePlayer(ctx *gin.Context) {
	utils.SetHeader(ctx)
	var player model.Player
	err := utils.RequestDecoding(ctx, &player)
	if err != nil {
		response.ShowResponse("FAILURE", utils.HTTP_BAD_REQUEST, "Error Decoding Values", err.Error(), ctx)
		return
	}

	err = player.Validate()
	if err != nil {
		response.ShowResponse("FAILURE", utils.HTTP_BAD_REQUEST, "Error Validating Request", err.Error(), ctx)
		return
	}

	players.CreatePlayerService(ctx, player)
}

func GetPlayer(ctx *gin.Context) {
	utils.SetHeader(ctx)
	players.GetAllPlayerService(ctx)
}

func UpdatePlayer(ctx *gin.Context) {
	utils.SetHeader(ctx)
	var playerUpdate request.PlayerUpdateRequest
	err := utils.RequestDecoding(ctx, &playerUpdate)
	if err != nil {
		response.ShowResponse("FAILURE", utils.HTTP_BAD_REQUEST, "Error Decoding Values", err.Error(), ctx)
		return
	}

	err = playerUpdate.Validate()
	if err != nil {
		response.ShowResponse("FAILURE", utils.HTTP_BAD_REQUEST, "Error Validating Request", err.Error(), ctx)
		return
	}
	
	players.UpdatePlayerService(ctx, playerUpdate)
}


func DeletePlayer(ctx *gin.Context){
	utils.SetHeader(ctx)
	players.DeletePlayerService(ctx)
}

func GetPlayerByRank(ctx *gin.Context){
	utils.SetHeader(ctx)
	players.GetPlayerByRankService(ctx)
}

func GetRandomPlayer(ctx *gin.Context){
	utils.SetHeader(ctx)
	players.GetRandomPlayerService(ctx)
}