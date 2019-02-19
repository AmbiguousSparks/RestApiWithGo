package controller

import (
	. "ApiUsers/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

var Users []User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	//All()
	var user User
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(user.All())
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	var user User

	user.ID = params["id"]

	user.ReadOne()

	json.NewEncoder(w).Encode(user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	var user User
	file, handle, err := r.FormFile("filePhoto")
	json.Unmarshal([]byte(r.PostFormValue("data")), &user)

	user.ID = params["id"]

	//verifica se alguma foto foi enviada
	if err != nil {
		//se não, ele só atualiza os dados
		res := user.UpdateNoPhoto()
		if res {
			jsonResponse(w, http.StatusOK, "User updated")
		}
		return
	}

	//se foi ele atualiza os dados e a foto
	res := user.UpdatePhoto()
	if res {
		saveFile(w, file, handle)
	}

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	params := mux.Vars(r)

	var user User

	user.ID = params["id"]

	//le os dados do usuario em questao para pegar qual foto corresponde ao mesmo
	user.ReadOne()

	//caminho aonde a imagem vai ser deletada
	path := "./uploads/"

	//deletar a foto do usuario
	os.Remove(path + user.Photo)

	res := user.Delete()

	if res {
		jsonResponse(w, http.StatusOK, "User successfully deleted!")
		return
	}
	jsonResponse(w, http.StatusBadRequest, "User not deleted!")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//pega o arquivo da foto recebida
	file, handle, err := r.FormFile("filePhoto")

	var user User

	err = json.Unmarshal([]byte(r.PostForm["data"][0]), &user)

	if err != nil {
		fmt.Println(err)
	}

	res := user.Save()

	if !res {
		fmt.Println()
		return
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")

	switch mimeType {
	case "image/jpg":
		saveFile(w, file, handle)
	case "image/jpeg":
		saveFile(w, file, handle)
	case "image/png":
		saveFile(w, file, handle)
	default:
		jsonResponse(w, http.StatusBadRequest, "The format file is not valid")
	}
}

func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader) {

	//função para salvar arquivos recebidos pelo servidor

	data, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
		return
	}

	//caminho aonde o arquivo vai ser salvo
	path := "./uploads/"
	err = ioutil.WriteFile(path+handle.Filename, data, 0666)
	if err != nil {
		log.Fatal(err)
	}

	jsonResponse(w, http.StatusCreated, "OK!")
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	/*função para dar uma resposta de string em formato json
	 */
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}
