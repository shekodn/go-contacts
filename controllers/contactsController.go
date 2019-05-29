package controllers

import (
    "net/http"
    "go-contacts/models"
    "encoding/json"
    u "go-contacts/utils"
)


var CreateContact = func(w http.ResponseWriter, r *http.Request) {
    //Grabs the id of the user that send the request
    user := r.Context().Value("user") . (uint)
    contact := &models.Contact{}

    err := json.NewDecoder(r.Body).Decode(contact)

    if err != nil {
        u.Respond(w, u.Message(false, "Error while decoding request body"))
        return
    }

    contact.UserId = user
    resp := contact.Create()
    u.Respond(w, resp)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {

    id := r.Context().Value("user") . (uint)
    data := models.GetContacts(uint(id))
    resp := u.Message(true, "success")
    resp["data"] = data
    u.Respond(w, resp)
}
