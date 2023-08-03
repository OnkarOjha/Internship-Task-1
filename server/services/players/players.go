package players

import (
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/utils"
	"math/rand"
	"sort"
	"time"

	"github.com/gin-gonic/gin"

	"strconv"
)

// Players is a slice to hold all player entries
var Players []model.Player

// Helpers
func getPlayerIndex(id int) int {
	for i, p := range Players {
		if p.ID == id {
			return i
		}
	}

	return -1
}

// CreatePlayerService creates a new player entry and adds it to the Players slice.
// It generates a new unique ID for the player and sends a JSON response with the newly created player details.
func CreatePlayerService(ctx *gin.Context, player model.Player) {

	// Generate a new unique ID for the player
	player.ID = len(Players) + 1

	Players = append(Players, player)

	response.ShowResponse("SUCCESS", utils.HTTP_OK, "New Player Created", Players, ctx)

}

// GetAllPlayerService retrieves all players from the Players slice, sorts them by score in descending order,
// and sends a JSON response with the list of players.
func GetAllPlayerService(ctx *gin.Context) {

	// Sort players by score in descending order
	sort.Slice(Players, func(i, j int) bool {
		return Players[i].Score > Players[j].Score
	})
	response.ShowResponse("SUCCESS", utils.HTTP_OK, "Requested Player", Players, ctx)

}

// UpdatePlayerService updates the name and score of an existing player with the provided ID.
// It receives the updated player details through the PlayerUpdateRequest and sends a JSON response
// with the updated player details.
func UpdatePlayerService(ctx *gin.Context, updatedPlayer request.PlayerUpdateRequest) {
	idParam, _ := ctx.Params.Get("id")

	playerID, err := strconv.Atoi(idParam)
	if err != nil {
		response.ShowResponse("FAILURE", utils.HTTP_BAD_REQUEST, "Error fetching ID", err.Error(), ctx)
		return
	}

	playerIndex := getPlayerIndex(playerID)

	if playerIndex == -1 {

		response.ShowResponse("FAILURE", utils.HTTP_BAD_REQUEST, "Player with given ID not found", nil, ctx)

		return
	}

	// Only update name and score
	Players[playerIndex].Name = updatedPlayer.Name
	Players[playerIndex].Score = updatedPlayer.Score

	response.ShowResponse("SUCCESS", utils.HTTP_OK, "Requested Player", Players, ctx)
}

// DeletePlayerService deletes a player with the given ID from the Players slice.
// It sends a JSON response with the updated list of players after removing the specified player.
func DeletePlayerService(ctx *gin.Context) {
	idParam, _ := ctx.Params.Get("id")

	playerID, err := strconv.Atoi(idParam)
	if err != nil {
		response.ShowResponse("FAILURE", utils.HTTP_BAD_REQUEST, "Error fetching ID", err.Error(), ctx)
		return
	}

	playerIndex := getPlayerIndex(playerID)

	if playerIndex == -1 {

		response.ShowResponse("FAILURE", utils.HTTP_BAD_REQUEST, "Player with given ID not found", nil, ctx)

		return
	}

	// Remove the player from the slice
	Players = append(Players[:playerIndex], Players[playerIndex+1:]...)

	response.ShowResponse("SUCCESS", utils.HTTP_OK, "Requested Player", Players, ctx)
}


// GetPlayerByRankService retrieves the player ranked at the given value from the Players slice,
// sorted by score in descending order. If the rank is out of range, it sends an error JSON response.
func GetPlayerByRankService(ctx *gin.Context) {
	val, _ := ctx.Params.Get("val")

	rank, err := strconv.Atoi(val)
	if err != nil {
		response.ShowResponse("FAILURE", utils.HTTP_BAD_REQUEST, "Error fetching ID", nil, ctx)
		return
	}

	if rank <= 0 || rank > len(Players) {
		response.ShowResponse("FAILURE", utils.HTTP_BAD_REQUEST, "Rank out of range", nil, ctx)
		return
	}

	// Sort players by score in descending order
	sort.Slice(Players, func(i, j int) bool {
		return Players[i].Score > Players[j].Score
	})

	player := Players[rank-1]
	response.ShowResponse("SUCCESS", utils.HTTP_OK, "Players according to their rank", player, ctx)

}

// GetRandomPlayerService fetches a random player from the Players slice and sends a JSON response.
// If no players are available, it sends an error JSON response.
func GetRandomPlayerService(ctx *gin.Context) {
	if len(Players) == 0 {
		response.ShowResponse("FAILURE", utils.HTTP_BAD_REQUEST, "No Player Found", nil, ctx)
		return
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(Players))
	response.ShowResponse("SUCCESS", utils.HTTP_OK, "Random Player fetched", Players[random], ctx)
}
