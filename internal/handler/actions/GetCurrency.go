package actions

import (
	base "GolangCurrencyMS/internal/handler"
	"GolangCurrencyMS/internal/payload"
	"GolangCurrencyMS/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetCurrencyController struct {
}

func NewGetCurrencyController() base.Controller {
	return &GetCurrencyController{}
}

func (c *GetCurrencyController) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get keyword from url
		currencyName := ctx.Param("item")
		sourceUrl := "https://api.hitbtc.com/api/3/public"

		var strings []string
		if currencyName == "all" {
			// Get all Symbol names from server
			strings = utils.GetAllKeyValues(sourceUrl + "/symbol")
		} else {
			// Set current keyword as Symbol name
			strings = append(strings, currencyName)
		}

		currencyList := make([]payload.DataResponse, len(strings))
		for i, item := range strings {

			newTicker := utils.GetDataFromAPI(sourceUrl+"/ticker/"+item, payload.TickerRequest{})
			newSymbol := utils.GetDataFromAPI(sourceUrl+"/symbol/"+item, payload.SymbolRequest{})
			fullNameForSymbol := utils.GetDataFromAPI(sourceUrl+"/currency/"+newSymbol.BaseCurrency, payload.CurrencyRequest{})
			// Create new currency Information
			newCurrency := payload.DataResponse{
				Id:          newSymbol.BaseCurrency,
				FullName:    fullNameForSymbol.FullName,
				Ask:         utils.ConvertTofloat(newTicker.Ask, 64),
				Bid:         utils.ConvertTofloat(newTicker.Bid, 64),
				Last:        utils.ConvertTofloat(newTicker.Last, 64),
				Open:        utils.ConvertTofloat(newTicker.Open, 64),
				Low:         utils.ConvertTofloat(newTicker.Low, 64),
				High:        utils.ConvertTofloat(newTicker.High, 64),
				FeeCurrency: newSymbol.FeeCurrency,
			}
			// Add individual currency information to the list
			currencyList[i] = newCurrency
		}

		if currencyName == "all" {
			ctx.JSON(http.StatusOK, gin.H{"currencies": currencyList})
		} else {
			ctx.JSONP(http.StatusOK, currencyList[0])
		}
	}
}

func (c *GetCurrencyController) Route() string {
	return "/currency/:item"
}

func (c *GetCurrencyController) Method() string {
	return http.MethodGet
}
