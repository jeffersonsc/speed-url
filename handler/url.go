package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jeffersonsc/speed-url/repo"

	"github.com/jeffersonsc/speed-url/lib/context"
	"github.com/jeffersonsc/speed-url/model"
)

func CreateURL(ctx *context.Context) {
	// Read Body of req
	body, err := ctx.Req.Body().Bytes()
	defer ctx.Req.Body().ReadCloser()
	if err != nil {
		log.Println("[urls-handler] erro ao pegar infos ", err.Error())
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": "formato do json inválido"})
		return
	}

	myurl := model.MyURL{}
	json.Unmarshal(body, &myurl)

	myurl, err = repo.SaveURL(myurl.LongURL)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Erro ao gerar url. ERROR:" + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, myurl)

}

func FindURL(ctx *context.Context) {
	key := ctx.Params("id")
	url, err := repo.FindURL(key, ctx.RemoteAddr())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Erro ao buscar url. ERROR:" + err.Error()})
		return
	}
	if url == "" {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Url inválida"})
		return
	}

	ctx.Redirect(url, http.StatusTemporaryRedirect)
}

func ShowURL(ctx *context.Context) {
	log.Println("URL>>>> ", ctx.Params("id"))
	myurl, err := repo.ExplaneURL(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Url não localizada. ERROR: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, myurl)
}
