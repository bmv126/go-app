package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Member of the Group
type Member struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// Members of the Group
var Members []Member = []Member{}

// GetAllMembers : To get all members
func GetAllMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Members)
}

// DeleteAll : Delete all members
func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Members = Members[:0]
	json.NewEncoder(w).Encode(Members)
}

// GetMember : GET operation
func GetMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	MemberName := mux.Vars(r)["name"]

	found := false
	for _, EachMember := range Members {
		if EachMember.Name == MemberName {
			json.NewEncoder(w).Encode(EachMember)
			found = true
			break
		}
	}
	if found == false {
		json.NewEncoder(w).Encode(struct {
			Error string
		}{MemberName + " not a member"})
	}
}

// DeleteMember : DELETE operation
func DeleteMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	MemberName := mux.Vars(r)["name"]

	toDeleteIndex := -1
	for index, EachMember := range Members {
		if EachMember.Name == MemberName {
			toDeleteIndex = index
			break
		}
	}
	if toDeleteIndex == -1 {
		json.NewEncoder(w).Encode(struct {
			Error string
		}{MemberName + " not a member"})
	} else {
		Members = append(Members[:toDeleteIndex], Members[toDeleteIndex+1:]...)
		json.NewEncoder(w).Encode(struct {
			Info string
		}{MemberName + " member removed successfully !!! "})
	}

}

// AddMember : POST operation
func AddMember(w http.ResponseWriter, r *http.Request) {

	var MemberInfo Member

	json.NewDecoder(r.Body).Decode(&MemberInfo)
	Members = append(Members, MemberInfo)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(struct {
		Info string
	}{MemberInfo.Name + " member added successfully !!! "})

}

// AddMembers : Add multiple members
func AddMembers(w http.ResponseWriter, r *http.Request) {

	var ListMemberInfo []Member

	json.NewDecoder(r.Body).Decode(&ListMemberInfo)
	Members = append(Members, ListMemberInfo...)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(struct {
		Info string
	}{" members added successfully !!! "})

}

// ModifyMember : PUT operation
func ModifyMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	MemberName := mux.Vars(r)["name"]

	var MemberInfo Member
	json.NewDecoder(r.Body).Decode(&MemberInfo)

	if MemberName != MemberInfo.Name {
		json.NewEncoder(w).Encode(struct {
			Error string
		}{MemberInfo.Name + " is incorrect. It does not match with routePath /" + MemberName})
		return
	}

	toModifyIndex := -1
	for index, EachMember := range Members {
		if EachMember.Name == MemberInfo.Name {
			toModifyIndex = index
			break
		}
	}
	if toModifyIndex == -1 {
		json.NewEncoder(w).Encode(struct {
			Error string
		}{MemberName + " not a member"})
	} else {
		Members[toModifyIndex] = MemberInfo
		json.NewEncoder(w).Encode(struct {
			Info string
		}{MemberName + " member modified successfully !!! "})
	}
}
